package kafka

import (
	"context"
	"time"

	"github.com/segmentio/kafka-go"

	"github.com/sirupsen/logrus"
)

type Kafka struct {
	w   *kafka.Writer
	log *logrus.Logger
}

func Connect(address, topic string) *Kafka {
	log := logrus.New()

	w := &kafka.Writer{
		Addr:   kafka.TCP(address),
		Topic:  topic,
		Logger: log,
	}

	return &Kafka{
		w:   w,
		log: log,
	}
}

func (k *Kafka) Push(parent context.Context, key, value []byte) (err error) {
	message := kafka.Message{
		Key:   key,
		Value: value,
		Time:  time.Now(),
	}

	return k.w.WriteMessages(parent, message)
}
