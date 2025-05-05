package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/aleksander/Go_API/config"
	"github.com/aleksander/Go_API/internal/infrastructure/router"
)

func main() {
	// Initialize database
	config.InitDB()

	// Setup routes
	router := router.SetupRouter(config.DB)

	// Start server
	port := ":8080"
	fmt.Printf("Server starting on port %s...\n", port)
	if err := http.ListenAndServe(port, router); err != nil {
		log.Fatal(err)
	}
}
