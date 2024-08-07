package main

import (
	"go_rabbitmq_producer/internal/adapters/rabbitmq"
	"go_rabbitmq_producer/internal/adapters/server"
	sale_usecase "go_rabbitmq_producer/internal/usecases/sale"
)

func init() {
	rabbitmq := rabbitmq.RabbitMQ{}
	rabbitmq.Connect()
}

func main() {
	u := sale_usecase.SaleUsecase{}
	u.CreateQueue()

	server := server.Server{}
	server.Run()
}
