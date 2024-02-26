package db

import (
    "fmt"

    "github.com/blox-eng/work/config"
    "github.com/blox-eng/work/model"

    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/postgres"
    log "github.com/sirupsen/logrus"
)

type sqlConn struct {
    DbPool *gorm.DB
}

var connector *sqlConn

func InitializeDatabaseConnector() *sqlConn {
    if connector != nil {
        log.Info("DataBase is initialized")
        return connector
    }
    log.Info("DataBase was not initialized ..initializing again")
    var err error
    connector, err = initializeDatabaseConnector()
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

    db, err := gorm.Open("postgres", dbUri)
    if err != nil {
        panic(err)
    }
    if maxCons := config.GetYamlValues().DBConfig.MaxConnection; maxCons > 0 {
        db.DB().SetMaxOpenConns(maxCons)
        db.DB().SetMaxIdleConns(maxCons / 3)
    }
    log.Info(db.AutoMigrate(&model.Blogs{}))   
    return &sqlConn{db}, nil
}

func GetDBConnection() *gorm.DB {
    return connector.DbPool
}
