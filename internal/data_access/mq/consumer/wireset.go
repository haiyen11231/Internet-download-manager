package consumer

import "github.com/google/wire"

var Wire = wire.NewSet(
	NewConsumer,
)