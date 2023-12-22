package db

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func ConnectDatabase(path string) (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(path), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&Todo{})

	return db, nil
}
