package services

import (
	"todo/db"
	model "todo/models"
	// "gorm.io/gorm"
)

func GetTodos(id string) []model.Todo {
	var todos []model.Todo
	if err := db.DB.Where("user_id = ?", id).Find(&todos).Error; err != nil {
		return []model.Todo{}
	}
	return todos
}
