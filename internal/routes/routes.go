package routes

import (
	"com.hypnovai.documentpublishing/internal/controllers"
	"com.hypnovai.documentpublishing/internal/middleware"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// SetupRoutes configures routes for the application
func SetupRoutes(router *gin.Engine, db *gorm.DB) {
	// Initialize controllers

	indexController := controllers.NewIndexController()
	loginController := controllers.NewLoginController(db)
	registerController := controllers.NewRegisterController(db)
	userController := controllers.NewUserController(db)

	// Define routes
	router.GET("/", indexController.Index)

	//Login
	router.GET("/login", loginController.ShowLoginForm)
	router.POST("/login", loginController.Login)

	//Register

	router.POST("/register", registerController.Register)

	// API routes
	//test using req bin to  http://localhost:8080/api/users http get, pass the token by custom
	api := router.Group("/api")
	api.Use(middleware.AuthMiddleware()) // Apply JWT authentication middleware
	{
		api.GET("/users", userController.GetUserList)
		api.POST("/users/change-password", userController.ChangePassword)
	}

}
