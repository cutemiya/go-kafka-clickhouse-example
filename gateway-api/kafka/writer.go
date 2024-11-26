package kafkaSettings

import (
	"github.com/segmentio/kafka-go"
)

func NewWriter(addr, topic string) *kafka.Writer {
	return &kafka.Writer{
		Addr:                   kafka.TCP(addr),
		Topic:                  topic,
		AllowAutoTopicCreation: true,
	}
}
