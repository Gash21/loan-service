package database

import (
	"github.com/Gash21/amartha-test/internal/shared/logger"
	"go.uber.org/zap"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	gormlogger "gorm.io/gorm/logger"
)

func InitSQLite(opts SQLDatabaseOpts) (*gorm.DB, error) {
	l := logger.WithId(opts.Logger, ContextName, "Init Sqlite")
	l.Info("Initializing SQLite database")

	gormOpts := &opts.GormConfig
	gormOpts.PrepareStmt = true

	queryLogger := logger.WithId(opts.Logger, ContextName, "ExecuteQuery")

	if opts.Debug {
		gormOpts.Logger = NewGormLogger(queryLogger, gormlogger.Info, true)
	} else {
		gormOpts.Logger = NewGormLogger(queryLogger, gormlogger.Warn, false)
	}

	db, err := gorm.Open(sqlite.Open(opts.DBName), gormOpts)
	if err != nil {
		l.Error("Failed to connect to database", zap.Error(err))
		return nil, err
	}

	l.Info("Successfully connected to database", zap.String("db_name", opts.DBName))
	return db, nil
}
