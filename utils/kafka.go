package utils

import "time"
import "github.com/segmentio/kafka-go"

type KafkaReader struct {
	reader *kafka.Reader
}

type KafkaWriter struct {
	writer *kafka.Writer
}

var kafkaReader *KafkaReader //单列模式
var kafkaWriter *KafkaWriter

func GetKafkaReader(topic string, brokers []string) *KafkaReader {
	if kafkaReader == nil {
		kafkaReader := &KafkaReader{}

		kafkaReader.writer = kafka.NewReader(kafka.ReaderConfig{
			Brokers:   brokers,
			Topic:     topic,
			Partition: 0,
			MinBytes:  1,
			MaxBytes:  10e6,
		})
	}
	return kafkaReader
}

func GetKafkaWriter(topic string, brokers []string) *KafkaWriter {
	if kafkaWriter == nil {
		kafkaWriter := &KafkaWriter{}

		kafkaWriter.writer = kafka.NewWriter(kafka.WriterConfig{
			Brokers:      brokers,
			Topic:        topic,
			Balancer:     &kafka.LeastBytes{},
			BatchTimeout: 5 * time.Millisecond,
		})
	}

	return kafkaWriter
}
