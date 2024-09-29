package models

import "time"

type Post struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	Markdown  string    `json:"markdown"`
	Images    []string  `json:"images" gorm:"type:text[]"`
	State     string    `json:"state"` // "draft" or "published"
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	AuthorID  uint      `json:"author_id"`
}
