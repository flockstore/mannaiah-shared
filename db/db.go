// Package db provides a reusable GORM connector with Zap integration via zapgorm2.
package db

import (
	"fmt"
	"strings"

	"github.com/flockstore/mannaiah-shared/logger"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
	"moul.io/zapgorm2"
)

// Config defines the database setup options.
type Config struct {
	Dialect  string // mysql, postgres, sqlite
	DSN      string // data source name
	LogLevel string // silent, error, warn, info
}

// New creates a new *gorm.DB instance with the given configuration.
func New(cfg Config) (*gorm.DB, error) {
	dialect, err := getDialect(cfg)
	if err != nil {
		return nil, err
	}

	zapGorm := zapgorm2.New(logger.L())
	zapGorm.SetAsDefault()
	zapGorm.LogLevel = parseLogLevel(cfg.LogLevel)

	db, err := gorm.Open(dialect, &gorm.Config{
		Logger: zapGorm,
	})
	if err != nil {
		return nil, err
	}

	return db, nil
}

// getDialect returns a GORM driver based on dialect.
func getDialect(cfg Config) (gorm.Dialector, error) {
	switch strings.ToLower(cfg.Dialect) {
	case "mysql":
		return mysql.Open(cfg.DSN), nil
	case "postgres":
		return postgres.Open(cfg.DSN), nil
	case "sqlite":
		return sqlite.Open(cfg.DSN), nil
	default:
		return nil, fmt.Errorf("unsupported dialect: %s", cfg.Dialect)
	}
}

// parseLogLevel converts a string into zapgorm2-compatible log level.
func parseLogLevel(level string) gormLogger.LogLevel {
	switch strings.ToLower(level) {
	case "silent":
		return gormLogger.Silent
	case "error":
		return gormLogger.Error
	case "warn":
		return gormLogger.Warn
	default:
		return gormLogger.Info
	}
}
