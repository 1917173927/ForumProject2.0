package repository

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Post struct {
	PostID  int    `gorm:"primaryKey;column:post_id"`
	Content string `gorm:"column:content"`
	UserID  int    `gorm:"column:user_id"`
}

func GetDBConnection() (*gorm.DB, error) {
	dsn := "root:coppklmja!BWZ@tcp(127.0.0.1:3306)/items?charset=utf8mb4&parseTime=True&loc=Local"
	return gorm.Open(mysql.Open(dsn), &gorm.Config{})
}
