package kafka

import (
	"github.com/IBM/sarama"
)

type Producer struct {
	broker string
}

func NewProducer(broker string) *Producer {
	return &Producer{broker: broker}
}

func (p *Producer) SendMessage(topic, message string) error {
	config := sarama.NewConfig()
	producer, err := sarama.NewSyncProducer([]string{p.broker}, config)
	if err != nil {
		return err
	}
	defer producer.Close()

	msg := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(message),
	}

	_, _, err = producer.SendMessage(msg)
	return err
}
