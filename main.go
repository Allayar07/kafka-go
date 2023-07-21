package main

import (
	"context"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/rs/zerolog/pkgerrors"
	"github.com/segmentio/kafka-go"
	"golang.org/x/sync/errgroup"
	"kafka-test/producer"
	//"log"
)

func main() {
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack
	zerolog.SetGlobalLevel(zerolog.DebugLevel)
	log.Logger = log.With().Caller().Logger()
	reader := producer.NewKafkaReader()
	writer := producer.NewKafkaWriter()

	ctx := context.Background()
	messages := make(chan kafka.Message, 1000)
	messageCommitChan := make(chan kafka.Message, 1000)

	g, ctx := errgroup.WithContext(ctx)

	g.Go(func() error {
		if err := reader.FetchMessage(ctx, messages); err != nil {
			log.Error().Stack().Err(err).Msg("")
			return err
		}
		return nil
	})

	g.Go(func() error {
		if err := writer.WriteMessageToOtherTopic(ctx, messages, messageCommitChan); err != nil {
			log.Error().Stack().Err(err).Msg("")
			return err
		}
		return nil
	})

	g.Go(func() error {
		if err := reader.CommitMessage(ctx, messageCommitChan); err != nil {
			log.Error().Stack().Err(err).Msg("")
			return err
		}
		return nil
	})

	err := g.Wait()
	if err != nil {
		log.Fatal().Err(err).Msg("")
	}
}
