package middleware

import (
	gconfig "xj/xapi-gateway/g_config"
	ghandle "xj/xapi-gateway/g_handle"
	glog "xj/xapi-gateway/g_log"
	"xj/xapi-gateway/utils"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// 访问控制（黑名单）
func FilterWithAccessControlInBlackIp() gin.HandlerFunc {
	return func(c *gin.Context) {
		IP_BLACK_LIST := gconfig.AppConfigDynamic.IPBlackList
		requestIP := utils.GetRequestIp(c)
		flag := true
		for _, val := range IP_BLACK_LIST {
			if val == requestIP {
				flag = false
			}
		}

		glog.Log.WithFields(logrus.Fields{
			"requestIP": requestIP,
			"pass":      flag,
		}).Info("middleware-访问控制-ip黑名单")

		if flag {
			c.Next()
		} else {
			ghandle.HandlerForbidden(c)
			return
		}
	}
}

// 访问控制（白名单）
func FilterWithAccessControlInWhiteIp() gin.HandlerFunc {
	return func(c *gin.Context) {
		IP_WHITE_LIST := gconfig.AppConfigDynamic.IPWhiteList
		requestIP := utils.GetRequestIp(c)
		flag := false
		for _, val := range IP_WHITE_LIST {
			if val == requestIP {
				flag = true
			}
		}

		glog.Log.WithFields(logrus.Fields{
			"requestIP": requestIP,
			"pass":      flag,
		}).Info("middleware-访问控制-ip白名单")

		if flag {
			c.Next()
		} else {
			ghandle.HandlerForbidden(c)
			return
		}
	}
}

// 访问控制（admin管理员ip）
func FilterWithAccessControlInAdminIp() gin.HandlerFunc {
	return func(c *gin.Context) {
		IP_LIST := gconfig.AppConfigDynamic.IPAdminList
		requestIP := utils.GetRequestIp(c)
		flag := false
		for _, val := range IP_LIST {
			if val == requestIP {
				flag = true
			}
		}

		glog.Log.WithFields(logrus.Fields{
			"requestIP": requestIP,
			"pass":      flag,
		}).Info("middleware-访问控制-ip管理员admin")

		if flag {
			c.Next()
		} else {
			ghandle.HandlerForbidden(c)
			return
		}
	}
}
