package services

import (
	"todo/db"
	model "todo/models"
)

func GetRefreshfeild(refresh string) (*model.RefreshToken, error) {
	hashed := GenerateHash(refresh)
	var feild model.RefreshToken
	err := db.DB.Where("token = ?", hashed).Find(&feild)
	if err.Error != nil {
		return nil, err.Error
	}
	// fmt.Println(feild)
	return &feild, nil
}
