package data_access

import (
	"github.com/google/wire"
	"github.com/haiyen11231/Internet-download-manager/internal/data_access/database"
)

var WireSet = wire.NewSet(
	database.WireSet,
)