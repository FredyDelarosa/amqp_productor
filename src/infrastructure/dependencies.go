package infrastructure

import (
	"sensor/src/application"
	"sensor/src/core"
	//"sensor/src/infrastructure/repositories"
)

type Dependencies struct {
	CreateEventUseCase *application.CreateEventUseCase
}

func NewDependencies() (*Dependencies, error) {
	db, err := core.InitDB()
	if err != nil {
		return nil, err
	}

	mysqlRepo := NewMySQLEventRepository(db)

	return &Dependencies{
		CreateEventUseCase: application.NewCreateEventUseCase(mysqlRepo),
	}, nil
}
