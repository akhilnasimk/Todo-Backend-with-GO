package services

import (
	// "fmt"
	"todo/db"
	model "todo/models"

	"gorm.io/gorm"
)

func FindUser(R model.Req) (*gorm.DB, model.Users) {
	var user model.Users
	result := db.DB.Where("email = ?", R.Email).First(&user)
	return result, user
}
