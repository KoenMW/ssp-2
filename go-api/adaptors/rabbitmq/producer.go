package rabbitmq

import (
	"context"
	"encoding/json"
	"message-api/domain"
	"os"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitMQProducer struct {
	conn  *amqp.Connection
	ch    *amqp.Channel
	queue amqp.Queue
}

func NewProducer() (*RabbitMQProducer, error) {
	conn, err := amqp.Dial(os.Getenv("RABBITMQ_URL"))
	if err != nil {
		return nil, err
	}

	ch, err := conn.Channel()
	if err != nil {
		conn.Close()
		return nil, err
	}

	q, err := ch.QueueDeclare(
		os.Getenv("QUEUE_NAME"), false, true, false, false, nil,
	)
	if err != nil {
		ch.Close()
		conn.Close()
		return nil, err
	}

	return &RabbitMQProducer{conn: conn, ch: ch, queue: q}, nil
}

func (p *RabbitMQProducer) Send(msg domain.Message) error {
	jsonBody, err := json.Marshal(msg)
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	return p.ch.PublishWithContext(ctx,
		"",
		p.queue.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        jsonBody,
		},
	)
}
