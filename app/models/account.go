package models

type Account struct {
	UserID   uint   `gorm:"primaryKey;column:user_id" json:"user_id"`
	Username string `gorm:"column:username" json:"username"`
	Password string `gorm:"column:password" json:"password"`
	UserType int    `gorm:"column:user_type" json:"user_type"`
	Name     string `gorm:"column:name" json:"name"`
}

func (Account) TableName() string {
	return "accounts"
}
