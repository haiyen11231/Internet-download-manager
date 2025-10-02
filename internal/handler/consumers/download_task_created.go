package consumers

import (
	"context"

	"github.com/haiyen11231/Internet-download-manager/internal/data_access/mq/producer"
	"github.com/haiyen11231/Internet-download-manager/internal/utils"
)

type DownloadTaskCreated interface {
	Handle(ctx context.Context, event producer.DownloadTaskCreated) error
}

type downloadTaskCreated struct {
	logger *zap.Logger
}

func NewDownloadTaskCreated(
	logger *zap.Logger,
) DownloadTaskCreated {
	return &downloadTaskCreated{
		logger: logger,
	}
}

func (d downloadTaskCreated) Handle(ctx context.Context, event producer.DownloadTaskCreated) error {
	logger := utils.LoggerWithContext(ctx, d.logger).With(zap.Any("event", event))
	logger.Info("download task created event received")

	return nil
}