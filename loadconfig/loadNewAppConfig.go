package loadconfig

import (
	"encoding/json"
	"os"
	"reflect"
	"sync"
	"time"
	gconfig "xj/xapi-gateway/g_config"
	glog "xj/xapi-gateway/g_log"
)

var appConfigMutex sync.Mutex

// 加载App配置数据
func LoadAppConfig() (*gconfig.AppConfiguration, error) {
	filePath := "conf/appconfig.json"
	config := &gconfig.AppConfiguration{}

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

// 加载App配置数据(可动态更新)
func LoadAppConfigDynamic() (*gconfig.AppConfigurationDynamic, error) {
	filePath := "conf/appdynamicconfig.json"
	config := &gconfig.AppConfigurationDynamic{}

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

func LoadNewAppDynamicConfig() {
	filePath := "conf/appdynamicconfig.json"
	ticker := time.NewTicker(3 * time.Second) // 每3秒检查一次配置文件
	defer ticker.Stop()

	var lastModTime time.Time
	var lastConfig *gconfig.AppConfigurationDynamic // 保存配置数据

	for range ticker.C {
		fileInfo, err := os.Stat(filePath)
		if err != nil {
			glog.Log.Errorf("Error reading config file: %v", err)
			continue
		}

		if fileInfo.ModTime() != lastModTime {
			lastModTime = fileInfo.ModTime()

			newConfig, err := LoadAppConfigDynamic()
			if err != nil {
				glog.Log.Errorf("Error loading config: %v", err)
				// todo 更新加载App配置数据失败，需报警
				continue
			}

			// 检查新配置与旧配置是否相同，避免不必要的重新加载
			appConfigMutex.Lock()
			if !reflect.DeepEqual(lastConfig, newConfig) {
				lastConfig = newConfig
				// 在这里使用最新的配置数据进行处理
				glog.Log.Errorf("Loaded new config: %+v", newConfig)
				gconfig.AppConfigDynamic = newConfig
			}
			appConfigMutex.Unlock()
		}
	}
}
