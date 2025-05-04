package main

import (
	"Go_API/routes"
	"fmt"
	"log"
	"net/http"
)

func main() {
	// Setup routes
	router := routes.SetupRoutes()

	// Start server
	port := ":8080"
	fmt.Printf("Server starting on port %s...\n", port)
	if err := http.ListenAndServe(port, router); err != nil {
		log.Fatal(err)
	}
}
