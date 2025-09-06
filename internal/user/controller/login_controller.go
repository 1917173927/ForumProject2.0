package controller

import (
	"api-main/internal/common/response"
	"api-main/internal/user/service"

	"github.com/gin-gonic/gin"
)

type LoginController struct{}

func (lc *LoginController) LoginHandler(c *gin.Context) {
	var req service.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.JSON(c.Writer, 400, "Invalid request body", nil)
		return
	}

	resp, err := service.Login(req)
	if err != nil {
		response.JSON(c.Writer, 401, err.Error(), nil)
		return
	}

	response.JSON(c.Writer, 200, "Success", resp)
}
