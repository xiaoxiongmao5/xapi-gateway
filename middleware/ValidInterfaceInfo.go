package middleware

import (
	"context"
	"net/http"
	"xj/xapi-gateway/rpc_api"
	"xj/xapi-gateway/utils"

	"dubbo.apache.org/dubbo-go/v3/common/logger"
	dubboConfig "dubbo.apache.org/dubbo-go/v3/config"
	"github.com/gin-gonic/gin"
)

var grpcInterfaceInfoImpl = new(rpc_api.IntefaceInfoClientImpl)

func ValidInterfaceInfo() gin.HandlerFunc {
	return func(c *gin.Context) {
		dubboConfig.SetConsumerService(grpcInterfaceInfoImpl)
		if err := dubboConfig.Load(); err != nil {
			panic(err)
		}
		domain, err := utils.GetDomainFromReferer(c.Request.Referer())
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"error": "get domain failed"})
			c.Abort()
			return
		}
		reply, err := grpcInterfaceInfoImpl.GetInterfaceInfo(context.Background(), &rpc_api.GetInterfaceInfoReq{
			Host:   domain,
			Path:   c.Param("path"),
			Method: c.Request.Method,
		})
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"error": "参数错误"})
			c.Abort()
			return
		}
		c.Set("cur_interfaceinfo", reply)
		logger.Infof("get reply~~: %v\n", reply)
	}
}
