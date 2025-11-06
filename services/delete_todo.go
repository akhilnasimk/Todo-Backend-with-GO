package services

import (
	"strconv"
	"todo/db"
	model "todo/models"

	"gorm.io/gorm"
)

func Delete_todo(id string) *gorm.DB {
	todoID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return &gorm.DB{Error: err}
	}
	res := db.DB.Delete(&model.Todo{}, todoID)
	return res
}
