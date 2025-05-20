package repository

import (
	device "github.com/aleksander/Go_API/internal/domain/device/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type deviceRepository struct {
	db *gorm.DB
}

func NewDeviceRepository(db *gorm.DB) device.Repository {
	return &deviceRepository{db: db}
}

func (r *deviceRepository) Create(device *device.Device) error {
	return r.db.Create(device).Error
}

func (r *deviceRepository) FindByID(id uuid.UUID) (*device.Device, error) {
	var device device.Device
	err := r.db.First(&device, id).Error
	return &device, err
}

func (r *deviceRepository) FindAll() ([]device.Device, error) {
	var devices []device.Device
	err := r.db.Find(&devices).Error
	return devices, err
}
