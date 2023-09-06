package utils

import (
	"fmt"
	"net"
	"net/url"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// 生成唯一的sessionID，用于溯源每个请求日志
func CreateUniSessionId() string {
	unixNaco := time.Now().UnixNano() // 获取当前时间的Unix纳秒时间戳
	unixNaco = (unixNaco * 100000) & 0x7FFFFFFF
	return strconv.FormatInt(unixNaco, 10)
}

func GetRequestIp(c *gin.Context) string {
	reqIp := c.ClientIP()
	if reqIp == "::1" {
		reqIp = "127.0.0.1"
	}
	return reqIp
}

func GetLocalIP() []string {
	var ipStr []string
	netInterfaces, err := net.Interfaces()
	if err != nil {
		fmt.Println("net.Interfaces error:", err.Error())
		return ipStr
	}

	for i := 0; i < len(netInterfaces); i++ {
		if (netInterfaces[i].Flags & net.FlagUp) != 0 {
			addrs, _ := netInterfaces[i].Addrs()
			for _, address := range addrs {
				if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
					//获取IPv6
					/*if ipnet.IP.To16() != nil {
						fmt.Println(ipnet.IP.String())
						ipStr = append(ipStr, ipnet.IP.String())
					}*/
					//获取IPv4
					if ipnet.IP.To4() != nil {
						// fmt.Println(ipnet.IP.String())
						ipStr = append(ipStr, ipnet.IP.String())
					}
				}
			}
		}
	}
	return ipStr
}

func GetDomainFromReferer(referer string) (string, error) {
	parsedURL, err := url.Parse(referer)
	if err != nil {
		return "", err
	}

	return parsedURL.Hostname(), nil
}
