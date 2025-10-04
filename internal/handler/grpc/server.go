package grpc

import (
	"context"
	"net"

	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/validator"
	"go.uber.org/zap"
	"google.golang.org/grpc"

	"github.com/haiyen11231/Internet-download-manager/internal/configs"
	"github.com/haiyen11231/Internet-download-manager/internal/generated/grpc/go_load"
	"github.com/haiyen11231/Internet-download-manager/internal/utils"
)

type Server interface {
	Start(ctx context.Context) error
}
type server struct {
	handler go_load.GoLoadServiceServer
	grpcConfig configs.GRPC
	logger     *zap.Logger
}

func NewServer (handler go_load.GoLoadServiceServer, grpcConfig configs.GRPC, logger *zap.Logger) Server {
	return &server{
		handler:   handler,
		grpcConfig: grpcConfig,
		logger:    logger,
	}
}

// listen to request from client toi port 8080 vaf xu ly request do
func (s server) Start(ctx context.Context) error {
	logger := utils.LoggerWithContext(ctx, s.logger)
	listener, err := net.Listen("tcp", s.grpcConfig.Address) // 0.0.0.0 allow you to ket noi tu ben ngoai (locahost only allow access from local machine)
	if err != nil {
		logger.With(zap.Error(err)).Error("failed to open tcp listener")
		return err
	}
	defer listener.Close()

	grpcServer := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			validator.UnaryServerInterceptor(),
		),
		grpc.ChainStreamInterceptor(
			validator.StreamServerInterceptor(),
		),
	)
	go_load.RegisterGoLoadServiceServer(grpcServer, s.handler)
	logger.With(zap.String("address", s.grpcConfig.Address)).Info("starting grpc server")

	return grpcServer.Serve(listener)
}