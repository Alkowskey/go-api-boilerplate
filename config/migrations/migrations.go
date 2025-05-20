package migrations

import (
	"gorm.io/gorm"

	acceleration "github.com/aleksander/Go_API/internal/domain/acceleration/models"
	device "github.com/aleksander/Go_API/internal/domain/device/models"
	user "github.com/aleksander/Go_API/internal/domain/user/models"
)

// Migrations is an array of all database models that need to be migrated
var Migrations = []interface{}{
	&user.User{},
	&device.Device{},
	&acceleration.Acceleration{},
}

// RunMigrations executes all database migrations
func RunMigrations(db *gorm.DB) error {
	return db.AutoMigrate(Migrations...)
}
