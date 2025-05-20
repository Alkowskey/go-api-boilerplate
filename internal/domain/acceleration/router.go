package acceleration

import (
	"github.com/aleksander/Go_API/internal/domain/acceleration/handler"
	"github.com/aleksander/Go_API/internal/domain/acceleration/repository"
	"github.com/aleksander/Go_API/internal/domain/acceleration/usecase"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func SetupRoutes(router *mux.Router, db *gorm.DB) {
	repo := repository.NewAccelerationRepository(db)
	importAccelerationDataUseCase := usecase.NewImportAccelerationDataUseCase(repo)
	getAccelerationDataUseCase := usecase.NewGetAccelerationDataUseCase(repo)
	accelerationHandler := handler.NewAccelerationHandler(importAccelerationDataUseCase, getAccelerationDataUseCase)

	router.HandleFunc("/acceleration", accelerationHandler.ImportAccelerationData).Methods("POST", "OPTIONS")
	router.HandleFunc("/acceleration/{id}", accelerationHandler.GetAccelerationData).Methods("GET", "OPTIONS")
}
