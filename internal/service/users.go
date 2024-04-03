package service

import (
	"github.com/blox-eng/app/internal/model"
	log "github.com/sirupsen/logrus"
)

func (s *service) ListUsers(user *model.User) ([]model.User, error) {
	records, err := s.sqlDB.ListUsers(user)
	if err != nil {
		log.Info("Failure: not getting data from Table", err)
		return []model.User{}, err
	}
	return records, nil
}
func (s *service) GetUser(id string) (model.User, error) {
	record, err := s.sqlDB.GetUser(id)
	if err != nil {
		log.Info("Failure: not getting data from Table", err)
		return model.User{}, err
	}
	return record, nil
}

func (s *service) CreateUser(user *model.User) (model.UserData, error) {
	record, err := s.sqlDB.CreateUser(user)
	if err != nil {
		log.Info("Failure: mot getting data from Table", err)
		return model.UserData{}, err
	}
	return record, nil
}

func (s *service) UpdateUser(id string, user *model.User) (model.UserData, error) {
	record, err := s.sqlDB.UpdateUser(id, user)
	if err != nil {
		log.Info("Failure: mot getting data from Table", err)
		return model.UserData{}, err
	}
	return record, nil
}

func (s *service) DeleteUser(id string) (string, error) {
	record, err := s.sqlDB.DeleteUser(id)
	if err != nil {
		log.Info("Failure: not getting data from Table", err)
		return "record is not available in the system", err
	}
	return record, nil
}
