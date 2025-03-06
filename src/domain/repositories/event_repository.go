package repositories

import "sensor/src/domain/entities"

type EventRepository interface {
	Create(event *entities.Event) error
	GetAll() ([]entities.Event, error)
	GetByID(id int) (*entities.Event, error)
}
