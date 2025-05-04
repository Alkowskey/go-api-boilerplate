package routes

import (
	"Go_API/config"
	"Go_API/handlers"

	"github.com/gorilla/mux"
)

func SetupRoutes() *mux.Router {
	router := mux.NewRouter()

	// Initialize database
	config.InitDB()

	// Create user handler
	userHandler := handlers.NewUserHandler()

	// API routes
	api := router.PathPrefix("/api").Subrouter()

	// Hello route
	api.HandleFunc("/hello", handlers.HelloHandler).Methods("GET")

	// Health route
	api.HandleFunc("/health", handlers.HealthHandler).Methods("GET")

	// User routes
	api.HandleFunc("/register", userHandler.Register).Methods("POST")
	api.HandleFunc("/login", userHandler.Login).Methods("POST")
	api.HandleFunc("/users", userHandler.ListUsers).Methods("GET")
	api.HandleFunc("/users/{id}", userHandler.GetUser).Methods("GET")

	return router
}
