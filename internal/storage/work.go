package storage

import (
	"errors"
	"fmt"

	"github.com/blox-eng/backend/internal/model"

	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func createWorkReport(wr *model.WorkReport) (model.WorkReportData, error) {
	db := GetDBConnection()
	if err := db.Table("work_reports").Create(&wr).Error; err != nil {
		log.Info("failure", model.WorkReport{}, err)
	}

	record := model.WorkReportData{
		WorkReport: *wr,
		Message:    "Work Report saved",
	}

	return record, nil
}

func getAllWorkReports() ([]model.WorkReport, error) {
	var record []model.WorkReport
	db := GetDBConnection()
	if err := db.Table("work_reports").Find(&record).Error; err != nil {
		log.Info("failure", []model.WorkReport{})
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return []model.WorkReport{}, fmt.Errorf("No work reports found")
		}
		return []model.WorkReport{}, fmt.Errorf("failed to get work reports: %w", err)
	}
	if len(record) == 0 {
		return []model.WorkReport{}, nil
	}
	return record, nil
}

func getWorkReport(id string) (model.WorkReport, error) {
	var record []model.WorkReport
	db := GetDBConnection()
	if err := db.Table("work_reports").Where("id=?", id).Find(&record).Error; err != nil {
		log.Info("failure", []model.WorkReport{})
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return model.WorkReport{}, fmt.Errorf("Work report with IDs %s not found", id)
		}
		return model.WorkReport{}, fmt.Errorf("failed to get work report: %w", err)
	}
	if len(record) == 0 {
		return model.WorkReport{}, gorm.ErrRecordNotFound
	}
	return record[0], nil
}

func updateWorkReport(id string, wr *model.WorkReport) (model.WorkReportData, error) {

	db := GetDBConnection()
	if err := db.Table("work_reports").Where("id=?", id).Updates(&wr).Error; err != nil {
		log.Info("failure", []model.WorkReport{})
	}
	record := model.WorkReportData{
		WorkReport: *wr,
		Message:    "record updated successfully",
	}
	return record, nil
}

func deleteWorkReport(id string) (string, error) {
	var wr model.WorkReport
	db := GetDBConnection()
	if err := db.Table("work_reports").Where("id=?", id).Delete(&wr).Error; err != nil {
		log.Info("failure", []model.WorkReport{})
		return "not able to delete", err
	}

	return "deleted successfully", nil
}
