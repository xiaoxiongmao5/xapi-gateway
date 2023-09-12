package utils

import (
	"fmt"
	"reflect"
	"strings"
	glog "xj/xapi-gateway/g_log"

	"github.com/sirupsen/logrus"
)

// 检查是否为空字符串
func AreEmptyStrings(values ...string) bool {
	for _, value := range values {
		if value == "" {
			return true
		}
	}
	return false
}

// 检查是否一样（使用 == 检查）
func CheckSame[T string | int](desc string, str1 T, str2 T) bool {
	res := false
	if str1 == str2 {
		res = true
	} else {
		res = false
	}
	glog.Log.WithFields(logrus.Fields{
		"got":    str1,
		"export": str2,
		"pass":   res,
	}).Info(desc)
	return res
}

// 检查字符串忽略大小写后是否一样（使用 EqualFold 检查）
func CheckSameStrFold(desc string, str1 string, str2 string) bool {
	res := false
	if strings.EqualFold(str1, str2) {
		res = true
	} else {
		res = false
	}
	glog.Log.WithFields(logrus.Fields{
		"got":    str1,
		"export": str2,
		"pass":   res,
		"notes":  "已忽略大小写",
	}).Info(desc)
	return res
}

// 检查数组是否一样（使用 DeepEqual 检查）
func CheckSameArr[T string | int | []int](desc string, str1 T, str2 T) bool {
	res := false
	if reflect.DeepEqual(str1, str2) {
		res = true
	} else {
		res = false
	}
	glog.Log.WithFields(logrus.Fields{
		"got":    fmt.Sprintf("%v", str1),
		"export": fmt.Sprintf("%v", str2),
		"pass":   res,
	}).Info(desc)
	return res
}
