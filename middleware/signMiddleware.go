package middleware

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strconv"
	"time"
	ghandle "xj/xapi-gateway/g_handle"
	"xj/xapi-gateway/rpc_api"
	"xj/xapi-gateway/utils"

	"dubbo.apache.org/dubbo-go/v3/common/logger"
	"github.com/gin-gonic/gin"
)

func SignMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 从请求头中获取参数
		accessKey := c.Request.Header.Get("accessKey")
		nonce := c.Request.Header.Get("nonce")
		timestamp := c.Request.Header.Get("timestamp")
		sign := c.Request.Header.Get("sign")
		body := ""

		fmt.Println("开始调用RPC: 获得用户信息")
		// 去数据库中查是否已分配给用户
		reply, err := grpcUserInfoImpl.GetInvokeUser(context.Background(), &rpc_api.GetInvokeUserReq{AccessKey: accessKey})
		if err != nil {
			ghandle.HandlerUnauthorized(c)
			return
		}
		logger.Infof("GetInvokeUser get reply~~: %v\n", reply)
		replyGetInvokeUser = reply

		if !utils.CheckSame[string]("校验: accessKey一致", accessKey, reply.Accesskey) {
			ghandle.HandlerUnauthorized(c)
			return
		}

		// todo 全局缓存随机数，标记是否使用过，5分钟后失效
		// if num, err := strconv.Atoi(nonce); err != nil || num > 10000 {
		// 	ghandle.HandlerUnauthorized(c)
		// 	return
		// }

		// 时间和当前时间不能超过5分钟
		fiveMinutes := int64(5 * 60)
		timestampNow := time.Now().Unix()
		if tsInt, err := strconv.ParseInt(timestamp, 10, 64); err != nil || timestampNow-tsInt >= fiveMinutes {
			fmt.Println("时间戳校验失败, 已超时5分钟")
			ghandle.HandlerUnauthorized(c)
			return
		}
		// 如果生成的签名不一致，则抛出异常
		signature := calculateSignature(accessKey, reply.Secretkey, nonce, timestamp, body)
		if !utils.CheckSame[string]("校验: 签名一致", signature, sign) {
			ghandle.HandlerUnauthorized(c)
			return
		}
		fmt.Println("[middleware 统一鉴权（API权限验证）]SignMiddleware complete!")
		c.Next()
	}
}

// 计算API签名
func calculateSignature(accessKey, secretKey, nonce, timestamp, requestBody string) string {
	// 将参数拼接成一个字符串
	concatenatedString := accessKey + nonce + timestamp + requestBody + secretKey

	// 计算 MD5 值
	signature := md5.Sum([]byte(concatenatedString))
	return hex.EncodeToString(signature[:])
}
