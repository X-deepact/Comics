package database

import (
	"be-client/config"
	"fmt"
	"log/slog"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DatabaseClient struct {
	db *gorm.DB
}

type Transaction struct {
	tx *gorm.DB
}

func NewDatabaseConnectClient(conf *config.DatabaseConfig) (*DatabaseClient, error) {
	dns := conf.BuildDnsMYSQL()
	slog.Info(fmt.Sprintf("connecting to database %s, %s", "url", dns))
	db, err := gorm.Open(mysql.Open(dns), &gorm.Config{
		IgnoreRelationshipsWhenMigrating: true,
		QueryFields:                      true,
		PrepareStmt:                      true,
		AllowGlobalUpdate:                true,
		Logger:                           logger.Default,
	})
	db = db.Set("gorm:table_options", "CHARSET=utf8mb4")
	db = db.Debug()
	if err != nil {
		slog.Error(fmt.Sprintf("failed to connect to database, %v", err))
		return nil, err
	}
	instance, err := db.DB()
	if err != nil {
		slog.Error(fmt.Sprintf("failed to connect to database, %v", err))
		return nil, err
	}
	instance.SetMaxIdleConns(conf.MaxIdleConnections)
	instance.SetMaxOpenConns(conf.MaxOpenConnections)
	instance.SetConnMaxIdleTime(conf.MaxConnIdleTime)
	return &DatabaseClient{
		db: db,
	}, nil
}

func (db *DatabaseClient) GetDB() *gorm.DB {
	return db.db
}
