package ghandle

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// 状态未授权 401
func HandlerUnauthorized(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"result": http.StatusUnauthorized, "msg": "状态未授权"})
	c.Abort()
}

// 禁止状态 403
func HandlerForbidden(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"result": http.StatusForbidden, "msg": "禁止状态"})
	c.Abort()
}

// 从上下文信息获取信息失败 204
func HandlerContextError(c *gin.Context, key string) {
	c.JSON(http.StatusOK, gin.H{"result": http.StatusNoContent, "msg": "从上下文信息获取[" + key + "]失败"})
	c.Abort()
}

// 从RPC远程获取信息失败 204
func HandlerGetContextByRPCError(c *gin.Context, key string) {
	c.JSON(http.StatusOK, gin.H{"result": http.StatusNoContent, "msg": "RPC获取[" + key + "]失败"})
	c.Abort()
}

// 未知错误 500
func HandlerServerError(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"result": http.StatusInternalServerError, "msg": "未知错误"})
	c.Abort()
}

// Dubbo加载失败 424
func HandlerDobboLoadFailed(c *gin.Context, msg string) {
	c.JSON(http.StatusOK, gin.H{"result": http.StatusFailedDependency, "msg": "Dubbo加载失败: " + msg})
	c.Abort()
}

// ——————————————————————以下是业务——————————————————————

// 业务成功 200
func HandlerSuccess(c *gin.Context, msg string, data any) {
	if data == nil {
		c.JSON(http.StatusOK, gin.H{"result": 0, "msg": msg})
		return
	}
	c.JSON(http.StatusOK, gin.H{"result": 0, "msg": msg, "data": data})
}

// 参数错误 200
func HandlerParamError(c *gin.Context, paramName string) {
	if paramName == "" {
		c.JSON(http.StatusOK, gin.H{"result": 1, "msg": "参数错误"})
	} else {
		c.JSON(http.StatusOK, gin.H{"result": 1, "msg": "参数[" + paramName + "]错误"})
	}
	c.Abort()
}

// 校验接口失败 200
func HandlerValidInterfaceFailed(c *gin.Context, msg string) {
	c.JSON(http.StatusOK, gin.H{"result": 2, "msg": msg})
	c.Abort()
}

// 第三方API接口调用失败,请联系管理员 -1
func HandlerInvokeError(c *gin.Context) {
	// http.StatusBadRequest
	c.JSON(http.StatusOK, gin.H{"result": -1, "msg": "第三方API接口调用失败,请联系管理员!"})
	c.Abort()
}

// 第三方API接口调用超时,请联系管理员 -2
func HandlerInvokeTimeout(c *gin.Context) {
	// http.StatusBadRequest
	c.JSON(http.StatusOK, gin.H{"result": -2, "msg": "第三方API接口调用超时,请联系管理员!"})
	c.Abort()
}
