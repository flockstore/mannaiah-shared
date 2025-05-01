// Package logger provides a reusable Zap logger with environment-based presets.
package logger

import (
	"github.com/flockstore/mannaiah-shared/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"sync"
)

var (
	once sync.Once
	log  *zap.Logger
)

// Config defines the logger setup options.
// Env selects between production and development modes.
// Encoding selects output format: "json" or "console".
type Config struct {
	Env           config.EnvType // EnvProduction or EnvDevelopment
	Level         string         // "debug", "info", "warn", "error"
	Encoding      string         // "json" or "console"
	DisableCaller bool           // disables logging of the calling function
	DisableStack  bool           // disables stacktrace logging on errors
}

// Init sets up a global Zap logger with the given config.
func Init(cfg Config) {
	once.Do(func() {
		var zapCfg zap.Config
		switch cfg.Env {
		case config.EnvProduction:
			zapCfg = zap.NewProductionConfig()
		default:
			zapCfg = zap.NewDevelopmentConfig()
		}

		zapCfg.Level = zap.NewAtomicLevelAt(parseLevel(cfg.Level))

		if cfg.Encoding != "" {
			zapCfg.Encoding = cfg.Encoding
		}

		zapCfg.DisableCaller = cfg.DisableCaller
		zapCfg.DisableStacktrace = cfg.DisableStack

		logger, err := zapCfg.Build()
		if err != nil {
			panic("cannot initialize logger: " + err.Error())
		}

		log = logger
	})
}

// L returns the shared global logger instance.
func L() *zap.Logger {
	if log == nil {
		panic("logger not initialized, call logger.Init() first")
	}
	return log
}

// parseLevel converts string to zapcore.Level, defaults to InfoLevel.
func parseLevel(lvl string) zapcore.Level {
	switch lvl {
	case "debug":
		return zapcore.DebugLevel
	case "warn":
		return zapcore.WarnLevel
	case "error":
		return zapcore.ErrorLevel
	default:
		return zapcore.InfoLevel
	}
}
