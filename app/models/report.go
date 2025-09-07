package models

import "gorm.io/gorm"

type Report struct {
	gorm.Model
	PostID    uint   `json:"post_id"`    // 被举报的帖子ID
	UserID    uint   `json:"user_id"`    // 举报人ID
	Reason    string `json:"reason"`     // 举报原因
	Status    int    `json:"status" gorm:"default:0"`      // 举报状态
	Type      string `json:"type"`        // 举报类型
}

func (Report) TableName() string {
	return "reports"
}
