package middleware

import (
	"fmt"
	"io"
	"net/http"
	ghandle "xj/xapi-gateway/g_handle"
	"xj/xapi-gateway/utils"

	"github.com/gin-gonic/gin"
)

// 路由转发
func ForwardApi() gin.HandlerFunc {
	return func(c *gin.Context) {
		targetURL := replyGetFullUserInterfaceInfo.Host + replyGetFullUserInterfaceInfo.Url

		queryString := c.Request.URL.RawQuery
		if queryString != "" {
			targetURL += "?" + queryString
		}

		// 判断请求方式是否一致
		if !utils.CheckSameStrFold("校验: 请求方式method一致", replyGetFullUserInterfaceInfo.Method, c.Request.Method) {
			ghandle.HandlerForbidden(c)
			return
		}

		// 创建转发请求
		request, err := http.NewRequest(c.Request.Method, targetURL, c.Request.Body)
		if err != nil {
			fmt.Println("error", "Failed to create request")
			ghandle.HandlerInvokeError(c)
			return
		}

		// 设置请求头
		for key, values := range c.Request.Header {
			for _, value := range values {
				request.Header.Add(key, value)
			}
		}
		// todo 这里可删可不删
		// 如果删除，则最好加上流量染色(为了溯源)，不过有请求ip说明是网关发出的，也可以溯源
		// 如果不删，思考是否会要求第三方服务添加跨域问题（好像也不会，跨域需要从前端发出才会有问题），不删的话，第三方也可以自行校验，也不错
		request.Header.Del("accessKey")
		request.Header.Del("nonce")
		request.Header.Del("timestamp")
		request.Header.Del("sign")
		request.Header.Del("gateway_transdata")
		uniSessionId, exists := c.Get("uniSessionId")
		if !exists {
			ghandle.HandlerContextError(c, "uniSessionId")
			return
		}
		request.Header.Add("from", "xapi-gateway")
		request.Header.Add("from_sid", uniSessionId.(string))

		// 发起转发请求
		client := &http.Client{}
		response, err := client.Do(request)
		if err != nil {
			fmt.Println("error", "Failed to forward request")
			ghandle.HandlerInvokeError(c)
			return
		}
		defer response.Body.Close()

		// 读取转发请求的响应内容
		body, err := io.ReadAll(response.Body)
		if err != nil {
			fmt.Println("error", "Failed to read response")
			ghandle.HandlerInvokeError(c)
			return
		}

		// 添加响应日志
		fmt.Println("响应码: ", response.StatusCode)

		// 调用成功
		if response.StatusCode == http.StatusOK {
			// 返回响应内容给请求方
			c.String(response.StatusCode, string(body))
			fmt.Println("ForwardApi complete![路由转发]")
			c.Next()
		} else {
			ghandle.HandlerInvokeError(c)
			return
		}
	}
}
