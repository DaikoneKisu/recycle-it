package db

import (
	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) {
	db.AutoMigrate(&Player{})
	db.AutoMigrate(&Lobby{})
}
