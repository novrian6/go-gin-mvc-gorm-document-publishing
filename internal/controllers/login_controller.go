package controllers

import (
	"net/http"
	"time"

	"com.hypnovai.documentpublishing/internal/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

const secretKey = "novasecret123"

// LoginController handles login related operations
type LoginController struct {
	DB *gorm.DB
}

// NewLoginController creates a new instance of LoginController
func NewLoginController(db *gorm.DB) *LoginController {
	return &LoginController{DB: db}
}

// ShowLoginForm displays the login form
func (lc *LoginController) ShowLoginForm(c *gin.Context) {
	// Render the login form HTML template
	c.HTML(http.StatusOK, "login.html", nil)
}

// Login handles user login
func (lc *LoginController) Login(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	// Query the user record from the database
	var user models.User
	if err := lc.DB.Where("username = ?", username).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}

	// Compare the hashed password with the provided password
	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}

	// Authentication successful
	// Generate JWT token
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = username
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix() // Token expiration time (24 hours)
	signedToken, err := token.SignedString([]byte(secretKey))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	c.SetCookie("token", signedToken, 3600, "/", "", false, true)

	c.JSON(http.StatusOK, gin.H{"token": signedToken})
}
