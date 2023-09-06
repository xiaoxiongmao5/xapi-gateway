package test

import (
	"context"
	"flag"
	"fmt"
	"os"
	"xj/xapi-gateway/rpc_api"

	_ "dubbo.apache.org/dubbo-go/v3/imports"

	"dubbo.apache.org/dubbo-go/v3/config"
)

func beforeInit() {
	// 使用命令行参数来指定配置文件路径
	configFile := flag.String("config", "../conf/dubbogo.yaml", "Path to Dubbo-go config file")
	flag.Parse()

	// 设置 DUBBO_GO_CONFIG_PATH 环境变量
	os.Setenv("DUBBO_GO_CONFIG_PATH", *configFile)
}

var grpcUserInfoImpl = new(rpc_api.UserInfoClientImpl)
var grpcInterfaceInfoImpl = new(rpc_api.IntefaceInfoClientImpl)
var grpcUserInterfaceInfoImpl = new(rpc_api.UserIntefaceInfoClientImpl)

func RunClientImpl() (res bool, err error) {
	beforeInit()
	config.SetConsumerService(grpcUserInfoImpl)
	config.SetConsumerService(grpcInterfaceInfoImpl)
	config.SetConsumerService(grpcUserInterfaceInfoImpl)
	if err = config.Load(); err != nil {
		return
	}
	accessKey := "H6GxH5ERXL4zVZ3IJrs2EZBRO0CizHxMvDXrxbWVQmE="
	reply1, err := grpcUserInfoImpl.GetInvokeUser(context.Background(), &rpc_api.GetInvokeUserReq{AccessKey: accessKey})
	if err != nil {
		return
	} else {
		fmt.Println("======================================")
		fmt.Println("GetInvokeUser reply1=", reply1)
		fmt.Println("---------------------------------------------------------")
	}

	var interfaceId int64
	interfaceId = 3
	reply2, err := grpcInterfaceInfoImpl.GetInterfaceInfoById(context.Background(), &rpc_api.GetInterfaceInfoByIdReq{InterfaceId: interfaceId})
	if err != nil {
		return
	} else {
		fmt.Println("======================================")
		fmt.Println("GetInterfaceInfoById reply2=", reply2)
		fmt.Println("---------------------------------------------------------")
	}

	reply3, err := grpcUserInterfaceInfoImpl.InvokeCount(context.Background(), &rpc_api.InvokeCountReq{InterfaceId: interfaceId, UserId: reply1.Id})
	if err != nil {
		return
	} else {
		fmt.Println("======================================")
		fmt.Printf("InvokeCount reply3=%v\n", reply3)
		fmt.Println("---------------------------------------------------------")
	}
	return true, nil
}
