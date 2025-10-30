package main

import (
	"log"
	"net/http"

	"mvc-app/config"
	"mvc-app/controllers"
	"mvc-app/routes"
	"mvc-app/services"

	"github.com/joho/godotenv"
)

func main() {
	// Load .env file if it exists
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system environment variables")
	}
	// Setup MySQL database
	dbConfig := config.GetDatabaseConfig()
	db, err := config.ConnectDatabase(dbConfig)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	// Create users table if it doesn't exist
	if err := config.CreateUsersTable(db); err != nil {
		log.Fatalf("Failed to create users table: %v", err)
	}

	// Initialize service and controller
	userService := services.NewUserService(db)
	userController := controllers.NewUserController(userService)

	// Setup routes
	router := routes.SetupRoutes(userController)

	// Start server
	log.Println("Connected to MySQL database")
	log.Println("Server starting on :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}