package kafka

import (
	"context"
	"fmt"
	"os"

	"github.com/segmentio/kafka-go"
)

func StartKafka() string {
	conf := kafka.ReaderConfig{
		Brokers: []string{"localhost:9092"},
		Topic:   "myTopic",
		GroupID: "1",
	}

	reader := kafka.NewReader(conf)

	for {
		m, err := reader.ReadMessage(context.Background())
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}

		message := string(m.Value)

		fmt.Println("Message is:", message)
		return message
	}
}
