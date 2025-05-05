package usecase

import (
	"strings"

	device "github.com/aleksander/Go_API/internal/domain/device/models"
)

type RegisterDeviceUseCase struct {
	repo device.Repository
}

func NewRegisterDeviceUseCase(repo device.Repository) *RegisterDeviceUseCase {
	return &RegisterDeviceUseCase{repo: repo}
}

type RegisterDeviceInput struct {
	Name string
}

type RegisterDeviceOutput struct {
	Device *device.Device
	Err    error
}

func (uc *RegisterDeviceUseCase) Execute(input RegisterDeviceInput) RegisterDeviceOutput {
	// Validation
	if strings.TrimSpace(input.Name) == "" {
		return RegisterDeviceOutput{Err: ErrNameRequired}
	}

	// Create device
	device := &device.Device{
		Name: strings.TrimSpace(input.Name),
	}

	// Save to repository
	if err := uc.repo.Create(device); err != nil {
		return RegisterDeviceOutput{Err: err}
	}

	return RegisterDeviceOutput{Device: device}
}
