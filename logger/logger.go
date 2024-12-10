package logger

import "go.uber.org/zap"

type Logger struct {
	logger *zap.SugaredLogger
}

func NewLogger() *zap.SugaredLogger {
	return zap.NewNop().Sugar()
}
