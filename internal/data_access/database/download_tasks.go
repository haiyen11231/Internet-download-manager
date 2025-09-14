package database

import "github.com/doug-martin/goqu/v9"

type DownloadTask struct {
	TaskID         uint64 `sql:"task_id"`
	AccountID      uint64 `sql:"account_id"`
	DownloadType   int    `sql:"download_type"`
	FileURL        string `sql:"file_url"`
	DownloadStatus string `sql:"download_status"`
	Metadata       string `sql:"metadata"`
}

type DownloadTaskDataAccessor interface {
	CreateDownloadTask(task DownloadTask) error
	GetDownloadTaskByID(taskID uint64) (*DownloadTask, error)
	UpdateDownloadTask(task DownloadTask) error
	DeleteDownloadTask(taskID uint64) error
	WithDatabase(database Database) DownloadTaskDataAccessor
}

type downloadTaskDataAccessor struct {
	database Database
}

func NewDownloadTaskDataAccessor(database *goqu.Database) DownloadTaskDataAccessor {
	return &downloadTaskDataAccessor{
		database: database,
	}
}