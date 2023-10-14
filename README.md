# xapi-gateway (API开放平台-网关)

## 项目的核心业务

使用网关做的事情（统一做的事情）

1. 统一日志（添加请求日志、响应日志）
2. 处理跨域
1. 接口保护（限制请求、信息脱敏、超时时间、降级熔断）
1. 发布控制（灰度更新）
3. 访问控制（黑白名单）
4. 统一鉴权（API签名认证）
1. 流量染色
5. 路由转发
1. 统一业务处理（验证请求接口是否存在、调用次数统计更新）

## 项目本地启动

⚠️ 注意：项目内使用了rpc远程调用，依赖 注册中心已启动、接口提供方已启动(具体见下面《关于 RPC 远程调用》的说明。

```cmd
go mod tidy
go run main.go
```

## 运行项目中的单元测试

```bash
go test -v ./test
go clean -testcache //清除测试缓存
```

## 关于 RPC 远程调用

该项目内的部分业务使用了dubbo-go 框架的rpc远程调用模式。

* 该项目角色是调用方（Consumer），依赖的提供方（Provide）是[xapi-backend 项目](https://github.com/xiaoxiongmao5/xapi-backend)

* 配置文件位置：/conf/dubbogo.yaml

* 具体业务为为：
    * 获得用户信息 `GetInvokeUser`
    * 获得接口信息 `GetInterfaceInfoByIdReq`
    * 更新接口调用次数 `InvokeCount`

