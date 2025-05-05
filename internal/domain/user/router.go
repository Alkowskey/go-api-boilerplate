package user

import (
	"github.com/aleksander/Go_API/internal/auth"
	"github.com/aleksander/Go_API/internal/domain/user/handler"
	"github.com/aleksander/Go_API/internal/domain/user/repository"
	"github.com/aleksander/Go_API/internal/domain/user/usecase"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func SetupRoutes(router *mux.Router, db *gorm.DB) {
	repo := repository.NewUserRepository(db)
	registerUseCase := usecase.NewRegisterUseCase(repo)
	authenticateUseCase := usecase.NewAuthenticateUseCase(repo)
	getUserUseCase := usecase.NewGetUserUseCase(repo)
	listUsersUseCase := usecase.NewListUsersUseCase(repo)

	userHandler := handler.NewUserHandler(
		registerUseCase,
		authenticateUseCase,
		getUserUseCase,
		listUsersUseCase,
	)

	// Public routes
	router.HandleFunc("/register", userHandler.Register).Methods("POST")
	router.HandleFunc("/login", userHandler.Login).Methods("POST")

	// Protected routes
	protected := router.PathPrefix("").Subrouter()
	protected.Use(auth.AuthMiddleware)

	// Protected routes
	protected.HandleFunc("/users", userHandler.ListUsers).Methods("GET")
	protected.HandleFunc("/users/{id}", userHandler.GetUser).Methods("GET")
}
