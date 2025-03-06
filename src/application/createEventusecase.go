package application

import (
	"encoding/json"
	"log"
	"sensor/src/core"
	"sensor/src/domain/entities"
	"sensor/src/domain/repositories"

	"github.com/rabbitmq/amqp091-go"
)

type CreateEventUseCase struct {
	repo repositories.EventRepository
}

func NewCreateEventUseCase(repo repositories.EventRepository) *CreateEventUseCase {
	return &CreateEventUseCase{repo: repo}
}

func (uc *CreateEventUseCase) Execute(event *entities.Event) error {
	err := uc.repo.Create(event)
	if err != nil {
		return err
	}

	if core.RabbitChannel == nil {
		log.Println("Ni para conectar al rabbit sirves.")
		return nil
	}

	body, _ := json.Marshal(event)
	err = core.RabbitChannel.Publish(
		"", "sensor_alerts", false, false,
		amqp091.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)

	if err != nil {
		log.Println("Error publishing message to RabbitMQ", err)
	}
	return err
}
