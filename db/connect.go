package db

import (
	"github.com/jeff-rdg/carcontrol/entity"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

func Init() *gorm.DB {
	db, err := gorm.Open(postgres.Open(os.Getenv("DB_DSN")), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	err = migrateEntities(db)
	if err != nil {
		panic("failed to migrate database")
	}

	return db
}

func migrateEntities(db *gorm.DB) error {
	return db.AutoMigrate(&entity.Company{})
}
