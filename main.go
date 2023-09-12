package main

import (
	"flag"
	"fmt"
	"os"
	gconfig "xj/xapi-gateway/g_config"
	glog "xj/xapi-gateway/g_log"
	"xj/xapi-gateway/loadconfig"
	"xj/xapi-gateway/middleware"

	_ "dubbo.apache.org/dubbo-go/v3/imports"

	"github.com/gin-gonic/gin"
)

func init() {
	// 实例化日志对象
	if logger, err := glog.SetupLogger(); err != nil {
		panic(err)
	} else {
		glog.Log = logger
	}

	// 使用命令行参数来指定配置文件路径
	configFile := flag.String("config", "conf/dubbogo.yaml", "Path to Dubbo-go config file")
	flag.Parse()

	// 设置 DUBBO_GO_CONFIG_PATH 环境变量
	os.Setenv("DUBBO_GO_CONFIG_PATH", *configFile)
}

func main() {
	defer glog.Log.Writer().Close()

	// 加载App配置数据
	if config, err := loadconfig.LoadAppConfig(); err != nil {
		glog.Log.Error("App配置加载失败, err=:", err)
		panic(err)
	} else {
		glog.Log.Info("App配置加载成功")
		gconfig.AppConfig = config
	}

	// 启动配置文件加载协程
	go loadconfig.LoadNewAppConfig()

	r := gin.Default()

	// 使用中间件格式化日志
	r.Use(middleware.LogMiddleware())

	// 使用中间件来处理跨域请求，并允许携带 Cookie
	r.Use(middleware.CORSMiddleware())

	// 访问控制（黑名单）
	r.Use(middleware.FilterWithAccessControlInBlackIp())

	// 定义一个路由组，用于匹配以 / 开头的路由
	apiGroup := r.Group("/")
	{
		// 匹配这个路由组中的所有请求方式和路径片段，无论是GET、POST、DELETE 等方式，以及后面跟着什么路径片段，都会被这个路由组匹配到。
		apiGroup.Any("/*path",
			middleware.LoadGrpcImpl(),           // 加载GrpcImpl
			middleware.SignMiddleware(),         // 统一鉴权（API权限验证）
			middleware.ValidUserInterfaceInfo(), // 验证请求的接口是否允许被该用户使用
			middleware.ForwardApi(),             // 路由转发
			middleware.InvokeCountMiddleware(),  // 调用次数统计更新
		)
	}

	port := fmt.Sprintf(":%d", gconfig.AppConfig.Server.Port)
	r.Run(port)
	// s := &http.Server{
	// 	Addr:           ":8080",
	// 	Handler:        router,
	// 	ReadTimeout:    10 * time.Second,
	// 	WriteTimeout:   10 * time.Second,
	// 	MaxHeaderBytes: 1 << 20,
	// }
	// s.ListenAndServe()
}
