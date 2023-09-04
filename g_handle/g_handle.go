package ghandle

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandlerNoAuth(c *gin.Context) {
	c.JSON(http.StatusForbidden, gin.H{"error": "无权限"})
	c.Abort()
}

func HandlerContextError(c *gin.Context, key string) {
	c.JSON(http.StatusInternalServerError, gin.H{"error": "从上下文信息获取[" + key + "]失败"})
	c.Abort()
}

func HandlerInvokeError(c *gin.Context) {
	c.JSON(http.StatusInternalServerError, gin.H{"error": "调用接口失败"})
	c.Abort()
}
