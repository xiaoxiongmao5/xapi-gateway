package loadconfig

import (
	"encoding/json"
	"os"
	"reflect"
	"time"
	gconfig "xj/xapi-gateway/g_config"
	glog "xj/xapi-gateway/g_log"
)

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

func LoadNewAppConfig() {
	filePath := "conf/appconfig.json"
	ticker := time.NewTicker(3 * time.Second) // 每3秒检查一次配置文件
	defer ticker.Stop()

	var lastModTime time.Time
	var lastConfig *gconfig.AppConfiguration // 保存配置数据

	for range ticker.C {
		fileInfo, err := os.Stat(filePath)
		if err != nil {
			glog.Log.Errorf("Error reading config file: %v", err)
			continue
		}

		if fileInfo.ModTime() != lastModTime {
			lastModTime = fileInfo.ModTime()

			newConfig, err := LoadAppConfig()
			if err != nil {
				glog.Log.Errorf("Error loading config: %v", err)
				// todo 更新加载App配置数据失败，需报警
				continue
			}

			// 检查新配置与旧配置是否相同，避免不必要的重新加载
			gconfig.AppConfigMutex.Lock()
			if !reflect.DeepEqual(lastConfig, newConfig) {
				lastConfig = newConfig
				// 在这里使用最新的配置数据进行处理
				glog.Log.Errorf("Loaded new config: %+v", newConfig)
				gconfig.AppConfig = newConfig
			}
			gconfig.AppConfigMutex.Unlock()
		}
	}
}
