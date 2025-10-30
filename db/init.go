package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init(conf string) error {
	var err error
	DB, err = gorm.Open(postgres.Open(conf), &gorm.Config{})
	if err != nil {
		return err
	}
	return nil
	
}
