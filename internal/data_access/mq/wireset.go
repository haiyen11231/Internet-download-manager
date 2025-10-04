package mq

import (
	"github.com/google/wire"

	"github.com/haiyen11231/Internet-download-manager/internal/data_access/mq/consumer"
	"github.com/haiyen11231/Internet-download-manager/internal/data_access/mq/producer"
)

var WireSet = wire.NewSet(
	consumer.WireSet,
	producer.WireSet,
)