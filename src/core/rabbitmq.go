package core

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/rabbitmq/amqp091-go"
)

var RabbitChannel *amqp091.Channel

func InitRabbitMQ() error {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error cargando el archivo .env: %v", err)
	}

	url := fmt.Sprintf("amqp://%s:%s@%s:%s/",
		os.Getenv("RABBIT_USER"),
		os.Getenv("RABBIT_PASSWORD"),
		os.Getenv("RABBIT_HOST"),
		os.Getenv("RABBIT_PORT"),
	)

	log.Println("Por favor rabbit conectate", url)

	conn, err := amqp091.Dial(url)
	if err != nil {
		log.Fatalf("No se conecto: %v", err)
		return err
	}

	RabbitChannel, err = conn.Channel()
	if err != nil {
		log.Fatalf("ni para crear un canal sirves: %v", err)
		return err
	}

	// Verificar si el canal es válido
	if RabbitChannel == nil {
		log.Fatalf("papi es tu codigo que no sirve")
	}

	fmt.Println("na hermano eres muy bueno")
	return nil
}
