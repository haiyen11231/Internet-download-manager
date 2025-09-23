//go:build wireinject
// +build wireinject

//
//go:generate go run github.com/google/wire/cmd/wire
package wiring

import (
	"github.com/google/wire"
	"github.com/haiyen11231/Internet-download-manager/internal/configs"
	"github.com/haiyen11231/Internet-download-manager/internal/data_access"
	"github.com/haiyen11231/Internet-download-manager/internal/handler"
	"github.com/haiyen11231/Internet-download-manager/internal/handler/grpc"
	"github.com/haiyen11231/Internet-download-manager/internal/logic"
	"github.com/haiyen11231/Internet-download-manager/internal/utils"
)

var WireSet = wire.NewSet(
	configs.WireSet,
	utils.WireSet,
	data_access.WireSet,
	logic.WireSet,
	handler.WireSet,
)

func InitializeGRPCServer(configFilePath configs.ConfigFilePath) (grpc.Server, func(), error) {
	wire.Build(WireSet)

	return nil, nil, nil
}