package storage

import (
	"errors"
	"fmt"

	"github.com/blox-eng/app/internal/hash"
	"github.com/blox-eng/app/internal/model"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func createUser(user *model.User) (model.UserData, error) {
	db := GetDBConnection()
	hasher := hash.NewHPasswordHash()
	hashedPassword, err := hasher.GenerateFromPassword(user.Password)

	if err != nil {
		log.Info("Failed to generate new hash", model.User{}, err)
	}

	user.Password = hashedPassword
	if err := db.Table("users").Create(&user).Error; err != nil {
		log.Info("failure", model.User{}, err)
	}

	record := model.UserData{
		User:    *user,
		Message: "User created",
	}
	return record, nil
}

func listUsers(user *model.User) ([]model.User, error) {
	db := GetDBConnection()
	var record []model.User
	db = db.Where(&user)
	if err := db.Table("users").Find(&record).Error; err != nil {
		log.Info("failure", []model.User{}, err)
	}

	if len(record) == 0 {
		return []model.User{}, fmt.Errorf("No users found")
	}

	return record, nil
}

func getUser(id string) (model.User, error) {
	var record []model.User
	db := GetDBConnection()
	if err := db.Table("users").Where("id=?", id).Find(&record).Error; err != nil {
		log.Info("failure", []model.User{})
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return model.User{}, fmt.Errorf("User with IDs %s not found", id)
		}
		return model.User{}, fmt.Errorf("failed to get user: %w", err)
	}
	if len(record) == 0 {
		return model.User{}, gorm.ErrRecordNotFound
	}
	return record[0], nil
}

func updateUser(id string, user *model.User) (model.UserData, error) {
	db := GetDBConnection()
	if err := db.Table("users").Where("id=?", id).Updates(&user).Error; err != nil {
		log.Info("failure", []model.User{})
	}
	record := model.UserData{
		User:    *user,
		Message: "record updated successfully",
	}
	return record, nil
}

func deleteUser(id string) (string, error) {
	db := GetDBConnection()
	if err := db.Table("users").Where("id=?", id).Delete(&model.User{}).Error; err != nil {
		log.Info("failure", []model.User{})
	}
	return "record deleted successfully", nil
}
