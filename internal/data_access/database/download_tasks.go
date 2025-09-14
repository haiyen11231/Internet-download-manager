package database

type DownloadTask struct {
	TaskID      uint64    `sql:"task_id"`
	UserID      uint64    `sql:"of_user_id"`
	DownloadType int       `sql:"download_type"`
	FileURL     string    `sql:"file_url"`
	DownloadStatus      string    `sql:"download_status"`
	Metadata    string    `sql:"metadata"`
}