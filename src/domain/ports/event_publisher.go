package ports

import "sensor/src/domain/entities"

type EventPublisher interface {
	PublishEvent(event *entities.Event) error
}
