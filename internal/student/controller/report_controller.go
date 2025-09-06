package controller

import (
	"api-main/internal/common/response"
	"api-main/internal/student/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ReportController struct{}

func (rc *ReportController) ReportPostHandler(c *gin.Context) {
	var req service.ReportPostRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.JSON(c.Writer, http.StatusBadRequest, "Invalid JSON format", nil)
		return
	}

	err := service.ReportPost(req)
	if err != nil {
		response.JSON(c.Writer, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	response.JSON(c.Writer, http.StatusOK, "Post reported successfully", nil)
}

func (rc *ReportController) GetReportStatusHandler(c *gin.Context) {
	userID, err := strconv.Atoi(c.Query("user_id"))
	if err != nil || userID == 0 {
		response.JSON(c.Writer, http.StatusBadRequest, "Invalid or missing user_id", nil)
		return
	}

	statuses, err := service.GetReportStatus(userID)
	if err != nil {
		response.JSON(c.Writer, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	response.JSON(c.Writer, http.StatusOK, "Success", statuses)
}
