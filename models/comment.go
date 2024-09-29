package models

import "time"

type Comment struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	PostID      uint      `json:"post_id"`
	AuthorEmail string    `json:"author_email"`
	Content     string    `json:"content"`
	CreatedAt   time.Time `json:"created_at"`
	IsDeleted   bool      `json:"is_deleted" gorm:"default:false"`
	IsBanned    bool      `json:"is_banned" gorm:"default:false"`
}
