package utils

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var dbSingleton *gorm.DB

func DB(config *DbConfig) (*gorm.DB, error) {
	if dbSingleton != nil {
		return dbSingleton, nil
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", config.Host, config.User,
		config.Password, config.Name, config.Port)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	dbSingleton = db
	return dbSingleton, nil
}
