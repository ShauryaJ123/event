package main

import (
	"log"

	"abc.com/calc/db"
	"abc.com/calc/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize the database
	db.InitDB()

	// Create a Gin router
	server := gin.Default()

	// Register routes
	routes.RegisterRoutes(server)

	// Log a message to confirm server startup
	log.Println("Server is starting on port 8080...")

	// Start the server
	err := server.Run(":8080")
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
