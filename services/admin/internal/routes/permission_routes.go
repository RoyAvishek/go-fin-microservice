package routes

import (
	"go-fin-microservice/services/admin/internal/handlers"

	"github.com/gin-gonic/gin"
)

func RegisterPermissionRoutes(r *gin.RouterGroup) {
	permissionGroup := r.Group("/permissions")
	{
		permissionGroup.GET("/", handlers.GetPermissions)
	}
}
