package middleware

import (
	"context"
	"fmt"
	"xj/xapi-gateway/rpc_api"

	"dubbo.apache.org/dubbo-go/v3/common/logger"
	"github.com/gin-gonic/gin"
)

// 调用成功，接口调用次数+1
func InvokeCountMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("开始调用RPC: 更新接口调用次数")
		// 更新接口调用次数
		reply, err := grpcUserInterfaceInfoImpl.InvokeCount(context.Background(), &rpc_api.InvokeCountReq{
			UserId:      replyGetInvokeUser.Id,
			InterfaceId: replyGetFullUserInterfaceInfo.Id,
		})
		if err != nil {
			logger.Error(err)
		}
		logger.Infof("更新接口调用次数 InvokeCount get reply~~: %v\n", reply)
		fmt.Println("[middleware 调用次数统计更新]InvokeCountMiddleware complete!")
		c.Next()
	}
}
