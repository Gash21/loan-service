package database

import (
	"go.uber.org/zap"
)

func NewDatabase(opts SQLDatabaseOpts) (*Database, error) {
	if opts.Logger == nil {
		opts.Logger = zap.NewNop()
	}

	connSqlite, err := InitSQLite(opts)
	if err != nil {
		opts.Logger.Error("Failed to connect to database", zap.Error(err))
		return nil, err
	}

	return &Database{
		Gorm: connSqlite,
	}, nil
}
