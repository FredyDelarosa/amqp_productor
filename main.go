package main

import (
	"log"
	"sensor/src/core"
	"sensor/src/infrastructure"

	"github.com/gin-gonic/gin"
)

func main() {

	err := core.InitRabbitMQ()
	if err != nil {
		log.Fatal("Error al conectar a RabbitMQ:", err)
	}

	deps, err := infrastructure.NewDependencies()
	if err != nil {
		log.Fatal("Error inicializando dependencias", err)
	}

	r := gin.Default()

	infrastructure.RegisterRoutes(r, deps.CreateEventUseCase)

	r.Run(":8080")
}
