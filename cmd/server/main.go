package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/Milad75Rasouli/IEC61850CRUD/database"
	"github.com/Milad75Rasouli/IEC61850CRUD/internal/api"
	"github.com/Milad75Rasouli/IEC61850CRUD/internal/service"
)

func main() {

	// Database Mongo
	dbconfig := database.MongoConnection{
		Port: 27017,
		IP:   "127.0.0.1",
		User: "root",
		Pass: "1234qwer",
	}

	ctx, cancel := context.WithTimeout(context.Background(), 1500*time.Millisecond)
	defer cancel()
	db, err := database.NewMongo(dbconfig, ctx)
	if err != nil {
		log.Fatalln(err)
	}

	// Consumer Kafka
	consumerConfig := service.ConsumerConnection{
		Port:     9092,
		IP:       "127.0.0.1",
		GroupID:  "FOO",
		PollTime: 100,
		Topic:    "buf",
	}
	consumer, err := service.NewConsumer(consumerConfig, db)
	if err != nil {
		log.Fatalln(err)
	}
	err = consumer.SubscribeTopic()
	if err != nil {
		log.Fatalln(err)
	}

	go consumer.InsertToDB()
	defer func() { consumer.Stop <- true }()

	fmt.Println("Consumer is up!")
	// API Gorilla
	server, err := api.NewApiServer(":5000", db)
	if err != nil {
		log.Fatalln(err)
	}
	err = server.Run()
	if err != nil {
		log.Fatalln(err)
	}

}
