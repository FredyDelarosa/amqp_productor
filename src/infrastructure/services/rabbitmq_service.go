package services

import (
	"context"
	"encoding/json"
	"log"
	"sensor/src/core"
	"sensor/src/domain/entities"
	"sensor/src/domain/ports"
	"time"

	"github.com/rabbitmq/amqp091-go"
)

type RabbitMQPublisher struct{}

func NewRabbitMQPublisher() ports.EventPublisher {
	return &RabbitMQPublisher{}
}

func (r *RabbitMQPublisher) PublishEvent(event *entities.Event) error {
	if core.RabbitChannel == nil {
		log.Println("RabbitMQ no conectado")
		return nil
	}

	body, _ := json.Marshal(event)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := core.RabbitChannel.PublishWithContext(ctx,
		"", "sensor_alerts", false, false,
		amqp091.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)

	if err != nil {
		log.Println("Error publicando evento a RabbitMQ:", err)
	} else {
		log.Println("Evento enviado a RabbitMQ:", string(body))
	}

	return err
}
