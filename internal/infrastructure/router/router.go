package router

import (
	"github.com/aleksander/Go_API/internal/domain/device"
	"github.com/aleksander/Go_API/internal/domain/user"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB) *mux.Router {
	router := mux.NewRouter()

	// API routes
	api := router.PathPrefix("/api").Subrouter()

	user.SetupRoutes(api, db)
	device.SetupRoutes(api, db)

	return router
}
