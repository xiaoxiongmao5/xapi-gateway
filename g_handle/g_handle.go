package ghandle

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// 状态未授权 401
func HandlerUnauthorized(c *gin.Context) {
	c.JSON(http.StatusUnauthorized, gin.H{"result": http.StatusUnauthorized, "msg": "状态未授权"})
	c.Abort()
}

// 禁止状态 403
func HandlerForbidden(c *gin.Context) {
	c.JSON(http.StatusForbidden, gin.H{"result": http.StatusForbidden, "msg": "禁止状态"})
	c.Abort()
}

// 从上下文信息获取信息失败 204
func HandlerContextError(c *gin.Context, key string) {
	c.JSON(http.StatusNoContent, gin.H{"result": http.StatusNoContent, "msg": "从上下文信息获取[" + key + "]失败"})
	c.Abort()
}

// 未知错误 500
func HandlerServerError(c *gin.Context) {
	c.JSON(http.StatusInternalServerError, gin.H{"result": http.StatusInternalServerError, "msg": "未知错误"})
	c.Abort()
}

// 业务成功 200
func HandlerSuccess(c *gin.Context, msg string, data any) {
	if data == nil {
		c.JSON(http.StatusOK, gin.H{"result": 0, "msg": msg})
		return
	}
	c.JSON(http.StatusOK, gin.H{"result": 0, "msg": msg, "data": data})
}

// Dubbo加载失败 424
func HandlerDobboLoadFailed(c *gin.Context, msg string) {
	c.JSON(http.StatusFailedDependency, gin.H{"result": http.StatusFailedDependency, "msg": "Dubbo加载失败: " + msg})
	c.Abort()
}

// 调用接口失败 400
func HandlerInvokeError(c *gin.Context) {
	c.JSON(http.StatusBadRequest, gin.H{"result": http.StatusBadRequest, "msg": "调用接口失败"})
	c.Abort()
}
