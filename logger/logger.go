package logger

import "go.uber.org/zap"

type Logger struct {
	logger *zap.SugaredLogger
}

func GetLogger() *zap.SugaredLogger {
	return zap.NewNop().Sugar()
}
