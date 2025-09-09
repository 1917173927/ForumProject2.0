package services

import (
	"api-main/app/models"
	"api-main/config/database"
	"errors"
	"fmt"
)

type ReportRequest struct {
	PostID    uint   `json:"PostID"`
	UserID    uint   `json:"UserID"`
	Reason    string `json:"Reason"`
	Type      string `json:"Type"`
}

type ReportResponse struct {
	ID        uint   `json:"ID"`
	PostID    uint   `json:"PostID"`
	UserID    uint   `json:"UserID"`
	Reason    string `json:"Reason"`
	Status    int    `json:"Status"`
	Type      string `json:"Type"`
}

func CreateReport(req ReportRequest) (ReportResponse, error) {
	if req.PostID == 0 || req.UserID == 0 {
		return ReportResponse{}, errors.New("post_id and user_id are required")
	}

	if req.Reason == "" {
		return ReportResponse{}, errors.New("reason cannot be empty")
	}

	if req.Type == "" {
		req.Type = "other"
	}

	db := database.GetDB()

	report := models.Report{
		PostID:    req.PostID,
		UserID:    req.UserID,
		Reason:    req.Reason,
		Status:    0,
		Type:      req.Type,
	}

	fmt.Printf("Creating report with data: %+v\n", report)
	if err := db.Create(&report).Error; err != nil {
		fmt.Printf("Database error: %v\n", err)
		return ReportResponse{}, errors.New("failed to create report")
	}
	fmt.Printf("Report created successfully with ID: %d\n", report.ID)

	return ReportResponse{
		ID:        report.ID,
		PostID:    report.PostID,
		UserID:    report.UserID,
		Reason:    report.Reason,
		Status:    report.Status,
		Type:      report.Type,
	}, nil
}

func GetReports() ([]models.Report, error) {
	db := database.GetDB()

	var reports []models.Report
	if err := db.Find(&reports).Error; err != nil {
		return nil, errors.New("failed to get reports")
	}

	return reports, nil
}

func GetPendingReports() ([]models.Report, error) {
	db := database.GetDB()

	var reports []models.Report
	if err := db.Where("status = ?", 0).Find(&reports).Error; err != nil {
		return nil, errors.New("failed to get pending reports")
	}

	return reports, nil
}

func DeletePostByReportID(reportID uint) error {
	db := database.GetDB()

	// 获取举报对应的帖子ID
	var report models.Report
	if err := db.Unscoped().First(&report, reportID).Error; err != nil {
		return errors.New("report not found")
	}

	// 删除对应帖子
	if err := db.Delete(&models.Post{}, report.PostID).Error; err != nil {
		return errors.New("failed to delete post")
	}

	// 更新举报状态为已处理
	if err := db.Model(&models.Report{}).Where("id = ?", reportID).Update("status", 1).Error; err != nil {
		return errors.New("failed to update report status")
	}

	return nil
}

func UpdateReportStatus(reportID uint, status int) error {
	if status != 1 && status != 2 {
		return errors.New("invalid status, must be 1(approved) or 2(rejected)")
	}

	db := database.GetDB()

	if err := db.Model(&models.Report{}).Where("id = ?", reportID).Update("status", status).Error; err != nil {
		return errors.New("failed to update report status")
	}

	return nil
}
