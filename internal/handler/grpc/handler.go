package grpc

import (
	"context"

	"github.com/haiyen11231/Internet-download-manager/internal/generated/grpc/go_load"
	"github.com/haiyen11231/Internet-download-manager/internal/logic"
)

// to forward input to logic layer + dkien condition khong the check bang validation
type Handler struct {
	go_load.UnimplementedGoLoadServiceServer
	accountLogic logic.Account
	downloadTaskLogic logic.DownloadTask
}

func NewHandler(accountLogic logic.Account, downloadTaskLogic logic.DownloadTask) go_load.GoLoadServiceServer {
	return &Handler{
		accountLogic:    accountLogic,
		downloadTaskLogic: downloadTaskLogic,
	}
}

// CreateAccount implements go_load.GoLoadServiceServer
func (h Handler) CreateAccount(ctx context.Context, req *go_load.CreateAccountRequest) (*go_load.CreateAccountResponse, error) {
	// View layer: co input tu req -> forward to logic layer to handle logic
	output, err := h.accountLogic.CreateAccount(ctx, logic.CreateAccountParams{
		AccountName: req.GetAccountName(),
		Password:    req.GetPassword(),
	})
	if err != nil {
		return nil, err
	}
	return &go_load.CreateAccountResponse{
		AccountId:          output.ID,
	}, nil
}

// CreateSession handles user login and returns a token
func (h Handler) CreateSession(ctx context.Context, req *go_load.CreateSessionRequest) (*go_load.CreateSessionResponse, error) {
	// Receive username + password
	// Check
	// - if username exist 
	// - password valid
	// return user info + session token

	output, err := h.accountLogic.CreateSession(ctx, logic.CreateSessionParams{
		AccountName: req.GetAccountName(),
		Password:    req.GetPassword(),
	})
	if err != nil {
		return nil, err
	}

	return &go_load.CreateSessionResponse{
		Account: output.Account,
		Token:   output.Token,
	}, nil
}

// CreateDownloadTask handles creating a download task
func (h Handler) CreateDownloadTask(ctx context.Context, req *go_load.CreateDownloadTaskRequest) (*go_load.CreateDownloadTaskResponse, error) {
	output, err := h.downloadTaskLogic.CreateDownloadTask(ctx, logic.CreateDownloadTaskParams{
		Token:        req.GetToken(),
		DownloadType: req.GetDownloadType(),
		URL:          req.GetUrl(),
	})
	if err != nil {
		return nil, err
	}

	return &go_load.CreateDownloadTaskResponse{
		DownloadTask: output.DownloadTask,
	}, nil
}

// GetDownloadTaskList returns a list of download tasks
func (h Handler) GetDownloadTaskList(ctx context.Context, req *go_load.GetDownloadTaskListRequest) (*go_load.GetDownloadTaskListResponse, error) {
	output, err := h.downloadTaskLogic.GetDownloadTaskList(ctx, logic.GetDownloadTaskListParams{
		Token:  req.GetToken(),
		Offset: req.GetOffset(),
		Limit:  req.GetLimit(),
	})
	if err != nil {
		return nil, err
	}
	
	return &go_load.GetDownloadTaskListResponse{
		DownloadTaskList:       output.DownloadTaskList,
		TotalDownloadTaskCount: output.TotalDownloadTaskCount,
	}, nil
}

// UpdateDownloadTask updates a download task
func (h Handler) UpdateDownloadTask(ctx context.Context, req *go_load.UpdateDownloadTaskRequest) (*go_load.UpdateDownloadTaskResponse, error) {
	output, err := h.downloadTaskLogic.UpdateDownloadTask(ctx, logic.UpdateDownloadTaskParams{
		Token:          req.GetToken(),
		DownloadTaskID: req.GetDownloadTaskId(),
		URL:            req.GetUrl(),
	})
	if err != nil {
		return nil, err
	}
	
	return &go_load.UpdateDownloadTaskResponse{
		DownloadTask: output.DownloadTask,
	}, nil
}

// DeleteDownloadTask deletes a download task
func (h Handler) DeleteDownloadTask(ctx context.Context, req *go_load.DeleteDownloadTaskRequest) (*go_load.DeleteDownloadTaskResponse, error) {
	if err := h.downloadTaskLogic.DeleteDownloadTask(ctx, logic.DeleteDownloadTaskParams{
		Token:          req.GetToken(),
		DownloadTaskID: req.GetDownloadTaskId(),
	}); err != nil {
		return nil, err
	}

	return &go_load.DeleteDownloadTaskResponse{}, nil
}

// GetDownloadTaskFile streams file content
func (h Handler) GetDownloadTaskFile(req *go_load.GetDownloadTaskFileRequest, stream go_load.GoLoadService_GetDownloadTaskFileServer) error {
	return nil
}