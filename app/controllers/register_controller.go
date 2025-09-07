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
		utils.JsonErrorResponse(c, http.StatusBadRequest, "Invalid request body")
		return
	}

	if req.Name == "" {
		utils.JsonErrorResponse(c, http.StatusBadRequest, "Name field is required")
		return
	}

	resp, err := services.Register(req)
	if err != nil {
		utils.JsonErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.JsonSuccessResponse(c, resp)
}
