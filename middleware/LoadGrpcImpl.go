package middleware

import (
	"fmt"
	ghandle "xj/xapi-gateway/g_handle"
	"xj/xapi-gateway/rpc_api"

	"dubbo.apache.org/dubbo-go/v3/config"
	"github.com/gin-gonic/gin"
)

var grpcInterfaceInfoImpl = new(rpc_api.IntefaceInfoClientImpl)
var grpcUserInfoImpl = new(rpc_api.UserInfoClientImpl)
var grpcUserInterfaceInfoImpl = new(rpc_api.UserIntefaceInfoClientImpl)
var replyGetInvokeUser *rpc_api.GetInvokeUserResp
var replyGetInterfaceInfoByIdReq *rpc_api.GetInterfaceInfoByIdResp

func init() {
	fmt.Println("LoadGrpcImpl init ~~")
}
func LoadGrpcImpl() gin.HandlerFunc {
	return func(c *gin.Context) {
		config.SetConsumerService(grpcUserInfoImpl)
		config.SetConsumerService(grpcInterfaceInfoImpl)
		config.SetConsumerService(grpcUserInterfaceInfoImpl)
		if err := config.Load(); err != nil {
			ghandle.HandlerDobboLoadFailed(c, "UserIntefaceInfoClientImpl")
			return
		}
		fmt.Println("LoadGrpcImpl complete![加载GrpcImpl]")
		c.Next()
	}
}
