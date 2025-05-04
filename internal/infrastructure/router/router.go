package router

import (
	"Go_API/internal/auth"
	"Go_API/internal/domain/user/handler"
	"Go_API/internal/domain/user/repository"
	"Go_API/internal/domain/user/usecase"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB) *mux.Router {
	router := mux.NewRouter()

	// Initialize repository
	userRepo := repository.NewUserRepository(db)

	// Initialize use cases
	registerUseCase := usecase.NewRegisterUseCase(userRepo)
	authenticateUseCase := usecase.NewAuthenticateUseCase(userRepo)
	getUserUseCase := usecase.NewGetUserUseCase(userRepo)
	listUsersUseCase := usecase.NewListUsersUseCase(userRepo)

	// Initialize handler
	userHandler := handler.NewUserHandler(
		registerUseCase,
		authenticateUseCase,
		getUserUseCase,
		listUsersUseCase,
	)

	// API routes
	api := router.PathPrefix("/api").Subrouter()

	// Public routes
	api.HandleFunc("/register", userHandler.Register).Methods("POST")
	api.HandleFunc("/login", userHandler.Login).Methods("POST")

	// Protected routes
	protected := api.PathPrefix("").Subrouter()
	protected.Use(auth.AuthMiddleware)
	protected.HandleFunc("/users", userHandler.ListUsers).Methods("GET")
	protected.HandleFunc("/users/{id}", userHandler.GetUser).Methods("GET")

	return router
}
