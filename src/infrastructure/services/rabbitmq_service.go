package services

import (
	"context"
	"encoding/json"
	"log"
	"sensor/src/core"
	"sensor/src/domain/entities"
	"time"

	"github.com/rabbitmq/amqp091-go"
)

type RabbitMQService struct{}

func NewRabbitMQService() *RabbitMQService {
	return &RabbitMQService{}
}

func (s *RabbitMQService) PublishEvent(event *entities.Event) error {
	if core.RabbitChannel == nil {
		log.Println("no se conecto a rabbit")
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
		log.Println("error al enviar el evento a rabbit", err)
	} else {
		log.Println("evento enviado a rabbit", string(body))
	}

	return err
}
