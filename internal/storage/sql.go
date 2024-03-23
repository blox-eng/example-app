package storage

import (
	"github.com/blox-eng/app/internal/model"
)

type SqlClient interface {
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
