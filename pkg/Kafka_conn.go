package pkg

import (
	"context"
	"github.com/segmentio/kafka-go"
)

func KafkaConn(ctx context.Context, network, address, topic string, partition int) (*kafka.Conn, error) {
	conn, err := kafka.DialLeader(ctx, network, address, topic, partition)
	if err != nil {
		return nil, err
	}
	return conn, nil
}
