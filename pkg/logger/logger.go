package logger

import (
	"fmt"

	"go.uber.org/zap"
)

// New creates a sugared zap logger. If level is empty, DefaultLogLevel is used.
func New(level string) *zap.SugaredLogger {
	if level == "" {
		level = DefaultLogLevel
	}

	var cfg zap.Config
	if level == "debug" {
		cfg = zap.NewDevelopmentConfig()
	} else {
		cfg = zap.NewProductionConfig()
	}

	cfg.Level = zap.NewAtomicLevelAt(parseLevel(level))

	logger, err := cfg.Build()
	if err != nil {
		panic(fmt.Sprintf("failed to build logger: %v", err))
	}
	return logger.Sugar()
}
