package routes

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	adminGroup := router.Group("/api/admin")
	{
		RegisterUserRoutes(adminGroup)
		RegisterRoleRoutes(adminGroup)
		RegisterPermissionRoutes(adminGroup)
	}
}
