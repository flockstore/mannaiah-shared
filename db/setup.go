package db

import (
	"github.com/flockstore/mannaiah-shared/config"
	"gorm.io/gorm"
)

// SetupEnvironmentDatabase creates a database from the provided environment
func SetupEnvironmentDatabase() (*gorm.DB, error) {

	dsn := config.MustGet(config.DatabaseDSN)
	d := config.MustGet(config.DatabaseDialect)
	lvl := config.MustGet(config.LogLevel)

	cfg := Config{
		Dialect:  d,
		LogLevel: lvl,
		DSN:      dsn,
	}

	return New(cfg)

}
