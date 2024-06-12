package main

import (
	"log"
	"os"

	"github.com/frhnfrnk/blog-platform-microservices/notification-service/internal/kafka"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	broker := os.Getenv("KAFKA_BROKER")
	kafkaConsumer := kafka.NewConsumer(broker)
	kafkaConsumer.Consume()
}
