package database

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
)

const ContextName = "Components.DB"

type Database struct {
	Gorm *gorm.DB
}

type SQLDatabaseOpts struct {
	GormConfig gorm.Config

	Debug  bool
	Logger *zap.Logger
	DBName string
}
