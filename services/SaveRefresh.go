package services

import (
	"time"
	model "todo/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func SaveOrUpdateRefreshToken(db *gorm.DB, userID string, hashedToken string, expiry time.Time) error {
	refresh := model.RefreshToken{
		UserID:    userID,
		Token:     hashedToken,
		CreatedAt: time.Now(),
		Revoked:   false,
		RevokedAt: &expiry,
	}

	return db.Clauses(clause.OnConflict{
		Columns: []clause.Column{{Name: "user_id"}}, // conflict column
		DoUpdates: clause.Assignments(map[string]interface{}{
			"token":      hashedToken,
			"created_at": time.Now(),
			"revoked":    false,
			"revoked_at": expiry,
		}),
	}).Create(&refresh).Error
}
