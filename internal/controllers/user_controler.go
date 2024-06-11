// controllers/user_controller.go
package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"com.hypnovai.documentpublishing/internal/models"
)

// / UserController handles user-related actions
type UserController struct {
	DB *gorm.DB
}

// NewUserController creates a new instance of UserController
func NewUserController(db *gorm.DB) *UserController {
	return &UserController{DB: db}
}

// GetUserList retrieves a list of all users from the database
func (uc *UserController) GetUserList(c *gin.Context) {
	// Check if user is authenticated
	if _, exists := c.Get("claims"); !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	// Query the database to retrieve all users
	var users []models.User
	if err := uc.DB.Find(&users).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve user data"})
		return
	}

	c.JSON(http.StatusOK, users)
}

// ChangePassword updates the user's password in the database
func (uc *UserController) ChangePassword(c *gin.Context) {
	// Retrieve user ID and new password from request
	username := c.PostForm("username")
	newPassword := c.PostForm("new_password")

	// Authenticate and authorize user - Implement as per your JWT authentication logic

	// Retrieve user from the database
	var user models.User
	if err := uc.DB.Where("username = ?", username).First(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to find user"})
		return
	}

	// Change the user's password
	if err := user.ChangePassword(uc.DB, newPassword); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update password"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Password updated successfully"})
}
