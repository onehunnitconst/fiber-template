package models

import (
	"gorm.io/gorm"
	"gorm.io/driver/sqlite"
)

func Connect() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("database.db"), &gorm.Config{})

	if err != nil {
		panic(err)
	}
	
	db.AutoMigrate(&Todo{})

	return db;
}