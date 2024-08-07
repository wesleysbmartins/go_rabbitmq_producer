package entities

import (
	"go_rabbitmq_producer/internal/adapters/rabbitmq/factory/args"
)

type Producer struct {
	Exchange struct {
		Name       string    `yaml:"name"`
		Kind       string    `yaml:"kind"`
		RoutingKey string    `yaml:"routing-key"`
		Durable    bool      `yaml:"durable"`
		AutoDelete bool      `yaml:"auto-delete"`
		Internal   bool      `yaml:"internal"`
		NoWait     bool      `yaml:"no-wait"`
		Args       args.Args `yaml:"args"`
		Bind       string    `yaml:"bind"`
	} `yaml:"exchange"`
	Queue struct {
		Name       string    `yaml:"name"`
		Durable    bool      `yaml:"durable"`
		Exclusive  bool      `yaml:"exclusive"`
		AutoDelete bool      `yaml:"auto-delete"`
		NoWait     bool      `yaml:"no-wait"`
		Args       args.Args `yaml:"args"`
		Bind       string    `yaml:"bind"`
	} `yaml:"queue"`
	Message struct {
		Headers       args.Args `yaml:"headers"`
		ContentType   string    `yaml:"content-type"`
		DeliveryMode  uint8     `yaml:"delivery-mode"`
		Priority      uint8     `yaml:"priority"`
		CorrelationId string    `yaml:"correlation-id"`
		ReplyTo       string    `yaml:"reply-to"`
		Expiration    string    `yaml:"expiration"`
		Type          string    `yaml:"type"`
	} `yaml:"message"`
}
