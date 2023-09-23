package gconfig

// App配置数据
type AppConfiguration struct {
	Server struct {
		Port int `json:"port"`
	} `json:"server"`
}

// App配置数据(可动态更新)
type AppConfigurationDynamic struct {
	IPWhiteList     []string `json:"ipWhiteList"`
	IPBlackList     []string `json:"ipBlackList"`
	IPAdminList     []string `json:"ipAdminList"`
	RateLimitConfig struct {
		RequestsPerSecond float64 `json:"requests_per_second"`
		BucketSize        int     `json:"bucket_size"`
	} `json:"rateLimitConfig"`
	InvokeTimeOut int `json:"invokeTimeOut"`
}

var (
	AppConfig        *AppConfiguration
	AppConfigDynamic *AppConfigurationDynamic
)
