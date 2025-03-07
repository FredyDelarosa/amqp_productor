package application

import (
	"sensor/src/domain/entities"
	"sensor/src/domain/repositories"
	"sensor/src/infrastructure/services"
)

type CreateEventUseCase struct {
	repo   repositories.EventRepository
	rabbit *services.RabbitMQService
}

func NewCreateEventUseCase(repo repositories.EventRepository, rabbit *services.RabbitMQService) *CreateEventUseCase {
	return &CreateEventUseCase{repo: repo, rabbit: rabbit}
}

func (uc *CreateEventUseCase) Execute(event *entities.Event) error {
	err := uc.repo.Create(event)
	if err != nil {
		return err
	}

	err = uc.rabbit.PublishEvent(event)
	return err
}
