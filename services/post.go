package services

import "github.com/dfryer1193/thehardway/models"
import "gorm.io/gorm"

func CreatePost(db *gorm.DB, post *models.Post) error {
	return db.Create(post).Error
}

func GetPublishedPosts(db *gorm.DB) ([]models.Post, error) {
	var posts []models.Post
	err := db.Where("state = ?", "published").Find(&posts).Error
	return posts, err
}

func UpdatePostState(db *gorm.DB, postID uint, state string) error {
	return db.Model(&models.Post{}).Where("id = ?", postID).Update("state", state).Error
}
