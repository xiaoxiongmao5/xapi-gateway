package middleware

import (
	"context"
	glog "xj/xapi-gateway/g_log"
	"xj/xapi-gateway/rpc_api"

	"dubbo.apache.org/dubbo-go/v3/common/logger"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// 调用成功，接口调用次数+1
func InvokeCountMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		glog.Log.Info("开始调用RPC: 更新接口调用次数")
		// 更新接口调用次数
		reply, err := grpcUserInterfaceInfoImpl.InvokeCount(context.Background(), &rpc_api.InvokeCountReq{
			UserId:      replyGetInvokeUser.Id,
			InterfaceId: replyGetFullUserInterfaceInfo.Id,
		})
		if err != nil {
			// todo 报警，但不影响用户剩余业务
			logger.Error(err)
		}
		glog.Log.Infof("更新接口调用次数 InvokeCount get reply~~: %v", reply)

		glog.Log.WithFields(logrus.Fields{
			"pass": true,
		}).Info("middleware-调用次数统计更新")

		c.Next()
	}
}
