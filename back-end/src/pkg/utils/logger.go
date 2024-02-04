package utils

import "go.uber.org/zap"

var Logger *zap.Logger

func InitLogging() {
	Logger, _ = zap.NewProduction()
	Logger.Named("unnamed-logger")
}

func InitLoggingWithName(name string) {
	Logger, _ = zap.NewProduction()
	Logger.Named(name)
}
