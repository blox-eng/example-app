package storage

import (
	"errors"
	"fmt"

	"github.com/blox-eng/app/internal/model"

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

func listWorkReports(wr *model.WorkReport) ([]model.WorkReport, error) {
	var record []model.WorkReport
	db := GetDBConnection()
	db = db.Where(&wr)
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

	result := db.Table("work_reports").Where("id=?", id).Delete(&wr)
	if result.Error != nil {
		return "not able to delete", result.Error
	}
	if result.RowsAffected == 0 {
		return "no matching record found", nil
	}

	return "deleted successfully", nil
}
