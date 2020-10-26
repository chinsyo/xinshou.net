package core

import (
	"go.uber.org/zap"
)

var SharedLogger *zap.SugaredLogger

func init() {
	logger, _ := zap.NewProduction()
	SharedLogger = logger.Sugar()
	defer logger.Sync()
}
