package service

import (
	"api-main/internal/admin/repository"
	"errors"

	"gorm.io/gorm"
)

type Report struct {
	ReportID int    `json:"report_id"`
	Username string `json:"username"`
	PostID   int    `json:"post_id"`
	Content  string `json:"content"`
	Reason   string `json:"reason"`
}

type AdminService struct {
	DB *gorm.DB
}

func (s *AdminService) GetPendingReports() ([]repository.Reported, error) {
	return repository.GetPendingReports(s.DB)
}

func (s *AdminService) ProcessReportApproval(reportID int, approval int) error {
	return repository.ProcessReportApproval(s.DB, reportID, approval)
}

func GetReports(userID string) ([]Report, error) {
	db, err := repository.GetDBConnection()
	if err != nil {
		return nil, errors.New("database connection error")
	}

	// 检查用户权限
	var account repository.Accounts
	if err := db.Where("user_id = ?", userID).First(&account).Error; err != nil {
		return nil, errors.New("user not found")
	}
	if account.UserType != 2 {
		return nil, errors.New("no admin privileges")
	}

	// 查询未审批的举报
	var reporteds []repository.Reporteds
	if err := db.Where("status = 0").Find(&reporteds).Error; err != nil {
		return nil, errors.New("failed to query reports")
	}

	reports := make([]Report, len(reporteds))
	for i, r := range reporteds {
		reports[i] = Report{
			ReportID: int(r.ID),
			Username: account.Username,
			PostID:   r.PostID,
			Content:  r.Content,
			Reason:   r.Reason,
		}
	}

	return reports, nil
}
