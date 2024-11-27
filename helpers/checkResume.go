package helpers

import (
	"coursework/internal/models"
	"errors"

	"gorm.io/gorm"
)

func CheckUserResumeExists(db *gorm.DB, userID uint) *models.Resume {
	var resume models.Resume
	err := db.Where("user_id = ?", userID).First(&resume).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil
	}
	return &resume
}
