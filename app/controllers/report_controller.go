package controllers

import (
	"fmt"
	"api-main/app/models"
	"api-main/app/services"
	"api-main/app/utils"
	"api-main/config/database"
	"net/http"
	"github.com/gin-gonic/gin"
)

type ReportController struct{}

func (rc *ReportController) CreateReportHandler(c *gin.Context) {
	var req services.ReportRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.JSON(c.Writer, http.StatusBadRequest, "Invalid request body", nil)
		return
	}

	resp, err := services.CreateReport(req)
	if err != nil {
		utils.JSON(c.Writer, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	utils.JSON(c.Writer, http.StatusOK, "Report created successfully", resp)
}

func (rc *ReportController) GetReportsHandler(c *gin.Context) {
	reports, err := services.GetReports()
	if err != nil {
		utils.JSON(c.Writer, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	utils.JSON(c.Writer, http.StatusOK, "Success", gin.H{"reports": reports})
}

func (rc *ReportController) GetPendingReportsHandler(c *gin.Context) {
	reports, err := services.GetPendingReports()
	if err != nil {
		utils.JSON(c.Writer, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	utils.JSON(c.Writer, http.StatusOK, "Success", gin.H{"reports": reports})
}

func (rc *ReportController) ReviewReportHandler(c *gin.Context) {
	var req struct {
		ReportID uint `json:"report_id"`
		Approval int  `json:"approval"`
		UserID   uint `json:"user_id"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		utils.JSON(c.Writer, http.StatusBadRequest, "Invalid request body", nil)
		return
	}

	// 验证用户是否为管理员
	db, err := database.GetDBConnection()
	if err != nil {
		utils.JSON(c.Writer, http.StatusInternalServerError, "Failed to connect to database", nil)
		return
	}

	fmt.Printf("Verifying user with ID: %d\n", req.UserID)
	var account models.Account
	if err := db.Unscoped().Where("user_id = ?", req.UserID).Order("user_id").First(&account).Error; err != nil {
		fmt.Printf("Error verifying user: %v\n", err)
		utils.JSON(c.Writer, http.StatusInternalServerError, "Failed to verify user", nil)
		return
	}

	if account.UserType != 2 {
		utils.JSON(c.Writer, http.StatusForbidden, "Only admin can review reports", nil)
		return
	}

	// 审核处理
	if req.Approval == 1 {
		// 通过审核，删除对应帖子
		if err := services.DeletePostByReportID(req.ReportID); err != nil {
			utils.JSON(c.Writer, http.StatusInternalServerError, err.Error(), nil)
			return
		}
	} else if req.Approval == 2 {
		// 驳回举报，仅更新举报状态
		if err := services.UpdateReportStatus(req.ReportID, 2); err != nil {
			utils.JSON(c.Writer, http.StatusInternalServerError, err.Error(), nil)
			return
		}
	} else {
		utils.JSON(c.Writer, http.StatusBadRequest, "Invalid approval status", nil)
		return
	}

	utils.JSON(c.Writer, http.StatusOK, "Report reviewed successfully", nil)
}
