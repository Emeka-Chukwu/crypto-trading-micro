package main

import (
	"fmt"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("llllll")
}

type KafkaPlacer struct {
	producer   *kafka.Producer
	topic      string
	deliverych chan kafka.Event
}

func NewKafkaPlacer(p *kafka.Producer, topic string) *KafkaPlacer {
	server := gin.Default()
	return &KafkaPlacer{
		producer:   p,
		topic:      topic,
		deliverych: make(chan kafka.Event, 10000),
	}
}
