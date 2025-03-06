package main

import (
	"log"
	"sensor/src/core"
	"sensor/src/infrastructure"

	"github.com/gin-gonic/gin"
)

func main() {
	deps, err := infrastructure.NewDependencies()
	if err != nil {
		log.Fatal("Error inicializando dependencias", err)
	}

	r := gin.Default()

	infrastructure.RegisterRoutes(r, deps.CreateEventUseCase)

	log.Println("iniciando rabbit")
	err = core.InitRabbitMQ()
	if err != nil {
		log.Fatal("el rabbit no jala", err)
	}

	r.Run(":8080")
}
