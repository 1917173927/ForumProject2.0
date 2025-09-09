package models

import "gorm.io/gorm"

type Report struct {
	gorm.Model
	PostID    uint   `json:"PostID"`    // 被举报的帖子ID
	UserID    uint   `json:"UserID"`    // 举报人ID
	Reason    string `json:"Reason"`     // 举报原因
	Status    int    `json:"Status" gorm:"default:0"`      // 举报状态
	Type      string `json:"Type"`        // 举报类型
}

func (Report) TableName() string {
	return "reports"
}
