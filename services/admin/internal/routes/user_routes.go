package routes

import (
	"go-fin-microservice/services/admin/internal/handlers"

	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(r *gin.RouterGroup) {
	userGroup := r.Group("/users")
	{
		userGroup.GET("/", handlers.GetUsers)
		userGroup.GET("/:id", handlers.GetUserByID)
		userGroup.GET("/", handlers.CreateUser)
		userGroup.GET("/:id", handlers.UpdateUser)
		userGroup.GET("/:id", handlers.DeleteUser)
	}
}
