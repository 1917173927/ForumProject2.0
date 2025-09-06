package controllers

import (
	"api-main/app/services"
	"api-main/app/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type LoginController struct{}

func (lc *LoginController) LoginHandler(c *gin.Context) {
	var req services.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.JSON(c.Writer, http.StatusBadRequest, "Invalid request body", nil)
		return
	}

	resp, err := services.Login(req)
	if err != nil {
		utils.JSON(c.Writer, http.StatusUnauthorized, err.Error(), nil)
		return
	}

	utils.JSON(c.Writer, http.StatusOK, "Login successful", resp)
}
