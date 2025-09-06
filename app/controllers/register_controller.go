package controllers

import (
	"api-main/app/services"
	"api-main/app/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RegisterController struct{}

func (rc *RegisterController) RegisterHandler(c *gin.Context) {
	var req services.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.JSON(c.Writer, http.StatusBadRequest, "Invalid request body", nil)
		return
	}

	if req.Name == "" {
		utils.JSON(c.Writer, http.StatusBadRequest, "Name field is required", nil)
		return
	}

	resp, err := services.Register(req)
	if err != nil {
		utils.JSON(c.Writer, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	utils.JSON(c.Writer, http.StatusOK, "User registered successfully", resp)
}
