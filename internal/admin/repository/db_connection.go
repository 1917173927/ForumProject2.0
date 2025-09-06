package repository

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Accounts struct {
	UserID   int    `gorm:"primaryKey;column:user_id" json:"user_id"`
	Username string `gorm:"column:username" json:"username"`
	UserType int    `gorm:"column:user_type" json:"user_type"`
}

type Reporteds struct {
	gorm.Model
	UserID  int    `gorm:"column:user_id" json:"user_id"`
	PostID  int    `gorm:"column:post_id" json:"post_id"`
	Content string `gorm:"column:content" json:"content"`
	Reason  string `gorm:"column:reason" json:"reason"`
	Status  int    `gorm:"column:status" json:"status"`
}

func GetDBConnection() (*gorm.DB, error) {
	dsn := "root:coppklmja!BWZ@tcp(127.0.0.1:3306)/items?charset=utf8mb4&parseTime=True&loc=Local"
	return gorm.Open(mysql.Open(dsn), &gorm.Config{})
}
