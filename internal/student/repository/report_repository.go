package repository

import (
	"errors"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Reported struct {
	gorm.Model
	UserID int    `gorm:"column:user_id" json:"user_id"`
	PostID int    `gorm:"column:post_id" json:"post_id"`
	Reason string `gorm:"column:reason" json:"reason"`
	Status int    `gorm:"column:status;default:0" json:"status"`
}

var db *gorm.DB

func InitDB() error {
	dsn := "root:coppklmja!BWZ@tcp(127.0.0.1:3306)/items?charset=utf8mb4&parseTime=True&loc=Local"
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return err
	}

	sqlDB, err := database.DB()
	if err != nil {
		return err
	}
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(10 * time.Second)

	db = database
	return db.AutoMigrate(&Reported{})
}

func SaveReport(report Reported) error {
	return db.Create(&report).Error
}

func SaveReportData(postID uint, reason string, userID uint) error {
	// Simulate saving the report to the database
	// Replace this with actual database logic
	if postID == 0 || reason == "" || userID == 0 {
		return errors.New("invalid report data")
	}

	report := Reported{
		UserID: int(userID),
		PostID: int(postID),
		Reason: reason,
	}

	// Assume the report is saved successfully
	return SaveReport(report)
}

func GetReportsByUserID(userID int) ([]Reported, error) {
	var reports []Reported
	if err := db.Where("user_id = ?", userID).Find(&reports).Error; err != nil {
		return nil, err
	}
	return reports, nil
}
