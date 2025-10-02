package producer

import (
	"context"
	"encoding/json"

	"github.com/haiyen11231/Internet-download-manager/internal/utils"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	MessageQueueDownloadTaskCreated = "download_task_created"
)

type DownloadTaskCreated struct {
	TaskID uint64 `json:"task_id"`
}

type DownloadTaskCreatedProducer interface {
	Produce(ctx context.Context, event DownloadTaskCreated) error
}

type downloadTaskCreatedProducer struct {
	client Client
	logger *zap.Logger
}

func NewDownloadTaskCreatedProducer(
	client Client,
	logger *zap.Logger,
) DownloadTaskCreatedProducer {
	return &downloadTaskCreatedProducer{
		client: client,
		logger: logger,
	}
}

func (d downloadTaskCreatedProducer) Produce(ctx context.Context, event DownloadTaskCreated) error {
	logger := utils.LoggerWithContext(ctx, d.logger)

	eventBytes, err := json.Marshal(event)
	if err != nil {
		logger.With(zap.Error(err)).Error("failed to marshal download task created event")
		return status.Error(codes.Internal, "failed to marshal download task created event")
	}

	err = d.client.Produce(ctx, MessageQueueDownloadTaskCreated, eventBytes)
	if err != nil {
		logger.With(zap.Error(err)).Error("failed to produce download task created event")
		return status.Error(codes.Internal, "failed to produce download task created event")
	}

	return nil
}