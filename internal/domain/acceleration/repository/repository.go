package repository

import (
	acceleration "github.com/aleksander/Go_API/internal/domain/acceleration/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type accelerationRepository struct {
	db *gorm.DB
}

func NewAccelerationRepository(db *gorm.DB) acceleration.Repository {
	return &accelerationRepository{db: db}
}

func (r *accelerationRepository) Create(device *acceleration.Acceleration) error {
	return r.db.Create(device).Error
}

func (r *accelerationRepository) FindByID(id uuid.UUID) (*acceleration.Acceleration, error) {
	var acceleration acceleration.Acceleration
	err := r.db.First(&acceleration, id).Error
	return &acceleration, err
}

func (r *accelerationRepository) FindAll() ([]acceleration.Acceleration, error) {
	var accelerations []acceleration.Acceleration
	err := r.db.Find(&accelerations).Error
	return accelerations, err
}
