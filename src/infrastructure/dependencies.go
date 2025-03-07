package infrastructure

import (
	"sensor/src/application"
	"sensor/src/core"
	"sensor/src/infrastructure/services"
)

type Dependencies struct {
	CreateEventUseCase *application.CreateEventUseCase
}

func NewDependencies() (*Dependencies, error) {
	db, err := core.InitDB()
	if err != nil {
		return nil, err
	}

	rabbitService := services.NewRabbitMQService()
	mysqlRepo := NewMySQLEventRepository(db)

	return &Dependencies{
		CreateEventUseCase: application.NewCreateEventUseCase(mysqlRepo, rabbitService),
	}, nil
}
