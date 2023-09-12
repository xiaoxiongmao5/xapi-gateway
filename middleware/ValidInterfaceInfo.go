package middleware

import (
	"context"
	"xj/xapi-gateway/enums"
	ghandle "xj/xapi-gateway/g_handle"
	glog "xj/xapi-gateway/g_log"
	"xj/xapi-gateway/rpc_api"
	"xj/xapi-gateway/utils"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func ValidUserInterfaceInfo() gin.HandlerFunc {
	return func(c *gin.Context) {
		interfaceIdstr := c.Request.Header.Get("gateway_transdata")
		interfaceId, err := utils.String2Int64(interfaceIdstr)
		if err != nil {
			ghandle.HandlerParamError(c, "gateway_transdata")
			c.Abort()
			return
		}
		glog.Log.Info("开始调用RPC: 获得用户接口信息, 接口ID=", interfaceId)

		reply, err := grpcUserInterfaceInfoImpl.GetFullUserInterfaceInfo(context.Background(), &rpc_api.GetFullUserInterfaceInfoReq{InterfaceId: interfaceId, UserId: replyGetInvokeUser.Id})

		if err != nil {
			ghandle.HandlerValidInterfaceFailed(c, "接口剩余可调用次数不足")
			c.Abort()
			return
		}
		glog.Log.Infof("get reply~~: %v", reply)
		replyGetFullUserInterfaceInfo = reply

		// 检查接口剩余可调用次数
		if reply.Leftnum <= 0 {
			ghandle.HandlerValidInterfaceFailed(c, "接口剩余可调用次数不足")
			c.Abort()
			return
		}
		// 检查用户调用该接口是否被禁用
		if reply.Banstatus != enums.UserInterfaceStatusOk {
			ghandle.HandlerValidInterfaceFailed(c, "该接口为禁用状态")
			c.Abort()
			return
		}
		// 检查接口是否正常状态
		if reply.Status != enums.InterfaceStatusOnline {
			ghandle.HandlerValidInterfaceFailed(c, "接口未上线")
			c.Abort()
			return
		}

		glog.Log.WithFields(logrus.Fields{
			"pass": true,
		}).Info("middleware-验证请求的接口是否允许被该用户使用")

		c.Next()
	}
}
