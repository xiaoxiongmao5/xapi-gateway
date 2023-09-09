package gconfig

import (
	"encoding/json"
	"log"
	"os"
	"reflect"
	"sync"
	"time"
)

// App配置数据
type AppConfiguration struct {
	Server struct {
		Port int `json:"port"`
	} `json:"server"`
	IPWhiteList []string `json:"ipWhiteList"`
}

var (
	AppConfigMutex sync.Mutex
	AppConfig      *AppConfiguration
)

// 加载App配置数据
func LoadAppConfig(filePath string) (*AppConfiguration, error) {
	config := &AppConfiguration{}

	// 打开项目配置文件
	configFile, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer configFile.Close()

	// 解码配置文件内容到结构体
	jsonParser := json.NewDecoder(configFile)
	if err := jsonParser.Decode(config); err != nil {
		return nil, err
	}

	return config, nil
}

func LoadNewAppConfig(configFilePath string) {
	ticker := time.NewTicker(3 * time.Second) // 每3秒检查一次配置文件
	defer ticker.Stop()

	var lastModTime time.Time
	var lastConfig *AppConfiguration // 保存配置数据

	for range ticker.C {
		fileInfo, err := os.Stat(configFilePath)
		if err != nil {
			log.Printf("Error reading config file: %v\n", err)
			continue
		}

		if fileInfo.ModTime() != lastModTime {
			lastModTime = fileInfo.ModTime()

			newConfig, err := LoadAppConfig(configFilePath)
			if err != nil {
				log.Printf("Error loading config: %v\n", err)
				// todo 更新加载App配置数据失败，需报警
				continue
			}

			// 检查新配置与旧配置是否相同，避免不必要的重新加载
			AppConfigMutex.Lock()
			if !reflect.DeepEqual(lastConfig, newConfig) {
				lastConfig = newConfig
				// 在这里使用最新的配置数据进行处理
				log.Printf("Loaded new config: %+v\n", newConfig)
				AppConfig = newConfig
			}
			AppConfigMutex.Unlock()
		}
	}
}
