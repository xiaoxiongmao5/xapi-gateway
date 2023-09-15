package router

import (
	"xj/xapi-gateway/controller"
	"xj/xapi-gateway/middleware"

	"github.com/gin-gonic/gin"
)

func ManagerRouter(r *gin.Engine) {
	router := r.Group("/manage", middleware.FilterWithAccessControlInAdminIp())
	router.GET("/config/ratelimit", controller.GetIPRateLimitConfig)
	router.PUT("/config/ratelimit", controller.UpdateIPRateLimitConfig)
}
