package routes

import (
	"go-fin-microservice/services/admin/internal/handlers"

	"github.com/gin-gonic/gin"
)

func RegisterRoleRoutes(r *gin.RouterGroup) {
	roleGroup := r.Group("/roles")
	{
		roleGroup.GET("/", handlers.GetRoles)
		roleGroup.GET("/:id", handlers.GetRoleByID)
		roleGroup.GET("/", handlers.CreateRole)
		roleGroup.GET("/:id", handlers.UpdateRole)
		roleGroup.GET("/:id", handlers.DeleteRole)
	}
}
