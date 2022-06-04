package utils

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func DB(config *DbConfig) *gorm.DB {
	if db != nil {
		return db
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", config.Host, config.User,
		config.Password, config.Name, config.Port)

	newDb, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db = newDb
	return db
}
