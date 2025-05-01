// Package logger provides tests for the shared logger configuration.
package logger

import (
	"bytes"
	"sync"
	"testing"

	"github.com/flockstore/mannaiah-shared/config"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// newBufferLogger returns a zap logger that writes to a buffer for testing.
func newBufferLogger(cfg Config, buf *bytes.Buffer) *zap.Logger {
	encoderCfg := zap.NewProductionEncoderConfig()
	encoderCfg.TimeKey = ""
	if cfg.Encoding == "json" {
		encoder := zapcore.NewJSONEncoder(encoderCfg)
		core := zapcore.NewCore(encoder, zapcore.AddSync(buf), parseLevel(cfg.Level))
		return zap.New(core)
	}
	encoder := zapcore.NewConsoleEncoder(encoderCfg)
	core := zapcore.NewCore(encoder, zapcore.AddSync(buf), parseLevel(cfg.Level))
	return zap.New(core)
}

func resetLogger() {
	log = nil
	once = sync.Once{}
}

// TestLoggerWrites verifies that the logger writes expected output to the buffer.
func TestLoggerWrites(t *testing.T) {
	var buf bytes.Buffer
	cfg := Config{Level: "debug", Encoding: "console"}
	log := newBufferLogger(cfg, &buf)

	log.Info("logger initialized", zap.String("env", "test"))

	output := buf.String()
	assert.Contains(t, output, "logger initialized")
	assert.Contains(t, output, "env")
	assert.Contains(t, output, "test")
}

// TestParseLevel checks correct level parsing.
func TestParseLevel(t *testing.T) {
	assert.Equal(t, zapcore.DebugLevel, parseLevel("debug"))
	assert.Equal(t, zapcore.InfoLevel, parseLevel("info"))
	assert.Equal(t, zapcore.WarnLevel, parseLevel("warn"))
	assert.Equal(t, zapcore.ErrorLevel, parseLevel("error"))
	assert.Equal(t, zapcore.InfoLevel, parseLevel("unknown"))
}

// TestInitLoggerPanicOnBuildError ensures Init panics if logger build fails.
func TestInitLoggerPanicOnBuildError(t *testing.T) {
	resetLogger()
	assert.Panics(t, func() {
		Init(Config{
			Env:      config.EnvDevelopment,
			Level:    "debug",
			Encoding: "invalid-encoding",
		})
	})
}

// TestLoggerAccessBeforeInit ensures L() panics if called before Init.
func TestLoggerAccessBeforeInit(t *testing.T) {
	resetLogger()
	assert.Panics(t, func() {
		_ = L()
	})
}

// TestInitProductionEnv ensures the production config path is exercised.
func TestInitProductionEnv(t *testing.T) {
	resetLogger()

	Init(Config{
		Env:           config.EnvProduction,
		Level:         "info",
		Encoding:      "json",
		DisableCaller: true,
		DisableStack:  true,
	})

	log := L()
	assert.NotNil(t, log)
	log.Info("production mode initialized")
}

// TestInitAndUseLogger validates logger initialization and global usage.
func TestInitAndUseLogger(t *testing.T) {
	resetLogger()

	Init(Config{
		Env:           config.EnvDevelopment,
		Level:         "debug",
		Encoding:      "console",
		DisableCaller: true,
		DisableStack:  true,
	})

	log := L()
	assert.NotNil(t, log)
	log.Debug("debug message")
	log.Info("info message")
	log.Warn("warn message")
	log.Error("error message")
}

// TestLoggerDoubleInit ensures Init only initializes once.
func TestLoggerDoubleInit(t *testing.T) {
	resetLogger()

	Init(Config{Env: config.EnvDevelopment})
	Init(Config{Env: config.EnvProduction})
	log := L()
	assert.NotNil(t, log)
}
