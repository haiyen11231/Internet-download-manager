package data_access

import (
	"github.com/google/wire"

	"github.com/haiyen11231/Internet-download-manager/internal/data_access/cache"
	"github.com/haiyen11231/Internet-download-manager/internal/data_access/database"
	"github.com/haiyen11231/Internet-download-manager/internal/data_access/mq"
)

var WireSet = wire.NewSet(
	cache.WireSet,
	database.WireSet,
	mq.WireSet,
)