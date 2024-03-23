package handler

import (
	"github.com/blox-eng/app/internal/storage"

	"github.com/blox-eng/app/internal/model"
)

type Service interface {
	ListWorkReports(wr *model.WorkReport) ([]model.WorkReport, error)
	GetWorkReport(id string) (model.WorkReport, error)
	CreateWorkReport(wr *model.WorkReport) (model.WorkReportData, error)
	UpdateWorkReport(id string, wr *model.WorkReport) (model.WorkReportData, error)
	DeleteWorkReport(id string) (string, error)
}

func NewService(sqlDB storage.SqlClient) Service {
	return &service{
		sqlDB: sqlDB,
	}
}

type service struct {
	sqlDB storage.SqlClient
}
