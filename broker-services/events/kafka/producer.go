package event_kafka

// "github.com/confluentinc/confluent-kafka-go/kafka"
// // "github.com/confluentinc/confluent-kafka-go/kafka"
// // "github.com/confluentinc/confluent-kafka-go/v2/kafka"

// type KafkaPlacer struct {
// 	producer   *kafka.Producer
// 	topic      string
// 	deliverych chan kafka.Event
// }

// func NewKafkaPlacer(p *kafka.Producer, topic string) *KafkaPlacer {
// 	return &KafkaPlacer{
// 		producer:   p,
// 		topic:      topic,
// 		deliverych: make(chan kafka.Event, 10000),
// 	}
// }

// func (kp *KafkaPlacer) placeEvent(eventType string, size int) error {
// 	var (
// 		format  = fmt.Sprintf("%s - %d", eventType, size)
// 		payload = []byte(format)
// 	)
// 	err := kp.producer.Produce(&kafka.Message{
// 		TopicPartition: kafka.TopicPartition{Topic: &kp.topic, Partition: kafka.PartitionAny},
// 		Value:          payload,
// 	},
// 		kp.deliverych,
// 	)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	<-kp.deliverych
// 	fmt.Printf("placed order on the queue %s\n", format)
// 	return nil
// }
