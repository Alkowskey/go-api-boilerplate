package usecase

import (
	acceleration "github.com/aleksander/Go_API/internal/domain/acceleration/models"
	"github.com/google/uuid"
)

type ImportAccelerationDataUseCase struct {
	repo acceleration.Repository
}

func NewImportAccelerationDataUseCase(repo acceleration.Repository) *ImportAccelerationDataUseCase {
	return &ImportAccelerationDataUseCase{repo: repo}
}

type ImportAccelerationDataInput struct {
	X float64
	Y float64
	Z float64
}

type ImportAccelerationDataOutput struct {
	Acceleration *acceleration.Acceleration
	Err          error
}

func (uc *ImportAccelerationDataUseCase) Execute(input ImportAccelerationDataInput, deviceId uuid.UUID) ImportAccelerationDataOutput {
	// Validation
	if input.X == 0 && input.Y == 0 && input.Z == 0 {
		return ImportAccelerationDataOutput{Err: ErrAccelerationDataRequired}
	}

	// Create acceleration data
	acceleration := &acceleration.Acceleration{
		DeviceID: deviceId,
		X:        input.X,
		Y:        input.Y,
		Z:        input.Z,
	}

	// Save to repository
	if err := uc.repo.Create(acceleration); err != nil {
		return ImportAccelerationDataOutput{Err: err}
	}

	return ImportAccelerationDataOutput{Acceleration: acceleration}
}
