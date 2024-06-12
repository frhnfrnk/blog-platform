package kafka

import (
	"log"

	"github.com/IBM/sarama"
)

type Consumer struct {
	broker string
}

func NewConsumer(broker string) *Consumer {
	return &Consumer{broker: broker}
}

func (c *Consumer) Consume() {
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true

	client, err := sarama.NewClient([]string{c.broker}, config)
	if err != nil {
		log.Fatalf("Error creating Kafka client: %v", err)
	}

	consumer, err := sarama.NewConsumerFromClient(client)
	if err != nil {
		log.Fatalf("Error creating Kafka consumer: %v", err)
	}

	partitions, err := consumer.Partitions("new-comments")
	if err != nil {
		log.Fatalf("Error getting partitions: %v", err)
	}

	for _, partition := range partitions {
		pc, err := consumer.ConsumePartition("new-comments", partition, sarama.OffsetNewest)
		if err != nil {
			log.Fatalf("Error starting partition consumer: %v", err)
		}

		go func(pc sarama.PartitionConsumer) {
			for msg := range pc.Messages() {
				log.Printf("New post: %s", string(msg.Value))
				// Handle the message (e.g., send a notification)
			}
		}(pc)
	}

	select {}
}
