package services

import (
	"api-main/app/models"
	"api-main/config/database"
	"errors"
	"fmt"
)

type ReportRequest struct {
	PostID    uint   `json:"post_id"`
	ReporterID uint   `json:"reporter_id"`
	Reason    string `json:"reason"`
	Type      string `json:"type"`
}

type ReportResponse struct {
	ID        uint   `json:"id"`
	PostID    uint   `json:"post_id"`
	ReporterID uint   `json:"reporter_id"`
	Reason    string `json:"reason"`
	Status    int    `json:"status"`
	Type      string `json:"type"`
}

func CreateReport(req ReportRequest) (ReportResponse, error) {
	if req.PostID == 0 || req.ReporterID == 0 {
		return ReportResponse{}, errors.New("post_id and reporter_id are required")
	}

	if req.Reason == "" {
		return ReportResponse{}, errors.New("reason cannot be empty")
	}

	if req.Type == "" {
		req.Type = "other"
	}

	db, err := database.GetDBConnection()
	if err != nil {
		return ReportResponse{}, errors.New("database connection error")
	}

	report := models.Report{
		PostID:    req.PostID,
		ReporterID: req.ReporterID,
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
		ReporterID: report.ReporterID,
		Reason:    report.Reason,
		Status:    report.Status,
		Type:      report.Type,
	}, nil
}

func GetReports() ([]models.Report, error) {
	db, err := database.GetDBConnection()
	if err != nil {
		return nil, errors.New("database connection error")
	}

	var reports []models.Report
	if err := db.Find(&reports).Error; err != nil {
		return nil, errors.New("failed to get reports")
	}

	return reports, nil
}

func GetPendingReports() ([]models.Report, error) {
	db, err := database.GetDBConnection()
	if err != nil {
		return nil, errors.New("database connection error")
	}

	var reports []models.Report
	if err := db.Where("status = ?", 0).Find(&reports).Error; err != nil {
		return nil, errors.New("failed to get pending reports")
	}

	return reports, nil
}

func DeletePostByReportID(reportID uint) error {
	db, err := database.GetDBConnection()
	if err != nil {
		return errors.New("database connection error")
	}

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

	db, err := database.GetDBConnection()
	if err != nil {
		return errors.New("database connection error")
	}

	if err := db.Model(&models.Report{}).Where("id = ?", reportID).Update("status", status).Error; err != nil {
		return errors.New("failed to update report status")
	}

	return nil
}
