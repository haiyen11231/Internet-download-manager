package http

import (
	"context"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"

	"github.com/haiyen11231/Internet-download-manager/internal/generated/grpc/go_load"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type server struct {
}

type Server interface {
	Start(ctx context.Context) error
}

func NewServer() Server {
	return &server{}
}

// grpc gw thiet lap server http, server nay se bien doi het cac req tu json sang grpc, roi forward den dia chi (0.0.0.0:8080) grpc server
func (s *server) Start(ctx context.Context) error {
	mux := runtime.NewServeMux()
	if err := go_load.RegisterGoLoadServiceHandlerFromEndpoint(ctx, mux, "0.0.0.0:8080", []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}); err != nil {
		return err
	}
	return http.ListenAndServe(":8081", mux) // request http toi cong 8081 se dc bien doi sang grpc va forward toi server grpc
}