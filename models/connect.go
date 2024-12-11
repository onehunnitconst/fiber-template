package models

import (
	"gorm.io/gorm"
	"gorm.io/driver/sqlite"
)

var Database *gorm.DB

func Connect() error {
	var err error

	Database, err := gorm.Open(sqlite.Open("database.db"), &gorm.Config{})

	if err != nil {
		panic(err)
	}
	
	Database.AutoMigrate(&Todo{})

	return nil;
}