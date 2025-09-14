package handler

import (
	"github.com/google/wire"
	"github.com/haiyen11231/Internet-download-manager/internal/handler/grpc"
	"github.com/haiyen11231/Internet-download-manager/internal/handler/http"
)

var WireSet = wire.NewSet(
	grpc.WireSet,
	http.WireSet,
)