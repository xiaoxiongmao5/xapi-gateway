package middleware

import (
	"net/http"
	"strconv"
	"time"
	"xj/xapi-gateway/utils"

	"github.com/gin-gonic/gin"
)

func HandlerNoAuth(c *gin.Context) {
	c.JSON(http.StatusForbidden, gin.H{"error": "无权限"})
	c.Abort()
}
func SignMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 从请求头中获取参数
		accessKey := c.Request.Header.Get("accessKey")
		nonce := c.Request.Header.Get("nonce")
		timestamp := c.Request.Header.Get("timestamp")
		sign := c.Request.Header.Get("sign")
		body := c.Request.Header.Get("body")

		// ToDo 去数据库中查是否已分配给用户
		if accessKey != "xiaoxiong" {
			HandlerNoAuth(c)
			return
		}
		secretKey := ""
		// 校验如果随机数大于1万，则抛出异常
		if num, err := strconv.Atoi(nonce); err != nil || num > 10000 {
			HandlerNoAuth(c)
			return
		}
		// 时间和当前时间不能超过5分钟
		fiveMinutes := int64(5 * 60)
		timestampNow := time.Now().Unix()
		if tsInt, err := strconv.ParseInt(timestamp, 10, 64); err != nil || timestampNow-tsInt >= fiveMinutes {
			HandlerNoAuth(c)
			return
		}
		// 如果生成的签名不一致，则抛出异常
		str := accessKey + nonce + timestamp + body
		signEcrypt := utils.GetAPISign(str, secretKey)
		if signEcrypt != sign {
			HandlerNoAuth(c)
			return
		}
		c.Next()
	}
}
