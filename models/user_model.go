package model

type Users struct {
	User_id  string `gorm:"user_id" `
	Name     string `gorm:"name" json:"name"`
	Email    string `gorm:"email" josn:"email"`
	Password string `gorm:"password" json:"password"`
}
