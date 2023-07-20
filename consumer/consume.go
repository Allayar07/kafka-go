package main

import (
	"context"
	"fmt"
	"kafka-test/pkg"
	"log"
	"time"
)

// get all messages from topic in for loop
func main() {
	kafkaConn, err := pkg.KafkaConn(context.Background(), "tcp", "localhost:9092", "topic_2", 0)
	if err != nil {
		log.Fatalln(err)
	}
	if err = kafkaConn.SetReadDeadline(time.Now().Add(time.Second * 10)); err != nil {
		log.Fatalln(err)
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
