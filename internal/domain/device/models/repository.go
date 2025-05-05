package device

import "github.com/google/uuid"

type Repository interface {
	Create(device *Device) error
	FindByID(id uuid.UUID) (*Device, error)
	FindAll() ([]Device, error)
}
