package router

import (
	"github.com/aleksander/Go_API/internal/domain/acceleration"
	"github.com/aleksander/Go_API/internal/domain/device"
	"github.com/aleksander/Go_API/internal/domain/user"
	"github.com/aleksander/Go_API/internal/infrastructure/middleware"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB) *mux.Router {
	router := mux.NewRouter()

	// API routes
	api := router.PathPrefix("/api").Subrouter()

	// Add CORS middleware to the API subrouter
	api.Use(middleware.CORSMiddleware)

	user.SetupRoutes(api, db)
	device.SetupRoutes(api, db)
	acceleration.SetupRoutes(api, db)

	return router
}
