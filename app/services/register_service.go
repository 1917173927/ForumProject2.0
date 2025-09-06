package services

import (
	"api-main/app/models"
	"api-main/config/database"
	"errors"
)

type RegisterRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	UserType int    `json:"user_type"`
	Name     string `json:"name"`
}

type RegisterResponse struct {
	Username string `json:"username"`
	Password string `json:"password"`
	UserType int    `json:"user_type"`
	Name     string `json:"name"`
}

func Register(req RegisterRequest) (RegisterResponse, error) {
	if req.Username == "" || req.Password == "" || req.Name == "" || (req.UserType != 1 && req.UserType != 2) {
		return RegisterResponse{}, errors.New("missing or invalid fields")
	}

	db, err := database.GetDBConnection()
	if err != nil {
		return RegisterResponse{}, errors.New("database connection error")
	}

	account := models.Account{
		Username: req.Username,
		Password: req.Password,
		UserType: req.UserType,
		Name:     req.Name,
	}

	if err := db.Create(&account).Error; err != nil {
		return RegisterResponse{}, errors.New("failed to create user")
	}

	return RegisterResponse{
		Username: account.Username,
		Password: account.Password,
		UserType: account.UserType,
		Name:     account.Name,
	}, nil
}
