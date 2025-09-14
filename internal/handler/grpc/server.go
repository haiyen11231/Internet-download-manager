package grpc

import (
	"context"
	"net"

	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/validator"
	"github.com/haiyen11231/Internet-download-manager/internal/generated/grpc/go_load"
	"google.golang.org/grpc"
)

type Server interface {
	Start(ctx context.Context) error
}
type server struct {
	handler go_load.GoLoadServiceServer
}

func NewServer (handler go_load.GoLoadServiceServer) Server {
	return &server{
		handler: handler, // cai dat logic sinh ra boi grpc
	}
}

// listen to request from client toi port 8080 vaf xu ly request do
func (s *server) Start(ctx context.Context) error {
	listener, err := net.Listen("tcp", "0.0.0.0:8080") // 0.0.0.0 allow you to ket noi tu ben ngoai (locahost only allow access from local machine)
	if err != nil {
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

	return grpcServer.Serve(listener)
}