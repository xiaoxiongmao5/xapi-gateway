package middleware

import (
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 路由转发
func ForwardApi() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 修改目标域名
		targetDomain := "http://127.0.0.1:8090"
		// 构建新的目标URL
		targetURL := targetDomain + c.Param("path")
		queryString := c.Request.URL.RawQuery
		if queryString != "" {
			targetURL += "?" + queryString
		}

		// 获取转发请求方式参数（?forward_method=POST）
		// forwardMethod := c.DefaultQuery("forward_method", "GET")
		forwardMethod := c.Request.Method
		// 创建转发请求
		request, err := http.NewRequest(forwardMethod, targetURL, c.Request.Body)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create request"})
			return
		}

		// 设置请求头
		for key, values := range c.Request.Header {
			for _, value := range values {
				request.Header.Add(key, value)
			}
		}

		// 发起转发请求
		client := &http.Client{}
		response, err := client.Do(request)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to forward request"})
			return
		}
		defer response.Body.Close()

		// 读取转发请求的响应内容
		body, err := io.ReadAll(response.Body)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read response"})
			return
		}

		// 返回响应内容给请求方
		c.String(response.StatusCode, string(body))
	}
}
