package services

import (
	"todo/db"
	model "todo/models"

	"gorm.io/gorm"
)

func Insert_todo(todo model.Todo) *gorm.DB {
	res := db.DB.Create(&todo)
	return res
}
