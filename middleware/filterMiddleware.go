package middleware

import (
	"net/http"
	"xj/xapi-gateway/utils"

	"github.com/gin-gonic/gin"
)

// 访问控制（黑白名单）
func FilterWithAccessControl() gin.HandlerFunc {
	return func(c *gin.Context) {
		IP_WHITE_LIST := []string{"127.0.0.1"}
		reqIp := utils.GetRequestIp(c)
		flag := false
		for _, val := range IP_WHITE_LIST {
			if val == reqIp {
				flag = true
			}
		}
		if flag {
			c.Next()
		} else {
			c.JSON(http.StatusForbidden, gin.H{"error": "no auth"})
			c.Abort()
			return
		}
	}
}
