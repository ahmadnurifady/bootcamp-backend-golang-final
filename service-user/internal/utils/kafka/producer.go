package kafka

import (
	"github.com/IBM/sarama"
	"time"
)

type Producer struct {
	producer sarama.SyncProducer
	//topic    string
}

func NewKafkaProducer(brokers []string, topic string) (*Producer, error) {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Producer.Retry.Max = 5
	config.Producer.Retry.Backoff = 100 * time.Millisecond
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Compression = sarama.CompressionSnappy

	//log.Printf("Creating Kafka producer with brokers: %v and topic: %s", brokers, topic)

	producer, err := sarama.NewSyncProducer(brokers, config)
	if err != nil {
		return nil, err
	}

	return &Producer{
		producer: producer,
		//topic:    topic,
	}, nil
}

func (kp *Producer) SendMessage(key string, value []byte, topic string) error {
	msg := &sarama.ProducerMessage{
		Topic: topic,
		Key:   sarama.StringEncoder(key),
		Value: sarama.ByteEncoder(value),
	}

	_, _, err := kp.producer.SendMessage(msg)
	return err
}

func (kp *Producer) Close() error {
	return kp.producer.Close()
}
