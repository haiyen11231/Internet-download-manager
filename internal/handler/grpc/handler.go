package grpc

import (
	"context"

	"github.com/haiyen11231/Internet-download-manager/internal/generated/grpc/go_load"
)

type Handler struct {
	go_load.UnimplementedGoLoadServiceServer
}

func NewHandler() *Handler {
	return &Handler{}
}

// CreateAccount implements go_load.GoLoadServiceServer
func (h *Handler) CreateAccount(ctx context.Context, req *go_load.CreateAccountRequest) (*go_load.CreateAccountResponse, error) {
	panic("unimplemented")
}

// CreateSession handles user login and returns a token
func (h *Handler) CreateSession(ctx context.Context, req *go_load.CreateSessionRequest) (*go_load.CreateSessionResponse, error) {
	panic("unimplemented")
}

// CreateDownloadTask handles creating a download task
func (h *Handler) CreateDownloadTask(ctx context.Context, req *go_load.CreateDownloadTaskRequest) (*go_load.CreateDownloadTaskResponse, error) {
	panic("unimplemented")
}

// GetDownloadTaskList returns a list of download tasks
func (h *Handler) GetDownloadTaskList(ctx context.Context, req *go_load.GetDownloadTaskListRequest) (*go_load.GetDownloadTaskListResponse, error) {
	panic("unimplemented")
}

// UpdateDownloadTask updates a download task
func (h *Handler) UpdateDownloadTask(ctx context.Context, req *go_load.UpdateDownloadTaskRequest) (*go_load.UpdateDownloadTaskResponse, error) {
	return nil, nil
}

// DeleteDownloadTask deletes a download task
func (h *Handler) DeleteDownloadTask(ctx context.Context, req *go_load.DeleteDownloadTaskRequest) (*go_load.DeleteDownloadTaskResponse, error) {
	return nil, nil
}

// GetDownloadTaskFile streams file content
func (h *Handler) GetDownloadTaskFile(req *go_load.GetDownloadTaskFileRequest, stream go_load.GoLoadService_GetDownloadTaskFileServer) error {
	return nil
}