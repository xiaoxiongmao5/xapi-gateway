package main

import (
	"flag"
	"os"
	"xj/xapi-gateway/middleware"

	_ "dubbo.apache.org/dubbo-go/v3/imports"

	"github.com/gin-gonic/gin"
)

/** 使用网关做的事情（统一做的事情）

1. 添加请求日志✅
1. 路由转发✅
1. 访问控制（黑白名单）✅
1. 统一鉴权（AK SK）✅
1. 发布控制（灰度更新）
1. 流量染色
1. 统一业务处理（调用次数统计更新）
1. 接口保护（限制请求、信息脱敏、超时时间、降级熔断）
1. 统一日志
*/

func init() {
	// 使用命令行参数来指定配置文件路径
	configFile := flag.String("config", "conf/dubbogo.yml", "Path to Dubbo-go config file")
	flag.Parse()

	// 设置 DUBBO_GO_CONFIG_PATH 环境变量
	os.Setenv("DUBBO_GO_CONFIG_PATH", *configFile)
}

func main() {

	router := gin.Default()

	// 请求日志
	router.Use(middleware.LogMiddleware())
	// 处理跨域
	router.Use(middleware.CORSMiddleware())

	// 定义一个路由组，用于匹配以 / 开头的路由
	apiGroup := router.Group("/")
	{
		// 匹配这个路由组中的所有请求方式和路径片段，无论是GET、POST、DELETE 等方式，以及后面跟着什么路径片段，都会被这个路由组匹配到。
		apiGroup.Any("/*path",
			middleware.FilterWithAccessControl(), // 访问控制（黑白名单）
			middleware.SignMiddleware(),          // 统一鉴权（API权限验证）
			middleware.ValidInterfaceInfo(),      // 验证请求的接口是否存在
			middleware.ForwardApi(),              // 路由转发
			middleware.InvokeCountMiddleware(),   //调用次数统计更新
		)
	}

	// 运行服务
	router.Run(":8080")
	// s := &http.Server{
	// 	Addr:           ":8080",
	// 	Handler:        router,
	// 	ReadTimeout:    10 * time.Second,
	// 	WriteTimeout:   10 * time.Second,
	// 	MaxHeaderBytes: 1 << 20,
	// }
	// s.ListenAndServe()
}
