package logger

import (
	"go.uber.org/zap"
)

var _logger *zap.SugaredLogger

func init() {
	loggerConfig := zap.NewProductionConfig()
	loggerConfig.Level = zap.NewAtomicLevelAt(zap.InfoLevel)
	logger, _ := loggerConfig.Build()
	_logger = logger.Sugar()

}
func EnableDevelopmentLogger() {
	loggerConfig := zap.NewDevelopmentConfig()
	loggerConfig.Level = zap.NewAtomicLevelAt(zap.InfoLevel)
	logger, _ := loggerConfig.Build()
	_logger = logger.Sugar()
}

func Get() *zap.SugaredLogger {
	return _logger
}
