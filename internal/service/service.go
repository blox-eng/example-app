package service

import (
	"github.com/blox-eng/app/internal/storage"

	"github.com/blox-eng/app/internal/model"
)

type Service interface {
	ListUsers(user *model.User) ([]model.User, error)
	CreateUser(user *model.User) (model.UserData, error)
	GetUser(id string) (model.User, error)
	UpdateUser(id string, user *model.User) (model.UserData, error)
	DeleteUser(id string) (string, error)

	CreateSession(session *model.Session) (model.SessionData, error)
	GetUserFromSession(sessionID string, userID string) (*model.User, error)

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
