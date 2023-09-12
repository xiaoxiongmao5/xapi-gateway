package middleware

import (
	ghandle "xj/xapi-gateway/g_handle"
	glog "xj/xapi-gateway/g_log"
	"xj/xapi-gateway/rpc_api"

	"dubbo.apache.org/dubbo-go/v3/config"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

var grpcInterfaceInfoImpl = new(rpc_api.IntefaceInfoClientImpl)
var grpcUserInfoImpl = new(rpc_api.UserInfoClientImpl)
var grpcUserInterfaceInfoImpl = new(rpc_api.UserIntefaceInfoClientImpl)
var replyGetInvokeUser *rpc_api.GetInvokeUserResp
var replyGetInterfaceInfoById *rpc_api.GetInterfaceInfoByIdResp
var replyGetFullUserInterfaceInfo *rpc_api.GetFullUserInterfaceInfoResp

func LoadGrpcImpl() gin.HandlerFunc {
	return func(c *gin.Context) {
		config.SetConsumerService(grpcUserInfoImpl)
		config.SetConsumerService(grpcInterfaceInfoImpl)
		config.SetConsumerService(grpcUserInterfaceInfoImpl)
		if err := config.Load(); err != nil {
			ghandle.HandlerDobboLoadFailed(c, "UserIntefaceInfoClientImpl")
			return
		}

		glog.Log.WithFields(logrus.Fields{
			"pass": true,
		}).Info("middleware-加载GrpcImpl")

		c.Next()
	}
}
