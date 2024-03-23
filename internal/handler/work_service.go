package handler

import (
	"github.com/blox-eng/app/internal/model"

	log "github.com/sirupsen/logrus"
)

func (s *service) ListWorkReports(wr *model.WorkReport) ([]model.WorkReport, error) {
	records, err := s.sqlDB.ListWorkReports(wr)
	if err != nil {
		log.Info("Failure: not getting data from Table", err)
		return []model.WorkReport{}, err
	}
	return records, nil
}

func (s *service) GetWorkReport(id string) (model.WorkReport, error) {
	record, err := s.sqlDB.GetWorkReport(id)
	if err != nil {
		log.Info("Failure: not getting data from Table", err)
		return model.WorkReport{}, err
	}
	return record, nil
}

func (s *service) CreateWorkReport(wr *model.WorkReport) (model.WorkReportData, error) {
	record, err := s.sqlDB.CreateWorkReport(wr)
	if err != nil {
		log.Info("Failure: mot getting data from Table", err)
		return model.WorkReportData{}, err
	}
	return record, nil
}

func (s *service) UpdateWorkReport(id string, wr *model.WorkReport) (model.WorkReportData, error) {
	record, err := s.sqlDB.UpdateWorkReport(id, wr)
	if err != nil {
		log.Info("Failure: mot getting data from Table", err)
		return model.WorkReportData{}, err
	}
	return record, nil
}

func (s *service) DeleteWorkReport(id string) (string, error) {
	record, err := s.sqlDB.DeleteWorkReport(id)
	if err != nil {
		log.Info("Failure: not getting data from Table", err)
		return "record is not available in the system", err
	}
	return record, nil

}
