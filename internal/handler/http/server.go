package http

import (
	"context"
	"net/http"
	"time"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"

	"github.com/haiyen11231/Internet-download-manager/internal/configs"
	"github.com/haiyen11231/Internet-download-manager/internal/generated/grpc/go_load"
	"github.com/haiyen11231/Internet-download-manager/internal/utils"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type server struct {
	grpcConfig configs.GRPC
	httpConfig configs.HTTP
	logger     *zap.Logger
}

type Server interface {
	Start(ctx context.Context) error
}

func NewServer(grpcConfig configs.GRPC, httpConfig configs.HTTP, logger *zap.Logger) Server {
	return &server{
		grpcConfig: grpcConfig,
		httpConfig: httpConfig,
		logger:     logger,
	}
}

// grpc gw thiet lap server http, server nay se bien doi het cac req tu json sang grpc, roi forward den dia chi (0.0.0.0:8080) grpc server
func (s server) Start(ctx context.Context) error {
	logger := utils.LoggerWithContext(ctx, s.logger)

	grpcMux := runtime.NewServeMux()
	if err := go_load.RegisterGoLoadServiceHandlerFromEndpoint(ctx, grpcMux, s.grpcConfig.Address, []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}); err != nil {
		return err
	}

	httpServer := http.Server{
		Addr:              s.httpConfig.Address,
		ReadHeaderTimeout: time.Minute,
		Handler: grpcMux,
	}

	logger.With(zap.String("address", s.httpConfig.Address)).Info("starting http server")
	return httpServer.ListenAndServe() // request http toi cong 8081 se dc bien doi sang grpc va forward toi server grpc
}