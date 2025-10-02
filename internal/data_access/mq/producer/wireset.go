package producer

import "github.com/google/wire"

var Wire = wire.NewSet(
	NewClient,
	NewDownloadTaskCreatedProducer,
)