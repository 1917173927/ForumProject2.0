package controller

import (
	"api-main/internal/admin/service"
	"api-main/internal/common/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ReportController struct{}

func (rc *ReportController) GetReportsHandler(c *gin.Context) {
	userID := c.Query("user_id")
	if userID == "" {
		response.JSON(c.Writer, http.StatusBadRequest, "Missing user_id parameter", nil)
		return
	}

	reports, err := service.GetReports(userID)
	if err != nil {
		response.JSON(c.Writer, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	response.JSON(c.Writer, http.StatusOK, "Success", reports)
}

type AdminReportController struct {
	Service *service.AdminService
}

func (arc *AdminReportController) GetPendingReportsHandler(c *gin.Context) {
	reports, err := arc.Service.GetPendingReports()
	if err != nil {
		response.JSON(c.Writer, http.StatusInternalServerError, "Failed to fetch pending reports", nil)
		return
	}

	response.JSON(c.Writer, http.StatusOK, "Success", reports)
}

func (arc *AdminReportController) ProcessReportApprovalHandler(c *gin.Context) {
	var req struct {
		ReportID int `json:"report_id"`
		Approval int `json:"approval"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		response.JSON(c.Writer, http.StatusBadRequest, "Invalid request body", nil)
		return
	}

	if err := arc.Service.ProcessReportApproval(req.ReportID, req.Approval); err != nil {
		response.JSON(c.Writer, http.StatusInternalServerError, "Failed to process report approval", nil)
		return
	}

	response.JSON(c.Writer, http.StatusOK, "Report processed successfully", nil)
}
