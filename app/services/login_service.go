package services

import (
	"api-main/app/models"
	"api-main/config/database"
	"errors"
	"log"

	"gorm.io/gorm"
)

type LoginRequest struct {
	Username string `json:"Username"`
	Password string `json:"Password"`
}

type LoginResponse struct {
	UserID   uint   `json:"UserID"`
	Username string `json:"Username"`
	UserType int    `json:"UserType"`
	Name     string `json:"Name"`
}

func Login(req LoginRequest) (LoginResponse, error) {
	if req.Username == "" || req.Password == "" {
		return LoginResponse{}, errors.New("username and password are required")
	}

	db := database.GetDB()

	var account models.Account
	if err := db.Where("username = ?", req.Username).First(&account).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Printf("Login failed: user '%s' not found", req.Username)
			return LoginResponse{}, errors.New("user not found")
		}
		log.Printf("Database error during login: %v", err)
		return LoginResponse{}, errors.New("database error")
	}

	if account.Password != req.Password {
		log.Printf("Login failed: invalid password for user '%s'", req.Username)
		return LoginResponse{}, errors.New("invalid password")
	}

	return LoginResponse{
		UserID:   account.UserID,
		Username: account.Username,
		UserType: account.UserType,
		Name:     account.Name,
	}, nil
}
