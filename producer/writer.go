package producer

import (
	"context"
	"github.com/segmentio/kafka-go"
)

type KafkaWriter struct {
	writer *kafka.Writer
}

func NewKafkaWriter() *KafkaWriter {
	writer := &kafka.Writer{
		Addr:  kafka.TCP("localhost:9092"),
		Topic: "topic_2",
	}
	return &KafkaWriter{writer: writer}
}

func (k *KafkaWriter) WriteMessageToOtherTopic(ctx context.Context, messages chan kafka.Message, committedMessages chan kafka.Message) error {
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case msg := <-messages:
			if err := k.writer.WriteMessages(ctx, kafka.Message{Value: msg.Value}); err != nil {
				return err
			}
			select {
			case <-ctx.Done():
				return ctx.Err()
			case committedMessages <- msg:
			}
		}
	}
}
