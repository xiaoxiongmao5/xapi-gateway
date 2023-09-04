package middleware

import (
	"context"
	ghandle "xj/xapi-gateway/g_handle"
	"xj/xapi-gateway/rpc_api"

	"dubbo.apache.org/dubbo-go/v3/common/logger"
	dubboConfig "dubbo.apache.org/dubbo-go/v3/config"
	"github.com/gin-gonic/gin"
)

var grpcUserInterfaceInfoImpl = new(rpc_api.UserIntefaceInfoClientImpl)

// 调用成功，接口调用次数+1
func InvokeCountMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		dubboConfig.SetConsumerService(grpcUserInterfaceInfoImpl)
		if err := dubboConfig.Load(); err != nil {
			panic(err)
		}
		userinfo, exists := c.Get("cur_userinfo")
		if !exists {
			ghandle.HandlerContextError(c, "cur_userinfo")
			return
		}
		interfaceinfo, exists := c.Get("cur_interfaceinfo")
		if !exists {
			ghandle.HandlerContextError(c, "cur_interfaceinfo")
			return
		}
		// 更新接口调用次数
		reply, err := grpcUserInterfaceInfoImpl.InvokeCount(context.Background(), &rpc_api.InvokeCountReq{
			InterfaceId: interfaceinfo.(*rpc_api.GetInterfaceInfoResp).Id,
			UserId:      userinfo.(*rpc_api.GetInvokeUserResp).Id,
		})
		if err != nil {
			logger.Error(err)
		}
		logger.Infof("更新接口调用次数 InvokeCount get reply~~: %v\n", reply)
		c.Next()
	}
}
