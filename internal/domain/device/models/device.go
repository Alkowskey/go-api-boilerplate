package device

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Device struct {
	ID        uuid.UUID      `json:"id" gorm:"type:uuid;primary_key;"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at,omitempty" gorm:"index"`
	Name      string         `json:"name" gorm:"not null"`
}

func (d *Device) BeforeCreate(tx *gorm.DB) error {
	d.ID = uuid.New()
	return nil
}
