package gconfig

import (
	"sync"
)

// App配置数据
type AppConfiguration struct {
	Server struct {
		Port int `json:"port"`
	} `json:"server"`
	IPWhiteList []string `json:"ipWhiteList"`
	IPBlackList []string `json:"ipBlackList"`
}

var (
	AppConfigMutex sync.Mutex
	AppConfig      *AppConfiguration
)
