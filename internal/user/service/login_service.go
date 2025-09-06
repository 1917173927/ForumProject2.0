package service

import (
	"api-main/internal/user/repository"
	"errors"
	"log"

	"gorm.io/gorm"
)

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	UserID   int `json:"user_id"`
	UserType int `json:"user_type"`
}

type Account struct {
	Username string `gorm:"column:username"`
	Password string `gorm:"column:password"`
	UserID   int    `gorm:"column:user_id"`
	UserType int    `gorm:"column:user_type"`
}

var db *gorm.DB

func InitService() error {
	var err error
	db, err = repository.GetDBConnection()
	if err != nil {
		return errors.New("failed to initialize database connection")
	}
	return nil
}

func Login(req LoginRequest) (LoginResponse, error) {
	if db == nil {
		return LoginResponse{}, errors.New("database connection is not initialized")
	}

	var account Account
	if err := db.Where("username = ?", req.Username).First(&account).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Printf("Login failed: user '%s' not found", req.Username)
			return LoginResponse{}, errors.New("user not found")
		}
		log.Printf("Database error during login for user '%s': %v", req.Username, err)
		return LoginResponse{}, err
	}

	if account.Password != req.Password {
		log.Printf("Login failed: invalid password for user '%s'", req.Username)
		return LoginResponse{}, errors.New("invalid password")
	}

	return LoginResponse{
		UserID:   account.UserID,
		UserType: account.UserType,
	}, nil
}
