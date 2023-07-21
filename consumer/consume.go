package main

import (
	"context"
	"fmt"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/rs/zerolog/pkgerrors"
	"kafka-test/pkg"

	"time"
)

// get all messages from topic in for loop
func main() {
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack
	zerolog.SetGlobalLevel(zerolog.DebugLevel)
	log.Logger = log.With().Caller().Logger()
	kafkaConn, err := pkg.KafkaConn(context.Background(), "tcp", "localhost:9092", "topic_2", 0)
	if err != nil {
		log.Fatal().Err(err).Msg("")
	}
	// SetReadDeadline is need because if something went wrong in send error message for me:)
	if err = kafkaConn.SetReadDeadline(time.Now().Add(time.Second * 10)); err != nil {
		log.Error().Stack().Err(err).Msg("")
	}

	batch := kafkaConn.ReadBatch(1e3, 1e9) // 1e3 = 1000, 1 kb
	bytes := make([]byte, 1e5)

	for {
		_, err = batch.Read(bytes)
		if err != nil {
			break
		}

		fmt.Println(string(bytes))
	}
}
