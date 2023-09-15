package middleware

import (
	"errors"
	"net/http"
	"sync"
	gconfig "xj/xapi-gateway/g_config"

	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
)

var (
	rateLimitMu sync.Mutex
	IPLimiter   *IPRateLimiter
)

func NewIPRateLimiter() *IPRateLimiter {
	return &IPRateLimiter{
		limiter: make(map[string]*rate.Limiter),
	}
}

type IPRateLimiter struct {
	mu      sync.Mutex
	limiter map[string]*rate.Limiter
}

func (i *IPRateLimiter) GetLimiter(ip string) *rate.Limiter {
	i.mu.Lock()
	defer i.mu.Unlock()

	limiter, exists := i.limiter[ip]
	if !exists {
		limiter = rate.NewLimiter(rate.Limit(gconfig.AppConfigDynamic.RateLimitConfig.RequestsPerSecond), gconfig.AppConfigDynamic.RateLimitConfig.BucketSize)
		i.limiter[ip] = limiter
	}

	return limiter
}

// 定义一个中间件函数来进行限流
func IPRateLimiterMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		ip := c.ClientIP()

		limiter := IPLimiter.GetLimiter(ip)

		if !limiter.Allow() {
			c.JSON(http.StatusOK, gin.H{"result": http.StatusTooManyRequests, "msg": "Rate limit exceeded"})
			c.Abort()
			return
		}

		c.Next()
	}
}

// 该方法用于动态更新具体IP的限流配置
func UpdateIPRateLimitConfig(ip string, requestsPerSecond float64, bucketSize int) error {
	// 在此处进行配置验证，确保新的配置是有效的
	if requestsPerSecond <= 0 || bucketSize <= 0 {
		return errors.New("Invalid configuration")
	}

	rateLimitMu.Lock()
	defer rateLimitMu.Unlock()

	limiter := IPLimiter.GetLimiter(ip)

	// 更新IP的限流配置
	limiter.SetLimit(rate.Limit(requestsPerSecond))
	limiter.SetBurst(bucketSize)

	return nil
}

// 该方法用于获得具体IP的限流配置
func GetIPRateLimitConfig(ip string) (requestsPerSecond float64, bucketSize int) {
	rateLimitMu.Lock()
	defer rateLimitMu.Unlock()

	limiter := IPLimiter.GetLimiter(ip)

	requestsPerSecond = float64(limiter.Limit())
	bucketSize = limiter.Burst()

	return
}
