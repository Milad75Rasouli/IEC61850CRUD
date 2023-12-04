// https://docs.confluent.io/kafka-clients/go/current/overview.html
package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

// type OrderdProducer
func main() {
	topic := "buf"

	p, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": "127.0.0.1:9092",
		"client.id":         "FOO",
		"acks":              "all"})

	if err != nil {
		fmt.Printf("Failed to create producer: %s\n", err)
		os.Exit(1)
	}

	fmt.Printf("%+v\n", p)
	deliverch := make(chan kafka.Event, 10000)
	msg := kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          []byte("FOO")}
	var index int
	for {
		index++
		signal := fmt.Sprintf("signal[%d]=%d", index*20, index*1000)
		msg.Value = []byte(signal)
		err = p.Produce(&msg,
			deliverch,
		)
		if err != nil {
			log.Fatal(err)
		}
		a := <-deliverch
		fmt.Printf("Produced a message:%+v\n", a)
		time.Sleep(time.Second * 3)
	}
}
