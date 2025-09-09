package migrations

import (
	"time"
	
	"gorm.io/gorm"
)

func CreateTables(db *gorm.DB) error {
	type Report struct {
		gorm.Model
		PostID     uint   `gorm:"not null"`
		ReporterID uint   `gorm:"not null"`
		Reason     string `gorm:"type:text;not null"`
		Status     int    `gorm:"default:0"`
		Type       string `gorm:"type:varchar(20)"`
	}

	type Account struct {
		UserID   int    `gorm:"primaryKey"`
		Password string `gorm:"type:varchar(255)"`
		UserType int
		Name     string `gorm:"type:varchar(100)"`
		Username string `gorm:"type:varchar(50)"`
	}

	type Post struct {
		PostID    int       `gorm:"primaryKey"`
		Time      time.Time `gorm:"type:timestamp"`
		UserID    int64     `gorm:"type:bigint"`
		Content   string    `gorm:"type:longtext"`
		CreatedAt time.Time `gorm:"type:datetime(3)"`
	}

	if err := db.AutoMigrate(&Report{}); err != nil {
		return err
	}
	if err := db.AutoMigrate(&Account{}); err != nil {
		return err
	}
	if err := db.AutoMigrate(&Post{}); err != nil {
		return err
	}
	return nil
}
