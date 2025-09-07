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
		utils.JsonErrorResponse(c, http.StatusBadRequest, "Invalid request body")
		return
	}

	resp, err := services.CreateReport(req)
	if err != nil {
		utils.JsonErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.JsonSuccessResponse(c, resp)
}

func (rc *ReportController) GetReportsHandler(c *gin.Context) {
	reports, err := services.GetReports()
	if err != nil {
		utils.JsonErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.JsonSuccessResponse(c, gin.H{"reports": reports})
}

func (rc *ReportController) GetPendingReportsHandler(c *gin.Context) {
	reports, err := services.GetPendingReports()
	if err != nil {
		utils.JsonErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.JsonSuccessResponse(c, gin.H{"reports": reports})
}

func (rc *ReportController) ReviewReportHandler(c *gin.Context) {
	var req struct {
		ReportID uint `json:"report_id"`
		Approval int  `json:"approval"`
		UserID   uint `json:"user_id"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		utils.JsonErrorResponse(c, http.StatusBadRequest, "Invalid request body")
		return
	}

	// 验证用户是否为管理员
	db, err := database.GetDBConnection()
	if err != nil {
		utils.JsonErrorResponse(c, http.StatusInternalServerError, "Failed to connect to database")
		return
	}

	fmt.Printf("Verifying user with ID: %d\n", req.UserID)
	var account models.Account
	if err := db.Unscoped().Where("user_id = ?", req.UserID).Order("user_id").First(&account).Error; err != nil {
		fmt.Printf("Error verifying user: %v\n", err)
		utils.JsonErrorResponse(c, http.StatusInternalServerError, "Failed to verify user")
		return
	}

	if account.UserType != 2 {
		utils.JsonErrorResponse(c, http.StatusForbidden, "Only admin can review reports")
		return
	}

	// 审核处理
	if req.Approval == 1 {
		// 通过审核，删除对应帖子
		if err := services.DeletePostByReportID(req.ReportID); err != nil {
			utils.JsonErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}
	} else if req.Approval == 2 {
		// 驳回举报，仅更新举报状态
		if err := services.UpdateReportStatus(req.ReportID, 2); err != nil {
			utils.JsonErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}
	} else {
		utils.JsonErrorResponse(c, http.StatusBadRequest, "Invalid approval status")
		return
	}

	utils.JsonSuccessResponse(c, nil)
}
