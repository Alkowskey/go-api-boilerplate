package device

import (
	"github.com/aleksander/Go_API/internal/domain/device/handler"
	"github.com/aleksander/Go_API/internal/domain/device/repository"
	"github.com/aleksander/Go_API/internal/domain/device/usecase"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func SetupRoutes(router *mux.Router, db *gorm.DB) {
	repo := repository.NewDeviceRepository(db)
	registerUseCase := usecase.NewRegisterDeviceUseCase(repo)
	deviceHandler := handler.NewDeviceHandler(registerUseCase)

	router.HandleFunc("/device", deviceHandler.RegisterDevice).Methods("POST")
}
