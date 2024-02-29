package db

import (
	"errors"
	"fmt"

	"github.com/blox-eng/work/model"

	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
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

func getAllWorkReports(id string) (model.WorkReport, error) {
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
		log.Info("failure", []model.Blogs{})
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
	if err := db.Table("work_report").Where("id=?", id).Delete(&wr).Error; err != nil {
		log.Info("failure", []model.WorkReport{})
		return "not able to delete", err
	}

	return "deleted successfully", nil
}
func createBlog(bl *model.Blogs) (model.BlogData, error) {
	db := GetDBConnection()
	if err := db.Table("blogs").Create(&bl).Error; err != nil {
		log.Info("failure", model.Blogs{}, err)
	}
	record := model.BlogData{
		Blog:    *bl,
		Message: "data saved",
	}
	return record, nil

}

func getAllBlog(id string) (model.Blogs, error) {
	var record []model.Blogs
	db := GetDBConnection()
	if err := db.Table("blogs").Where("id=?", id).Find(&record).Error; err != nil {
		log.Info("failure", []model.Blogs{})
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return model.Blogs{}, fmt.Errorf("blog with IDs %s not found", id)
		}
		return model.Blogs{}, fmt.Errorf("failed to get blog: %w", err)
	}
	if len(record) == 0 {
		return model.Blogs{}, gorm.ErrRecordNotFound
	}
	return record[0], nil
}

func updateBlog(id string, bl *model.Blogs) (model.BlogData, error) {

	db := GetDBConnection()
	if err := db.Table("blogs").Where("id=?", id).Updates(&bl).Error; err != nil {
		log.Info("failure", []model.Blogs{})
	}
	record := model.BlogData{
		Blog:    *bl,
		Message: "record updated successfully",
	}
	return record, nil
}

func deleteBlog(id string) (string, error) {
	var bl model.Blogs
	db := GetDBConnection()
	if err := db.Table("blogs").Where("id=?", id).Delete(&bl).Error; err != nil {
		log.Info("failure", []model.Blogs{})
		return "not able to delete", err
	}

	return "deleted successfully", nil
}
