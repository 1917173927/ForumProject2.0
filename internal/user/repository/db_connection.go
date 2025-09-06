package repository

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Account struct {
	Username string `gorm:"column:username"`
	Password string `gorm:"column:password"`
	UserType int    `gorm:"column:user_type"`
	Name     string `gorm:"column:name"`
}

func GetDBConnection() (*gorm.DB, error) {
	dsn := "root:coppklmja!BWZ@tcp(127.0.0.1:3306)/items?charset=utf8mb4&parseTime=True&loc=Local"
	return gorm.Open(mysql.Open(dsn), &gorm.Config{})
}
