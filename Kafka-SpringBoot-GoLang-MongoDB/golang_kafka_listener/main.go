package main

import (
	"fmt"
	"golang_kafka_listener/kafka"
	"golang_kafka_listener/mongodb"
	"time"
)

func main() {
	fmt.Println("Starting...")

	message := kafka.StartKafka()

	collection := mongodb.GetClientInstance()

	mongodb.InsertMessage(collection, message)

	fmt.Println("Kafka has been started...")

	time.Sleep(time.Minute * 10)
}
