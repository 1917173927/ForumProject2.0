package models

import (
	"gorm.io/gorm"
	"sync"
)

var (
	db     *gorm.DB
	dbOnce sync.Once
)

type Account struct {
	UserID   uint   `gorm:"primaryKey;column:user_id" json:"UserID"`
	Username string `gorm:"column:username;not null" json:"Username" validate:"required"`
	Password string `gorm:"column:password;not null" json:"Password" validate:"required"`
	UserType int    `gorm:"column:user_type" json:"UserType"`
	Name     string `gorm:"column:name" json:"Name"`
}

func (Account) TableName() string {
	return "accounts"
}


