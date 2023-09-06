package middleware

import (
	"fmt"
	"time"
	"xj/xapi-gateway/utils"

	"github.com/gin-gonic/gin"
)

// 添加请求日志
func LogMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now()                //开始时间
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
		fmt.Println("LogMiddleware complete![请求日志]")
		c.Next()
		endTime := time.Now() //结束时间
		subtime := endTime.Sub(startTime)
		fmt.Printf("总耗时: %d 毫秒(%.2f 秒)\n", subtime.Milliseconds(), subtime.Seconds())
	}
}
