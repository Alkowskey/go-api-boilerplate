package acceleration

import "github.com/google/uuid"

type Repository interface {
	Create(acceleration *Acceleration) error
	FindByID(id uuid.UUID) (*Acceleration, error)
	FindAll() ([]Acceleration, error)
}
