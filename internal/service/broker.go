package service

import (
	"context"
	"fmt"

	"github.com/Milad75Rasouli/IEC61850CRUD/database"
	"github.com/Milad75Rasouli/IEC61850CRUD/model"
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type ConsumerConnection struct {
	IP       string
	Port     int
	Topic    string
	GroupID  string
	PollTime int
}

type Consumer struct {
	topic     string
	consumer  *kafka.Consumer
	pollTimer int
	db        database.Storage
	Stop      chan bool
}

func NewConsumer(conn ConsumerConnection, db database.Storage) (*Consumer, error) {

	url := fmt.Sprintf("%s:%d", conn.IP, conn.Port)
	consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": url,
		"group.id":          conn.GroupID,
		"auto.offset.reset": "smallest"})

	if err != nil {
		return nil, err
	}
	return &Consumer{
		topic:     conn.Topic,
		consumer:  consumer,
		db:        db,
		pollTimer: conn.PollTime,
	}, nil
}

func (c *Consumer) SubscribeTopic() error {
	err := c.consumer.Subscribe(c.topic, nil)
	if err != nil {
		return err
	}
	return nil
}

func (c *Consumer) InsertToDB() error {
	fmt.Println("Consumer is started.")
	for {
		ev := c.consumer.Poll(c.pollTimer)
		//fmt.Printf("%+v", ev)
		switch ev.(type) {
		case *kafka.Message:
			msg := ev.String()
			//fmt.Println("Consumer Received:", msg)
			c.db.AddSignal("IEC61850", model.Signal{Key: "IEC61850-Signal", Value: msg[5:]}, context.Background())
		case kafka.Error:
			return fmt.Errorf(ev.String())
		}
	}

}
