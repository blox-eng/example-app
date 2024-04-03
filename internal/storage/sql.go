package storage

import (
	"github.com/blox-eng/app/internal/model"
)

type SqlClient interface {
	// User
	ListUsers(user *model.User) ([]model.User, error)
	CreateUser(user *model.User) (model.UserData, error)
	GetUser(id string) (model.User, error)
	UpdateUser(id string, user *model.User) (model.UserData, error)
	DeleteUser(id string) (string, error)

	// Session
	CreateSession(session *model.Session) (model.SessionData, error)
	GetUserFromSession(sessionID string, userID string) (*model.User, error)

	// WorkReport
	ListWorkReports(wr *model.WorkReport) ([]model.WorkReport, error)
	CreateWorkReport(wr *model.WorkReport) (model.WorkReportData, error)
	GetWorkReport(id string) (model.WorkReport, error)
	UpdateWorkReport(id string, wr *model.WorkReport) (model.WorkReportData, error)
	DeleteWorkReport(id string) (string, error)
}

func (c *sqlClient) CreateWorkReport(wr *model.WorkReport) (model.WorkReportData, error) {
	return createWorkReport(wr)
}

func (c *sqlClient) ListWorkReports(wr *model.WorkReport) ([]model.WorkReport, error) {
	return listWorkReports(wr)
}

func (c *sqlClient) GetWorkReport(id string) (model.WorkReport, error) {
	return getWorkReport(id)
}

func (c *sqlClient) UpdateWorkReport(id string, wr *model.WorkReport) (model.WorkReportData, error) {
	return updateWorkReport(id, wr)
}

func (c *sqlClient) DeleteWorkReport(id string) (string, error) {
	return deleteWorkReport(id)
}

func (c *sqlClient) CreateUser(user *model.User) (model.UserData, error) {
	return createUser(user)
}

func (c *sqlClient) GetUser(id string) (model.User, error) {
	return getUser(id)
}

func (c *sqlClient) ListUsers(user *model.User) ([]model.User, error) {
	return listUsers(user)
}

func (c *sqlClient) UpdateUser(id string, user *model.User) (model.UserData, error) {
	return updateUser(id, user)
}

func (c *sqlClient) DeleteUser(id string) (string, error) {
	return deleteUser(id)
}

func (c *sqlClient) CreateSession(session *model.Session) (model.SessionData, error) {
	return createSession(session)
}

func (c *sqlClient) GetUserFromSession(sessionID string, userID string) (*model.User, error) {
	return getUserFromSession(sessionID, userID)
}
