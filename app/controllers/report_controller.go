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
		ReportID uint `json:"ReportID"`
		Approval int  `json:"Approval"`
		UserID   uint `json:"UserID"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		utils.JsonErrorResponse(c, http.StatusBadRequest, "Invalid request body")
		return
	}

	// 验证用户是否为管理员
	db := database.GetDB()


	fmt.Printf("Verifying user with ID: %d\n", req.UserID)
	var account models.Account
	if err := db.Where("user_id = ?", req.UserID).First(&account).Error; err != nil {
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
		// 通过审核，删除对应帖子并更新举报状态
		if err := services.DeletePostByReportID(req.ReportID); err != nil {
			utils.JsonErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}
		if err := services.UpdateReportStatus(req.ReportID, 1); err != nil {
			utils.JsonErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}
	} else if req.Approval == 2 {
		// 驳回举报，更新举报状态并记录原因
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
