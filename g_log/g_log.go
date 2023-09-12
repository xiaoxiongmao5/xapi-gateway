package glog

import (
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/sirupsen/logrus"
)

// 全局的 logrus.Logger 实例，用于应用中的所有日志记录
var Log *logrus.Logger

// LogConfig 包含日志配置信息
type LogConfig struct {
	LogFilePath  string // 日志文件路径
	LogLevel     string // 日志级别
	MaxAge       int    // 最大保存天数
	RotationTime int    // 日志切割时间间隔（小时）
}

func SetupLogger() (*logrus.Logger, error) {
	config := LogConfig{
		LogFilePath:  "log/app.log",
		LogLevel:     "info",
		MaxAge:       7,  // 7天
		RotationTime: 24, // 24小时
	}

	logger := logrus.New()

	// 设置日志级别
	logLevel, err := logrus.ParseLevel(config.LogLevel)
	if err != nil {
		logLevel = logrus.DebugLevel
	}
	logger.SetLevel(logLevel)

	// 设置日志输出
	logWriter, err := rotatelogs.New(
		// 分割后的文件名称
		config.LogFilePath+".%Y%m%d.log",
		// 生成软链，指向最新日志文件
		rotatelogs.WithLinkName(config.LogFilePath),
		// 设置最大保存时间(7天)
		rotatelogs.WithMaxAge(time.Duration(config.MaxAge)*24*time.Hour),
		// 设置日志切割时间间隔(1天)
		rotatelogs.WithRotationTime(time.Duration(config.RotationTime)*time.Hour),
	)
	if err != nil {
		logger.Fatal("Failed to create log file: ", err)
		return nil, err
	}

	logger.SetOutput(logWriter)
	logger.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})
	// logger.SetFormatter(&logrus.TextFormatter{
	// 	TimestampFormat: "2006-01-02 15:04:05",
	// })

	return logger, nil
}
