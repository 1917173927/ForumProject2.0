package repository

import (
	"gorm.io/gorm"
)

type Reported struct {
	gorm.Model
	UserID  int    `gorm:"column:user_id" json:"user_id"`
	PostID  int    `gorm:"column:post_id" json:"post_id"`
	Content string `gorm:"column:content" json:"content"`
	Reason  string `gorm:"column:reason" json:"reason"`
	Status  int    `gorm:"column:status" json:"status"`
}

func GetPendingReports(db *gorm.DB) ([]Reported, error) {
	var reports []Reported
	if err := db.Where("status = 0").Find(&reports).Error; err != nil {
		return nil, err
	}
	return reports, nil
}

func ProcessReportApproval(db *gorm.DB, reportID int, approval int) error {
	if approval == 1 {
		// 删除帖子和举报记录
		return db.Transaction(func(tx *gorm.DB) error {
			var report Reported
			if err := tx.First(&report, reportID).Error; err != nil {
				return err
			}
			if err := tx.Where("post_id = ?", report.PostID).Delete(&Reported{}).Error; err != nil {
				return err
			}
			return tx.Delete(&report).Error
		})
	} else {
		// 更新举报状态为已处理
		return db.Model(&Reported{}).Where("id = ?", reportID).Update("status", 1).Error
	}
}

func InitDB() error {
	// Placeholder for initializing admin repository database connection
	return nil
}
