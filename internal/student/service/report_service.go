package service

import (
	"api-main/internal/student/repository"
	"errors"
)

type ReportPostRequest struct {
	UserID int    `json:"user_id"`
	PostID int    `json:"post_id"`
	Reason string `json:"reason"`
}

func ReportPost(req ReportPostRequest) error {
	if req.UserID == 0 || req.PostID == 0 || req.Reason == "" {
		return errors.New("invalid input data")
	}

	report := repository.Reported{
		UserID: req.UserID,
		PostID: req.PostID,
		Reason: req.Reason,
	}

	return repository.SaveReport(report)
}

type ReportStatusResponse struct {
	PostID int    `json:"post_id"`
	Reason string `json:"reason"`
	Status int    `json:"status"`
}

func GetReportStatus(userID int) ([]ReportStatusResponse, error) {
	reports, err := repository.GetReportsByUserID(userID)
	if err != nil {
		return nil, err
	}

	response := make([]ReportStatusResponse, 0)
	for _, report := range reports {
		response = append(response, ReportStatusResponse{
			PostID: report.PostID,
			Reason: report.Reason,
			Status: report.Status,
		})
	}

	return response, nil
}
