/*
/api/admin/permissions	GET	List all permissions
*/

package handlers

import (
	"net/http"

	"go-fin-microservice/services/admin/internal/database"
	"go-fin-microservice/services/admin/internal/models"

	"github.com/gin-gonic/gin"
)

// GetRoles retrieves all permissions from the database
func GetPermissions(c *gin.Context) {
	var permissions []models.Permission
	if err := database.DB.Find(&permissions).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve permissions"})
		return
	}
	c.JSON(http.StatusOK, permissions)
}
