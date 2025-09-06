package controller

import (
	"api-main/internal/common/response"
	"api-main/internal/user/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RegisterController struct{}

func (rc *RegisterController) RegisterHandler(c *gin.Context) {
	var req service.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.JSON(c.Writer, http.StatusBadRequest, "Invalid request body", nil)
		return
	}

	// Ensure the Name field is present in the request
	if req.Name == "" {
		response.JSON(c.Writer, http.StatusBadRequest, "Name field is required", nil)
		return
	}

	resp, err := service.Register(req)
	if err != nil {
		response.JSON(c.Writer, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	response.JSON(c.Writer, http.StatusOK, "User registered successfully", resp)
}
