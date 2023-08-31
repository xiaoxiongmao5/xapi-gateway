package middleware

import (
	"fmt"
	"xj/xapi-gateway/utils"

	"github.com/gin-gonic/gin"
)

// 添加请求日志
func LogMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("请求唯一标识：")
		fmt.Println("请求参数：", c.Params)         //  [{path /interface/invoke}]
		fmt.Println("请求路径URL：", c.Request.URL) //	/interface/invoke?token%20=%20123
		fmt.Println("请求方法Method：", c.Request.Method)
		// fmt.Println("请求参数Body：", c.Request.Body)
		// fmt.Println("请求参数RemoteAddr：", c.Request.RemoteAddr)
		fmt.Println("请求参数RequestURI：", c.Request.RequestURI) //	/interface/invoke?token%20=%20123
		fmt.Println("请求头Header：", c.Request.Header)
		fmt.Println("请求参数Cookies：", c.Request.Cookies())       //	[token=eyJhbGciOiJIUzxxx]
		fmt.Println("请求参数Host[目标地址]：", c.Request.Host)         //	localhost:8080
		fmt.Println("请求参数Referer[来源地址]：", c.Request.Referer()) //	http://localhost:8001/
		domain, _ := utils.GetDomainFromReferer(c.Request.Referer())
		fmt.Println("请求参数Referer[来源域名]：", domain)   //	localhost
		fmt.Println("访问IP：", utils.GetRequestIp(c)) //	127.0.0.1
		fmt.Println("本地IP：", utils.GetLocalIP())    //	192.168.2.104

	}
}
