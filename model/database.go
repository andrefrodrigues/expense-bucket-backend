package model

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func SetupDatabase(dsn string) (*gorm.DB, error) {

	db, err := gorm.Open(postgres.Open(dsn))
	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&User{})

	return db, nil
}
