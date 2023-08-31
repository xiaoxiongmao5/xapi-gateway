package utils

import "time"

// 获得当前时间的毫秒级时间戳
func GetCurrentTimeMillis() int64 {
	unixTimestampNano := time.Now().UnixNano()
	return unixTimestampNano / int64(time.Millisecond)
}

// 获取当前时间的秒级时间戳
func GetCurrentTimeUnixSecond() int64 {
	return time.Now().Unix()
}

// 获取当前时间的纳秒级时间戳
func GetCurrentTimeUnixNano() int64 {
	return time.Now().UnixNano()
}
