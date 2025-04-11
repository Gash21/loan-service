package database

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Database struct {
	Gorm *gorm.DB
}

type SQLDatabaseOpts struct {
	Logger *zap.Logger
	DSN    string
}
