package sale_usecase

import (
	"fmt"
	configs "go_rabbitmq_producer/config"
	"go_rabbitmq_producer/internal/adapters/rabbitmq/factory/exchanges"
	"go_rabbitmq_producer/internal/adapters/rabbitmq/factory/queues"
	"go_rabbitmq_producer/internal/entities"
)

type SaleUsecase struct{}

type ISaleUsecase interface {
	CreateQueue()
	Produce()
}

var p *entities.Producer

func (u *SaleUsecase) CreateQueue() {
	config := configs.Config{}
	config.Load("sale", &p)

	exchange := exchanges.Exchange{}

	exchange.Create(p.Exchange.Name, p.Exchange.Kind, p.Exchange.Durable, p.Exchange.AutoDelete, p.Exchange.Internal, p.Exchange.NoWait, p.Exchange.Args)

	if p.Exchange.Bind != "" {
		exchange.Bind(p.Exchange.Bind, p.Exchange.RoutingKey, p.Exchange.Name, p.Exchange.NoWait, p.Exchange.Args)
	}

	queue := queues.Queue{}

	queue.Create(p.Queue.Name, p.Queue.Durable, p.Queue.AutoDelete, p.Queue.Exclusive, p.Queue.NoWait, p.Queue.Args)

	queue.Bind(p.Queue.Name, p.Exchange.RoutingKey, p.Exchange.Name, p.Queue.NoWait, p.Queue.Args)

}

func (u *SaleUsecase) Produce(dto SaleMessageDTO) error {

	queue := queues.Queue{}

	message := queues.Message{
		ContentType: "application/json",
		Body:        dto.Sale,
	}

	err := queue.Produce(p.Exchange.Name, p.Exchange.RoutingKey, false, false, message)
	if err != nil {
		panic(fmt.Sprintf("ERROR TO PUBLISH IN QUEUE: %s", err.Error()))
	}

	return err
}
