package model

import "time"


type RefreshToken struct {
	RefreshID int        `gorm:"primaryKey;autoIncrement"`
	UserID    string     `gorm:"column:user_id;type:uuid;unique;not null"`
	User      Users      `gorm:"foreignKey:UserID;references:User_id;constraint:OnDelete:CASCADE"`
	Token     string     `gorm:"uniqueIndex;not null"`
	CreatedAt time.Time  `gorm:"not null"`
	Revoked   bool       `gorm:"default:false"`
	RevokedAt *time.Time `gorm:"default:null"`
}
