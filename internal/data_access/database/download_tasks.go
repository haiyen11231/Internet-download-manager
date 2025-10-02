package database

import (
	"context"

	"github.com/doug-martin/goqu/v9"
	"github.com/haiyen11231/Internet-download-manager/internal/generated/grpc/go_load"
	"go.uber.org/zap"
)

type DownloadTask struct {
	TaskID         uint64 `db:"task_id"`
	AccountID      uint64 `db:"account_id"`
	DownloadType   go_load.DownloadType `db:"download_type"`
	FileURL        string                `db:"file_url"`
	DownloadStatus go_load.DownloadStatus `db:"download_status"`
	Metadata       string                `db:"metadata"`
}

type DownloadTaskDataAccessor interface {
	CreateDownloadTask(ctx context.Context, task DownloadTask) (uint64, error)
	GetDownloadTaskListOfUser(ctx context.Context, userID, offset, limit uint64) ([]DownloadTask, error)
	GetDownloadTaskCountOfUser(ctx context.Context, userID uint64) (uint64, error)
	UpdateDownloadTask(ctx context.Context, task DownloadTask) error
	DeleteDownloadTask(ctx context.Context, taskID uint64) error
	WithDatabase(database Database) DownloadTaskDataAccessor
}

type downloadTaskDataAccessor struct {
	database Database
	logger   *zap.Logger
}

func NewDownloadTaskDataAccessor(database *goqu.Database, logger *zap.Logger) DownloadTaskDataAccessor {
	return &downloadTaskDataAccessor{
		database: database,
		logger:   logger,
	}
}

func (d downloadTaskDataAccessor) CreateDownloadTask(ctx context.Context, task DownloadTask) (uint64, error) {
	return 1, nil
}

func (d downloadTaskDataAccessor) GetDownloadTaskListOfUser(ctx context.Context, userID, offset, limit uint64) ([]DownloadTask, error) {
	panic("unimplemented")
}

func (d downloadTaskDataAccessor) GetDownloadTaskCountOfUser(ctx context.Context, userID uint64) (uint64, error) {
	panic("unimplemented")
}	

func (d downloadTaskDataAccessor) UpdateDownloadTask(ctx context.Context, task DownloadTask) error {
	panic("unimplemented")
}

func (d downloadTaskDataAccessor) DeleteDownloadTask(ctx context.Context, taskID uint64) error {
	panic("unimplemented")
}

func (d downloadTaskDataAccessor) WithDatabase(database Database) DownloadTaskDataAccessor {
	return &downloadTaskDataAccessor{
		database: database,
		logger:   d.logger,
	}	
}
