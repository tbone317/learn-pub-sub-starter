package pubsub

import (
	"context"
	"encoding/json"
	"fmt"

	amqp "github.com/rabbitmq/amqp091-go"
)

func PublishJSON[T any](ch *amqp.Channel, exchange, key string, val T) error {
	jsonBytes, err := json.Marshal(val)
	if err != nil {
		return fmt.Errorf("could not marshal value to JSON: %w", err)
	}
	ctx := context.Background()
	return ch.PublishWithContext(ctx, exchange, key, false, false, amqp.Publishing{
		ContentType: "application/json",
		Body:        jsonBytes,
	})
}
