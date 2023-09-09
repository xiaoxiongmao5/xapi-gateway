package middleware

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// CORSMiddleware 是处理跨域请求的中间件
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 设置允许跨域请求的来源域，这里需要设置为请求的 Origin
		c.Writer.Header().Set("Access-Control-Allow-Origin", c.GetHeader("Origin"))

		// 允许的 HTTP 方法
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, OPTIONS")

		// 允许的请求标头
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, Nonce, timestamp, accessKey, sign")

		// 允许携带 Cookie
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		// 如果是预检请求（OPTIONS 请求），直接返回成功状态
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(200)
			return
		}
		fmt.Println("[middleware 处理跨域]CORSMiddleware complete!")
		c.Next()
	}
}
