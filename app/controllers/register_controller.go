package controllers

import (
	"api-main/app/services"
	"api-main/app/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RegisterController struct{}

// @Summary User registration
// @Description Register a new user
// @Tags auth
// @Accept json
// @Produce json
// @Param request body services.RegisterRequest true "Register request"
// @Success 200 {object} services.RegisterResponse
// @Failure 400 {object} utils.ErrorResponse
// @Failure 500 {object} utils.ErrorResponse
// @Router /register [post]
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
