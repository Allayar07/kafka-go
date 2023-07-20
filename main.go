package main

import (
	"context"
	"github.com/segmentio/kafka-go"
	"golang.org/x/sync/errgroup"
	"kafka-test/producer"
	"log"
)

func main() {
	//kafkaReader := producer.NewKafkaReader()
	//kafkaWriter := producer.NewKafkaWriter()
	//
	//ctx := context.Background()
	//messages := make(chan kafka.Message, 1000)
	//committedMessages := make(chan kafka.Message, 1000)
	//
	//g, ctxg := errgroup.WithContext(ctx)
	//
	//g.Go(func() error {
	//	if err := kafkaReader.FetchMessage(ctxg, messages); err != nil {
	//		fmt.Println("error is :", err)
	//		return err
	//	}
	//	return nil
	//})
	//
	//g.Go(func() error {
	//	if err := kafkaWriter.WriteMessageToOtherTopic(ctxg, messages, committedMessages); err != nil {
	//		fmt.Println("error is :", err)
	//		return err
	//	}
	//	return nil
	//})
	//
	//g.Go(func() error {
	//	if err := kafkaReader.CommitMessage(ctxg, committedMessages); err != nil {
	//		fmt.Println("error is :", err)
	//		return err
	//	}
	//	return nil
	//})
	//
	//if WaitErr := g.Wait(); WaitErr != nil {
	//	log.Fatalln(WaitErr)
	//}
	reader := producer.NewKafkaReader()
	writer := producer.NewKafkaWriter()

	ctx := context.Background()
	messages := make(chan kafka.Message, 1000)
	messageCommitChan := make(chan kafka.Message, 1000)

	g, ctx := errgroup.WithContext(ctx)

	g.Go(func() error {
		return reader.FetchMessage(ctx, messages)
	})

	g.Go(func() error {
		return writer.WriteMessageToOtherTopic(ctx, messages, messageCommitChan)
	})

	g.Go(func() error {
		return reader.CommitMessage(ctx, messageCommitChan)
	})

	err := g.Wait()
	if err != nil {
		log.Fatalln(err)
	}
}
