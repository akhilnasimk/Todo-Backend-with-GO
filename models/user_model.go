package model

type Users struct {
	User_id  string `gorm:"column:user_id;type:uuid;primaryKey"`
	Name     string `gorm:"name" json:"name"`
	Email    string `gorm:"column:email" json:"email"`
	Password string `gorm:"password" json:"password"`
	Age      int    `gorm:"default:20"`
}
