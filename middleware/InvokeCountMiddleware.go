package middleware

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// 调用成功，接口调用次数+1
func InvokeCountMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// todo 更新接口调用次数
		fmt.Println("更新接口调用次数")
		c.Next()
	}
}
