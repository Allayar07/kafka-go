package producer

import (
	"context"
	"github.com/pkg/errors"
	"github.com/segmentio/kafka-go"
	"log"
)

type KafkaReader struct {
	reader *kafka.Reader
}

func NewKafkaReader() *KafkaReader {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{"localhost:9092"},
		Topic:   "topic_1",
		GroupID: "group_1",
	})
	return &KafkaReader{
		reader: reader,
	}
}

func (k *KafkaReader) FetchMessage(ctx context.Context, messages chan kafka.Message) error {
	for {
		msg, err := k.reader.FetchMessage(ctx)
		if err != nil {
			return err
		}
		select {
		// check context is not expired, if it expired then return error
		case <-ctx.Done():
			return ctx.Err()
		// if context is not expired then send message to messages channel
		case messages <- msg:
			log.Println("successfully fetched message: ", string(msg.Value))
		}
	}
}

func (k *KafkaReader) CommitMessage(ctx context.Context, commitMessages <-chan kafka.Message) error {
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case msg := <-commitMessages:
			if err := k.reader.CommitMessages(ctx, msg); err != nil {
				return errors.Wrap(err, "Reader.CommitMessages")
			}
			log.Printf("committed successfully message: %s ", string(msg.Value))
		}
	}
}
