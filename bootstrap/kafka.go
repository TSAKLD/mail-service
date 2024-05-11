package bootstrap

import (
	"github.com/segmentio/kafka-go"
)

func KafkaConnect(addr string, topic string, groupID string) (*kafka.Reader, error) {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  []string{addr},
		Topic:    topic,
		MaxBytes: 10e6, // 10MB,
		GroupID:  groupID,
	})

	return r, nil
}
