package queues

import (
	"encoding/json"
	"fmt"
	"go_rabbitmq_producer/internal/adapters/rabbitmq"
	"go_rabbitmq_producer/internal/adapters/rabbitmq/factory/args"

	"github.com/streadway/amqp"
)

type Queue struct{}

type IQueue interface {
	Create(name string, durable bool, autoDelete bool, exclusive bool, noWait bool, args args.Args)
	Bind(routingKey string, exchangeName string, noWait bool, args args.Args)
	Produce(exchangeName string, routingKey string, mandatory bool, immediate bool, message Message) error
}

func (q *Queue) Create(name string, durable bool, autoDelete bool, exclusive bool, noWait bool, args args.Args) {
	if name == "" {
		panic("Param name is required!")
	}

	rabbitmq := rabbitmq.RabbitMQ{}
	channel := rabbitmq.Channel()
	_, err := channel.QueueDeclare(name, durable, autoDelete, exclusive, noWait, args.Handle())
	if err != nil {
		panic(fmt.Sprintf("RABBITMQ DECLARE QUEUE ERROR\n%s", err.Error()))
	}
}

func (q *Queue) Bind(queueName string, routingKey string, exchangeName string, noWait bool, args args.Args) {
	rabbitmq := rabbitmq.RabbitMQ{}
	channel := rabbitmq.Channel()
	err := channel.QueueBind(queueName, routingKey, exchangeName, noWait, args.Handle())
	if err != nil {
		panic(fmt.Sprintf("RABBITMQ BINDING QUEUE ERROR\n%s", err.Error()))
	}
}

func (q *Queue) Produce(exchangeName string, routingKey string, mandatory bool, immediate bool, message Message) error {
	rabbitmq := rabbitmq.RabbitMQ{}
	channel := rabbitmq.Channel()

	body, _ := json.Marshal(message.Body)

	publishing := amqp.Publishing{
		Headers:       message.Headers.Handle(),
		ContentType:   message.ContentType,
		Body:          body,
		DeliveryMode:  message.DeliveryMode,
		Priority:      message.Priority,
		CorrelationId: message.CorrelationId,
		ReplyTo:       message.ReplyTo,
		Expiration:    message.Expiration,
		MessageId:     message.MessageId,
		Timestamp:     message.Timestamp,
		Type:          message.Type,
		UserId:        message.UserId,
		AppId:         message.AppId,
	}

	err := channel.Publish(exchangeName, routingKey, mandatory, immediate, publishing)
	if err != nil {
		panic(fmt.Sprintf("RABBITMQ PUBLISH IN QUEUE DIRECT ERROR\n%s", err.Error()))
	}

	return err
}
