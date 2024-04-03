package storage

import (
	"errors"
	"fmt"

	"github.com/blox-eng/app/internal/model"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func createSession(session *model.Session) (model.SessionData, error) {
	db := GetDBConnection()
	session.SessionID = uuid.New().String()
	if err := db.Table("sessions").Create(&session).Error; err != nil {
		log.Info("failure", []model.Session{})
	}
	record := model.SessionData{
		Session: *session,
		Message: "Session created",
	}
	return record, nil
}

func getUserFromSession(sessionID string, userID string) (*model.User, error) {
	var record model.Session

	db := GetDBConnection()
	err := db.Preload("User", func(db *gorm.DB) *gorm.DB {
		return db.Select("ID", "Email")
	}).Where("session_id = ? AND user_id = ?", sessionID, userID).First(&record).Error

	if err != nil {
		log.Info("failure", []model.User{})
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &model.User{}, fmt.Errorf("User with IDs %s not found", userID)
		}
		return &model.User{}, fmt.Errorf("failed to get user: %w", err)
	}

	if record.User.ID == 0 {
		return &model.User{}, fmt.Errorf("No User associated with this session")
	}
	return &record.User, nil
}
