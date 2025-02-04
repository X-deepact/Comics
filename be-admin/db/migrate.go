package db

import (
	"gorm.io/gorm"
	"log"
)

func RunMigrations(db *gorm.DB) error {
	log.Println("Running migrations...")

	// Add your models here to migrate
	models := []interface{}{
		&User{},
		/*&Comic{},*/
	}

	for _, model := range models {
		if err := db.AutoMigrate(model); err != nil {
			log.Printf("Failed to migrate model %T: %v", model, err)
			return err
		}
	}

	log.Println("Migrations completed successfully!")
	return nil
}
