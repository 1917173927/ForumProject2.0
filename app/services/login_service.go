package services

import (
	"api-main/app/models"
	"api-main/config/database"
	"errors"
	"log"

	"gorm.io/gorm"
)

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Username string `json:"username"`
	UserType int    `json:"user_type"`
	Name     string `json:"name"`
}

func Login(req LoginRequest) (LoginResponse, error) {
	if req.Username == "" || req.Password == "" {
		return LoginResponse{}, errors.New("username and password are required")
	}

	db, err := database.GetDBConnection()
	if err != nil {
		return LoginResponse{}, errors.New("database connection error")
	}

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
		Username: account.Username,
		UserType: account.UserType,
		Name:     account.Name,
	}, nil
}
