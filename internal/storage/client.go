package storage

import (
	"fmt"

	"github.com/blox-eng/backend/config"
	"github.com/blox-eng/backend/internal/model"

	log "github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Config struct {
	DBConnection string
}

type sqlClient struct {
	config *Config
}

func NewClient(config *Config) SqlClient {
	return &sqlClient{
		config: config,
	}
}

type sqlConn struct {
	DbPool *gorm.DB
}

var connector *sqlConn

func InitializeDatabaseConnector() *sqlConn {
	if connector != nil {
		log.Info("Database is initialized")
		return connector
	}
	log.Info("Database was not initialized ...initializing again")
	var err error
	connector, err = initializeDatabaseConnector()
	log.Info("Connection established: ", &connector)
	if err != nil {
		panic(err)
	}
	return connector
}

// DB Initialization

func initializeDatabaseConnector() (*sqlConn, error) {
	log.Info(config.GetYamlValues().DBConfig, config.GetYamlValues().DBConfig.Port)
	dbUri := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s",
		config.GetYamlValues().DBConfig.Server, config.GetYamlValues().DBConfig.Username, config.GetYamlValues().DBConfig.Schema, config.GetYamlValues().DBConfig.Password) //Build connection string

	db, err := gorm.Open(postgres.Open(dbUri), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}
	if maxCons := config.GetYamlValues().DBConfig.MaxConnection; maxCons > 0 {
		sqlDB.SetMaxOpenConns(maxCons)
		sqlDB.SetMaxIdleConns(maxCons / 3)
	}
	// https://gorm.io/docs/migration.html - unused columns WILL NOT BE DELETED
	db.AutoMigrate(&model.WorkReport{})
	return &sqlConn{db}, nil
}

func GetDBConnection() *gorm.DB {
	return connector.DbPool
}
