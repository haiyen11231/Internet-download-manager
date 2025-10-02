package mq

var WireSet = wire.NewSet(
	consumer.WireSet,
	producer.WireSet,
)