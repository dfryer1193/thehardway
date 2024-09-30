package models

import (
	"errors"
	"gorm.io/gorm"
	"time"
)

type Post struct {
	ID        string    `json:"id" gorm:"primaryKey"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	Images    []string  `json:"images" gorm:"type:text[]"`
	State     string    `json:"state"` // "draft" or "published"
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// Before saving, render the Markdown to HTML
func (p *Post) BeforeSave(tx *gorm.DB) (err error) {
	if p.Content == "" {
		return errors.New("content cannot be empty")
	}
	return nil
}

// UpdatePost updates an existing post
func UpdatePost(db *gorm.DB, post Post) error {
	if err := db.Save(&post).Error; err != nil {
		return err
	}
	return nil
}

// GetPostByID fetches a post by ID
func GetPostByID(db *gorm.DB, id string) (Post, error) {
	var post Post
	if err := db.First(&post, id).Error; err != nil {
		return post, err
	}

	return post, nil
}
