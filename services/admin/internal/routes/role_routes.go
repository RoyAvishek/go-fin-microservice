package routes

import (
	"go-fin-microservice/services/admin/internal/handlers"

	"github.com/gin-gonic/gin"
)

func RegisterRoleRoutes(r *gin.RouterGroup) {
	roleGroup := r.Group("/roles")
	{
		roleGroup.GET("/", handlers.GetRoles)         // GET /api/admin/roles
		roleGroup.GET("/:id", handlers.GetRoleByID)   // GET /api/admin/roles/:id
		roleGroup.POST("/", handlers.CreateRole)      // POST /api/admin/roles
		roleGroup.PUT("/:id", handlers.UpdateRole)    // PUT /api/admin/roles/:id
		roleGroup.DELETE("/:id", handlers.DeleteRole) // DELETE /api/admin/roles/:id
	}
}
