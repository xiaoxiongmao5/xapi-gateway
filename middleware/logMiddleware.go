package middleware

import (
	"fmt"
	"time"
	glog "xj/xapi-gateway/g_log"
	"xj/xapi-gateway/utils"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type AppHook struct {
	RequestID string
}

func (h *AppHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

func (h *AppHook) Fire(entry *logrus.Entry) error {
	entry.Data["req_id"] = h.RequestID
	return nil
}

func LogMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now()                 // 开始时间
		requestID := utils.CreateUniSessionId() //1365038848
		c.Set("uniSessionId", requestID)

		// 添加reqId到每次请求的日志中
		h := &AppHook{RequestID: requestID}
		glog.Log.AddHook(h)

		// 将标识添加到响应头中，以便客户端可以获取它
		c.Writer.Header().Set("X-Request-ID", requestID)

		// 记录请求信息
		domain, err := utils.GetDomainFromReferer(c.Request.Referer())
		if err != nil {
			glog.Log.Error("获得请求来源域名Referer失败, err=: ", err.Error())
		}
		localIp, err := utils.GetLocalIP()
		if err != nil {
			glog.Log.Error("获得本机IP失败, err=: ", err.Error())
		}
		// fmt.Println("请求路径参数: ", c.Params)       //[{path /api/name}]
		// fmt.Println("请求路径带参数: ", c.Request.URL) // /api/name?name=xiaohua123
		// fmt.Println("请求来源地址Referer: ", c.Request.Referer()) //

		glog.Log.WithFields(logrus.Fields{
			"请求路径":   c.Request.RequestURI,
			"请求方式":   c.Request.Method,
			"目标Host": c.Request.Host,
			"来源域名":   domain,
			"来源IP":   utils.GetRequestIp(c),
			"本机IP":   localIp,
		}).Info("请求日志")

		glog.Log.WithFields(logrus.Fields{
			"请求头":     c.Request.Header,
			"Cookies": c.Request.Cookies(),
		}).Info("请求日志补充")
		// glog.Log.Infof("Received 请求路径: %s, 请求方式: %s, 目标Host: %s, 来源域名: %s, 来源IP: %s", c.Request.RequestURI, c.Request.Method, c.Request.Host, domain, utils.GetRequestIp(c))

		c.Next() // 处理请求

		endTime := time.Now() // 结束时间
		// latencyTm := time.Since(startTime)

		// 记录响应信息
		respStatus := c.Writer.Status()
		latencyTm := endTime.Sub(startTime) // 执行时间总耗时
		totaltm := ""
		if latencyTm < 1*time.Millisecond {
			totaltm = fmt.Sprintf("%dµs", latencyTm.Microseconds()) //微秒 1微秒 = 1000纳秒
		} else if latencyTm < 1*time.Second {
			totaltm = fmt.Sprintf("%dms", latencyTm.Milliseconds()) //毫秒 1毫秒 = 1000微秒
		} else {
			totaltm = fmt.Sprintf("%.2fs", latencyTm.Seconds()) //秒 1秒 = 1000毫秒
		}
		glog.Log.WithFields(logrus.Fields{
			"响应码": respStatus,
			"总耗时": totaltm,
		}).Info("响应日志")
		// glog.Log.Infof("响应码: %d, 总耗时: %s", respStatus, totaltm)
	}
}
