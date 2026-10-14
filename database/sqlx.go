package database

import (
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	baseDatabase "github.com/vucongthanh92/go-base-utils/database"
	"github.com/vucongthanh92/go-base-utils/logger"
	"github.com/vucongthanh92/go-base-utils/models"
)

type ReadDb *sqlx.DB
type WriteDb *sqlx.DB

func Open(cfg *models.DatabaseConfig) (ReadDb, WriteDb) {
	readDb := baseDatabase.MustConnect(cfg.ReadDbCfg.DbType, cfg.ReadDbCfg.ConnectionString)
	readDb.SetMaxIdleConns(cfg.ReadDbCfg.MaxIdleConns)
	readDb.SetMaxOpenConns(cfg.ReadDbCfg.MaxOpenConns)
	readDb.SetConnMaxLifetime(time.Duration(cfg.ReadDbCfg.ConnMaxLifetime) * time.Minute)
	readDb.MapperFunc(func(s string) string { return s })

	writeDb := baseDatabase.MustConnect(cfg.WriteDbCfg.DbType, cfg.WriteDbCfg.ConnectionString)
	writeDb.SetMaxIdleConns(cfg.WriteDbCfg.MaxIdleConns)
	writeDb.SetMaxOpenConns(cfg.WriteDbCfg.MaxOpenConns)
	writeDb.SetConnMaxLifetime(time.Duration(cfg.WriteDbCfg.ConnMaxLifetime) * time.Minute)
	writeDb.MapperFunc(func(s string) string { return s })

	logger.Info("Connected to read & write database!")
	return readDb, writeDb
}
