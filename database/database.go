package database

import (
	"gorm.io/gorm"
)

type Database struct {
	DB *gorm.DB
}

func InitDatabade(db *gorm.DB) *Database {
	return &Database{DB: db}
}
