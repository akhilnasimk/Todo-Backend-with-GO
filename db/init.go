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
	errr := DB.AutoMigrate(&model.Todo{})
	if errr != nil {
		fmt.Println("migration failed ", errr)
		return errr
	}
	MER := DB.AutoMigrate(&model.Users{}, &model.Todo{})
	if MER != nil {
		return MER
	}
	fmt.Println("migration completed and Todo is ready ")
	return nil

}
