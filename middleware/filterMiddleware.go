package middleware

import (
	"fmt"
	gconfig "xj/xapi-gateway/g_config"
	ghandle "xj/xapi-gateway/g_handle"
	"xj/xapi-gateway/utils"

	"github.com/gin-gonic/gin"
)

// 访问控制（黑白名单）
func FilterWithAccessControl() gin.HandlerFunc {
	return func(c *gin.Context) {
		IP_WHITE_LIST := gconfig.AppConfig.IPWhiteList
		reqIp := utils.GetRequestIp(c)
		flag := false
		for _, val := range IP_WHITE_LIST {
			if val == reqIp {
				flag = true
			}
		}
		if flag {
			fmt.Println("[middleware 访问控制（黑白名单）]FilterWithAccessControl complete!")
			c.Next()
		} else {
			ghandle.HandlerForbidden(c)
			return
		}
	}
}
