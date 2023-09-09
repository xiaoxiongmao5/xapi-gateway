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
		startTime := time.Now() //开始时间

		uniSessionId := utils.CreateUniSessionId()
		c.Set("uniSessionId", uniSessionId)

		fmt.Println("请求唯一ID: ", uniSessionId)               //675481600
		fmt.Println("请求路径参数: ", c.Params)                   //[{path /api/name}]
		fmt.Println("请求路径带参数: ", c.Request.URL)             // /api/name?name=xiaohua123
		fmt.Println("请求路径带参数: ", c.Request.RequestURI)      // /api/name?name=xiaohua123
		fmt.Println("请求方式", c.Request.Method)               //GET
		fmt.Println("请求头: ", c.Request.Header)              //map[Accept-Encoding:[gzip] Accesskey:[H6GxH5ERXL4zVZ3IJrs2EZBRO0CizHxMvDXrxbWVQmE=] Gateway_transdata:[2] Nonce:[3968843963288542780443340209407078562082975207184102505760380310447773194108299198709643218798055836] Sign:[faeeeaf3161abcdc46e46353551b63f0] Timestamp:[1693994604] User-Agent:[Go-http-client/1.1]]
		fmt.Println("请求Cookies: ", c.Request.Cookies())     //[token=eyJhbGciOiJIUzxxx]
		fmt.Println("请求目标地址Host: ", c.Request.Host)         //localhost:8080
		fmt.Println("请求来源地址Referer: ", c.Request.Referer()) //
		domain, err := utils.GetDomainFromReferer(c.Request.Referer())
		if err != nil {
			fmt.Println("获得请求来源域名Referer失败， err=: ", err.Error())
		}
		fmt.Println("请求来源域名Referer: ", domain)       //localhost
		fmt.Println("访问IP: ", utils.GetRequestIp(c)) //127.0.0.1
		fmt.Println("本机IP: ", utils.GetLocalIP())    //[192.168.2.104]
		// fmt.Println("请求参数RemoteAddr: ", c.Request.RemoteAddr)
		// fmt.Println("请求参数Body: ", c.Request.Body)
		fmt.Println("[middleware 请求日志]LogMiddleware complete!")

		c.Next()

		endTime := time.Now() //结束时间
		subtime := endTime.Sub(startTime)
		fmt.Printf("总耗时: %d 毫秒(%.2f 秒)\n", subtime.Milliseconds(), subtime.Seconds())
	}
}
