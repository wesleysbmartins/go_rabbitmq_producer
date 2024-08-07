package queues

import (
	"go_rabbitmq_producer/internal/adapters/rabbitmq/factory/args"
	"time"
)

type Message struct {
	Headers       args.Args
	ContentType   string
	Body          interface{}
	DeliveryMode  uint8
	Priority      uint8
	CorrelationId string
	ReplyTo       string
	Expiration    string
	MessageId     string
	Timestamp     time.Time
	Type          string
	UserId        string
	AppId         string
}
