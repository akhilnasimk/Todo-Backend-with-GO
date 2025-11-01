package model

import "gorm.io/gorm"

type Todo struct {
	gorm.Model
	//creating the foring key  relation with gorm
	UserID      string `gorm:"column:user_id;type:varchar;not null"`                                               // the actual foreign key feild
	// User        Users  `gorm:"foreignKey:UserID;references:User_id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"` // saying the forign key is UserID and it referes to the User_id of tbe User struct
	Messsage    string `gorm:"not null " json:"message"`
	IsCompleted bool   `gorm:"default:false"`
}
