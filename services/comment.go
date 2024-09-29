package services

import (
	"github.com/dfryer1193/thehardway/models"
	"gorm.io/gorm"
)

func AddComment(db *gorm.DB, comment *models.Comment) error {
	return db.Create(comment).Error
}

func DeleteComment(db *gorm.DB, commentID uint) error {
	return db.Model(&models.Comment{}).Where("id = ?", commentID).Update("is_deleted", true).Error
}

func BanUserByEmail(db *gorm.DB, email string) error {
	return db.Model(&models.Comment{}).Where("author_email = ?", email).Update("is_banned", true).Error
}
