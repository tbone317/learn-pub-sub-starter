package main

import (
	"fmt"
	"log"

	"github.com/bootdotdev/learn-pub-sub-starter/internal/pubsub"
	"github.com/bootdotdev/learn-pub-sub-starter/internal/routing"

	ampq "github.com/rabbitmq/amqp091-go"
)

func main() {
	fmt.Println("Starting Peril server...")
	rabbitConnString := "amqp://guest:guest@localhost:5672/"
	conn, err := ampq.Dial(rabbitConnString)
	if err != nil {
		log.Fatalf("could not connect to RabbitMQ: %v", err)
	}
	defer conn.Close()
	publishCh, err := conn.Channel()
	if err != nil {
		log.Fatalf("could not create RabbitMQ channel: %v", err)
	}
	defer publishCh.Close()
	fmt.Println("Connected to RabbitMQ successfully!")

	pubsub.PublishJSON(publishCh,
		routing.ExchangePerilDirect,
		routing.PauseKey,
		routing.PlayingState{IsPaused: true},
	)
}
