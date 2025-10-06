package handler

import (
	"github.com/google/wire"

	"github.com/haiyen11231/Internet-download-manager/internal/handler/consumers"
	"github.com/haiyen11231/Internet-download-manager/internal/handler/grpc"
	"github.com/haiyen11231/Internet-download-manager/internal/handler/http"
	"github.com/haiyen11231/Internet-download-manager/internal/handler/jobs"
)

var WireSet = wire.NewSet(
	grpc.WireSet,
	http.WireSet,
	consumers.WireSet,
	jobs.WireSet,
)