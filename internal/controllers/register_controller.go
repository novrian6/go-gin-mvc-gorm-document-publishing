// controllers/register_controller.go
package controllers

import (
	"net/http"

	"com.hypnovai.documentpublishing/internal/models"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// RegisterController handles user registration
type RegisterController struct {
	DB *gorm.DB
}

// NewRegisterController creates a new instance of RegisterController
func NewRegisterController(db *gorm.DB) *RegisterController {
	return &RegisterController{DB: db}
}

// Register processes registration form submission
func (rc *RegisterController) Register(c *gin.Context) {
	var newUser models.User
	err := c.ShouldBindJSON(&newUser)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newUser.PasswordHash), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}
	newUser.PasswordHash = string(hashedPassword)

	// Validate username and password here...

	// Insert new user into the database
	err = newUser.Register(rc.DB)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User registered successfully"})
}
