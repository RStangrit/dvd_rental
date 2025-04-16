package migrations

import (
	"fmt"
	"main/internal/models"
	"main/pkg/db"
)

func launchMigrationsGORM() error {
	models.ReorderModels()

	if len(models.ModelRegistry) == 0 {
		return fmt.Errorf("no models registered for migration")
	}

	for _, model := range models.ModelRegistry {
		fmt.Printf("Starting migration for table: %T\n", model)
	}

	return db.GORM.AutoMigrate(models.ModelRegistry...)
}
