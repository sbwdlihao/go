package kafka

import (
	"testing"
	"fmt"
	"github.com/Shopify/sarama"
)

func TestKafka0 (t *testing.T) {
	config := sarama.NewConfig()
	producer, err := sarama.NewAsyncProducer([]string{"localhost:9092"}, config)
	if err == nil {
		fmt.Println("producer create success")
		fmt.Println(producer)
	} else {
		fmt.Println(err)
	}
}

