package middleware

import (
	"context"
	"fmt"
	ghandle "xj/xapi-gateway/g_handle"
	"xj/xapi-gateway/rpc_api"
	"xj/xapi-gateway/utils"

	"dubbo.apache.org/dubbo-go/v3/common/logger"
	"github.com/gin-gonic/gin"
)

func ValidInterfaceInfo() gin.HandlerFunc {
	return func(c *gin.Context) {
		interfaceIdstr := c.Request.Header.Get("gateway_transdata")
		interfaceId, err := utils.String2Int64(interfaceIdstr)
		if err != nil {
			ghandle.HandlerSuccess(c, "参数错误", nil)
			c.Abort()
			return
		}
		fmt.Println("开始调用RPC: 获得接口信息, 接口ID=", interfaceId)

		reply, err := grpcInterfaceInfoImpl.GetInterfaceInfoById(context.Background(), &rpc_api.GetInterfaceInfoByIdReq{InterfaceId: interfaceId})
		if err != nil {
			ghandle.HandlerSuccess(c, "参数错误", nil)
			c.Abort()
			return
		}
		logger.Infof("get reply~~: %v\n", reply)
		replyGetInterfaceInfoByIdReq = reply
		fmt.Println("ValidInterfaceInfo complete![验证请求的接口是否存在]")
		c.Next()
	}
}
