package storage

import (
	"github.com/blox-eng/backend/internal/model"
)

type SqlClient interface {
	GetAllWorkReports() ([]model.WorkReport, error)
	CreateWorkReport(wr *model.WorkReport) (model.WorkReportData, error)
	GetWorkReport(id string) (model.WorkReport, error)
	UpdateWorkReport(id string, wr *model.WorkReport) (model.WorkReportData, error)
	DeleteWorkReport(id string) (string, error)
}

func (c *sqlClient) CreateWorkReport(wr *model.WorkReport) (model.WorkReportData, error) {
	return createWorkReport(wr)
}

func (c *sqlClient) GetAllWorkReports() ([]model.WorkReport, error) {
	return getAllWorkReports()
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
