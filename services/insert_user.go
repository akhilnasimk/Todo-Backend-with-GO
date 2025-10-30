package services

import (
	"todo/db"
	model "todo/models"

	"gorm.io/gorm"
)

func InserUser(U model.Users) *gorm.DB {
	err := db.DB.Create(&U)
	return err
}
