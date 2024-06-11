package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	//s"document_platform/internal/routes"
	"gorm.io/driver/postgres"

	"com.hypnovai.documentpublishing/internal/routes"
	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

const (
	host     = "172.16.137.133" //"192.168.1.161"
	port     = 5432
	user     = "postgres"
	password = "zuruck"
	dbname   = "document_platform_db"
)

func main() {
	// Connect to PostgreSQL database
	db, err := connectDB()

	// Initialize Gin
	router := gin.Default()

	// Get the absolute path of the directory containing the main.go file
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	// Navigate up one level to the project root
	projectRoot := filepath.Dir(filepath.Dir(dir))

	// Define the relative path to the views directory
	viewsDir := filepath.Join(projectRoot, "internal", "views")

	router.LoadHTMLGlob(filepath.Join(viewsDir, "*.html"))

	// Load routes
	routes.SetupRoutes(router, db)

	// Start the Gin server
	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}

// connectDB establishes a connection to the PostgreSQL database using GORM
func connectDB() (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
