package application

import (
	"sensor/src/domain/entities"
	"sensor/src/domain/ports"
	"sensor/src/domain/repositories"
)

type CreateEventUseCase struct {
	repo      repositories.EventRepository
	publisher ports.EventPublisher
}

func NewCreateEventUseCase(repo repositories.EventRepository, publisher ports.EventPublisher) *CreateEventUseCase {
	return &CreateEventUseCase{repo: repo, publisher: publisher}
}

func (uc *CreateEventUseCase) Execute(event *entities.Event) error {
	if err := uc.repo.Create(event); err != nil {
		return err
	}

	return uc.publisher.PublishEvent(event)
}
