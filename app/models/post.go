package models

import "time"

type Post struct {
	PostID    uint      `gorm:"primaryKey;column:post_id;not null" json:"PostID"`
	Content   string    `gorm:"column:content;not null" json:"Content"`
	UserID    uint      `gorm:"column:user_id;not null" json:"UserID"`
	CreatedAt time.Time `gorm:"autoCreateTime;column:created_at;not null" json:"CreatedAt"`
}

func (Post) TableName() string {
	return "posts"
}
