package controller

import (
	"net/http"
	"xj/xapi-gateway/enums"
	ghandle "xj/xapi-gateway/g_handle"
	glog "xj/xapi-gateway/g_log"
	"xj/xapi-gateway/middleware"

	"github.com/gin-gonic/gin"
)

type RateLimitConfig struct {
	IP                string  `json:"ip"`
	RequestsPerSecond float64 `json:"requests_per_second"`
	BucketSize        int     `json:"bucket_size"`
}

// @Summary		获得具体IP的限流配置
// @Description	获得具体IP的限流配置
// @Tags			管理配置
// @Accept			application/x-www-form-urlencoded
// @Produce		application/json
// @Param			ip	query		string	true	"ip地址"
// @Success		200	{object}	object	"具体IP的限流配置"
// @Router			/manage/config/ratelimit [get]
func GetIPRateLimitConfig(c *gin.Context) {
	ip := c.Query("ip")
	if ip == "" {
		c.JSON(http.StatusOK, gin.H{"result": enums.ParameterError, "msg": "参数错误"})
		return
	}
	requestsPerSecond, bucketSize := middleware.GetIPRateLimitConfig(ip)

	resData := RateLimitConfig{
		IP:                ip,
		RequestsPerSecond: requestsPerSecond,
		BucketSize:        bucketSize,
	}

	ghandle.HandlerSuccess(c, "IP的限流配置获取成功", resData)
}

// @Summary		更新具体IP限流配置失败
// @Description	更新具体IP限流配置失败
// @Tags			管理配置
// @Accept			application/x-www-form-urlencoded
// @Produce		application/json
// @Param			request	body		RateLimitConfig	true	"限流配置"
// @Success		200		{object}	object
// @Router			/manage/config/ratelimit [put]
func UpdateIPRateLimitConfig(c *gin.Context) {
	var params *RateLimitConfig
	if err := c.ShouldBindJSON(&params); err != nil {
		glog.Log.Errorf("RateLimitConfig err=%v", err.Error())
		c.JSON(http.StatusOK, gin.H{"result": enums.ParameterError, "msg": "参数错误"})
		return
	}
	if params.RequestsPerSecond <= 0 || params.BucketSize <= 0 {
		glog.Log.Errorf("参数IP限流配置无效, RequestsPerSecond=%v, BucketSize=%v", params.RequestsPerSecond, params.BucketSize)
		c.JSON(http.StatusOK, gin.H{"result": enums.ParameterError, "msg": "限流配置无效"})
		return
	}
	if err := middleware.UpdateIPRateLimitConfig(params.IP, params.RequestsPerSecond, params.BucketSize); err != nil {
		c.JSON(http.StatusOK, gin.H{"result": enums.UpdateIPRateLimitConfigFailed, "msg": "更新IP限流配置失败"})
		return
	}

	ghandle.HandlerSuccess(c, "更新IP限流配置成功", nil)
}
