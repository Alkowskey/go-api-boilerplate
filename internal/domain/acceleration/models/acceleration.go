package acceleration

import (
	"time"

	device "github.com/aleksander/Go_API/internal/domain/device/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Acceleration represents acceleration data from a device
type Acceleration struct {
	ID        uuid.UUID      `json:"id" gorm:"type:uuid;primary_key;"`
	DeviceID  uuid.UUID      `json:"device_id" gorm:"type:uuid;not null"`
	Device    *device.Device `json:"device,omitempty" gorm:"foreignKey:DeviceID;references:ID"`
	X         float64        `json:"x" bson:"x"` // X-axis acceleration in m/s²
	Y         float64        `json:"y" bson:"y"` // Y-axis acceleration in m/s²
	Z         float64        `json:"z" bson:"z"` // Z-axis acceleration in m/s²
	Timestamp time.Time      `json:"timestamp" bson:"timestamp"`
	CreatedAt time.Time      `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time      `json:"updated_at" bson:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at,omitempty" gorm:"index"`
}

func (acceleration *Acceleration) BeforeCreate(tx *gorm.DB) error {
	acceleration.ID = uuid.New()
	return nil
}

// GetMagnitude calculates the total acceleration magnitude
func (a *Acceleration) GetMagnitude() float64 {
	return (a.X*a.X + a.Y*a.Y + a.Z*a.Z)
}
