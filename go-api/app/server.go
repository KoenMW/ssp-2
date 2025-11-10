package main

import (
	"log"
	"message-api/adaptors/rabbitmq"
	"message-api/adaptors/rest"
	"net/http"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Println(".env file not found, relying on environment variables")
	}

	producer, err := rabbitmq.NewProducer()
	if err != nil {
		log.Fatalf("failed to start producer: %v", err)
	}

	handler := &rest.Handler{Producer: producer}
	http.HandleFunc("/messages", handler.Messages)
	http.ListenAndServe(":8080", nil)
}
