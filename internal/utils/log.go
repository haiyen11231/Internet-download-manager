package utils

import (
	"context"

	"github.com/haiyen11231/Internet-download-manager/internal/configs"
	"go.uber.org/zap"
)

func getZapLoggerLevel(level string) zap.AtomicLevel {
	switch level {
	case "debug":
		return zap.NewAtomicLevelAt(zap.DebugLevel)	
	case "info":
		return zap.NewAtomicLevelAt(zap.InfoLevel)
	case "warn":
		return zap.NewAtomicLevelAt(zap.WarnLevel)	
	case "error":
		return zap.NewAtomicLevelAt(zap.ErrorLevel)
	case "panic":
		return zap.NewAtomicLevelAt(zap.PanicLevel)
	default:
		return zap.NewAtomicLevelAt(zap.InfoLevel)
	}
}

func InitializeLogger(logConfig configs.Log) (*zap.Logger, func(), error) {
	zapLoggerConfig := zap.NewProductionConfig()
	zapLoggerConfig.Level = getZapLoggerLevel(logConfig.Level)

	logger, er := zapLoggerConfig.Build()
	if er != nil {
		return nil, nil, er
	}

	cleanup := func() {
		// deliberately ignore the returned error here
		_ = logger.Sync()
	}

	return logger, cleanup, nil
}

func LoggerWithContext(ctx context.Context, logger *zap.Logger) *zap.Logger {
	return logger
	// return logger.With(zap.String("request_id", ctx.Value("request_id").(string)))
}