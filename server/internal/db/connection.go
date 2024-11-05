package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewConnection() (*gorm.DB, error) {
	const DATA_SOURCE_NAME = "host=localhost user=postgres password=bellingham dbname=recycle_it port=5432"
	db, err := gorm.Open(postgres.Open(DATA_SOURCE_NAME), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}
