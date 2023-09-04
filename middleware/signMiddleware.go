package middleware

import (
	"context"
	"fmt"
	"strconv"
	ghandle "xj/xapi-gateway/g_handle"
	"xj/xapi-gateway/rpc_api"
	"xj/xapi-gateway/utils"

	"dubbo.apache.org/dubbo-go/v3/common/logger"
	dubboConfig "dubbo.apache.org/dubbo-go/v3/config"
	"github.com/gin-gonic/gin"
)

var grpcUserInfoImpl = new(rpc_api.UserInfoClientImpl)

func SignMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		dubboConfig.SetConsumerService(grpcUserInfoImpl)
		if err := dubboConfig.Load(); err != nil {
			panic(err)
		}
		// 从请求头中获取参数
		accessKey := c.Request.Header.Get("accessKey")
		nonce := c.Request.Header.Get("nonce")
		timestamp := c.Request.Header.Get("timestamp")
		sign := c.Request.Header.Get("sign")
		body := ""

		// 去数据库中查是否已分配给用户
		reply, err := grpcUserInfoImpl.GetInvokeUser(context.Background(), &rpc_api.GetInvokeUserReq{AccessKey: accessKey})
		if err != nil {
			ghandle.HandlerNoAuth(c)
			return
		}
		c.Set("cur_userinfo", reply)
		logger.Infof("GetInvokeUser get reply~~: %v\n", reply)

		if accessKey != reply.Accesskey {
			ghandle.HandlerNoAuth(c)
			return
		}

		// 校验如果随机数大于1万，则抛出异常
		if num, err := strconv.Atoi(nonce); err != nil || num > 10000 {
			ghandle.HandlerNoAuth(c)
			return
		}
		// 时间和当前时间不能超过5分钟
		fiveMinutes := int64(5 * 60)
		timestampNow := utils.GetCurrentTimeMillis()
		if tsInt, err := strconv.ParseInt(timestamp, 10, 64); err != nil || timestampNow-tsInt >= fiveMinutes {
			ghandle.HandlerNoAuth(c)
			return
		}
		// 如果生成的签名不一致，则抛出异常
		str := accessKey + nonce + timestamp + body
		signEcrypt := utils.GetAPISign(str, reply.Secretkey)
		if signEcrypt != sign {
			ghandle.HandlerNoAuth(c)
			return
		}
		fmt.Println("API权限验证 通过")
		c.Next()
	}
}
