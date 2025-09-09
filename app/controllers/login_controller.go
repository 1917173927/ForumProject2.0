package controllers

import (
	"api-main/app/services"
	"api-main/app/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type LoginController struct{}

func (lc *LoginController) LoginHandler (c *gin.Context) {
	var req services.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.JsonErrorResponse(c, http.StatusBadRequest, "Invalid request body")
		return
	}

	resp, err := services.Login(req)
	if err != nil {
		utils.JsonErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	token, err := utils.GenerateToken(resp.UserID)
	if err != nil {
		utils.JsonErrorResponse(c, http.StatusInternalServerError, "Failed to generate token")
		return
	}

	utils.JsonSuccessResponse(c, gin.H{
		"token": token,
		"user": resp,
	})
}
