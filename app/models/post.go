package models

import "time"

type Post struct {
	PostID    uint      `gorm:"primaryKey;column:post_id"`
	Content   string    `gorm:"column:content;not null"`
	UserID    uint      `gorm:"column:user_id;not null"`
	CreatedAt time.Time `gorm:"autoCreateTime;column:created_at"`
}

func (Post) TableName() string {
	return "posts"
}
