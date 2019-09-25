package logger

import "go.uber.org/zap"

var Logger *zap.SugaredLogger

// Инициализация логера
func InitLogger() (*zap.SugaredLogger, error) {
	logger, err := zap.NewProduction()
	if err != nil {
		return nil, err
	}
	Logger = logger.Sugar()
	return Logger, nil
}

func GetLogger() (*zap.SugaredLogger, error) {
	if Logger != nil {
		return Logger, nil
	}
	return InitLogger()
}
