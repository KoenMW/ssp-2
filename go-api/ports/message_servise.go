package ports

import "message-api/domain"

type MessageProducer interface {
	Send(msg domain.Message) error
}
