package service

import (
	"github.com/blox-eng/app/internal/model"
	log "github.com/sirupsen/logrus"
)

func (s *service) CreateSession(session *model.Session) (model.SessionData, error) {
	record, err := s.sqlDB.CreateSession(session)
	if err != nil {
		log.Info("Failure: mot getting data from Table", err)
		return model.SessionData{}, err
	}
	return record, nil
}
func (s *service) GetUserFromSession(sessionID string, userID string) (*model.User, error) {
	record, err := s.sqlDB.GetUserFromSession(sessionID, userID)
	if err != nil {
		log.Info("Failure: mot getting data from Table", err)
		return &model.User{}, err
	}
	return record, nil
}
