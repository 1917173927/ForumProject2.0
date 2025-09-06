package migrations

import (
	"gorm.io/gorm"
)

func CreateReportsTable(db *gorm.DB) error {
	type Report struct {
		gorm.Model
		PostID     uint   `gorm:"not null"`
		ReporterID uint   `gorm:"not null"`
		Reason     string `gorm:"type:text;not null"`
		Status     int    `gorm:"default:0"`
		Type       string `gorm:"type:varchar(20)"`
	}

	return db.AutoMigrate(&Report{})
}
