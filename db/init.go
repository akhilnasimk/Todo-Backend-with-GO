package db

import (
	"fmt"
	model "todo/models"

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
	MER := DB.AutoMigrate(&model.Users{}, &model.Todo{}, &model.RefreshToken{})
	if MER != nil {
		return MER
	}
	fmt.Println("migration completed and Todo is ready ")
	return nil

}
