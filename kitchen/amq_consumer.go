package main

import (
	"context"
	"encoding/json"
	"log"

	common "github.com/balajiss36/omsv3/common"
	pb "github.com/balajiss36/omsv3/common/api"
	broker "github.com/balajiss36/omsv3/common/broker"
	"github.com/balajiss36/omsv3/kitchen/gateway"
)

type Consumer struct {
	config  *common.Config
	gateway gateway.KitchenGateway
}

func NewConsumer(config *common.Config, gateway gateway.KitchenGateway) *Consumer {
	return &Consumer{config, gateway}
}

func (g *Consumer) Listen(ctx context.Context) {
	ch, err := broker.ConnectMQ(&broker.RabbitMQ{
		Host:     g.config.RABBIT_MQ_HOST,
		User:     g.config.RABBIT_MQ_USER,
		Password: g.config.RABBIT_MQ_PASSWORD,
		Port:     g.config.RABBIT_MQ_PORT,
	})
	if err != nil {
		log.Printf("error connecting to rabbitmq: %v", err)
	}
	defer ch.Close()

	queue, err := ch.QueueDeclare("order.paid", false, false, false, false, nil)
	if err != nil {
		log.Fatalf("error declaring queue: %v", err)
	}

	msgs, err := ch.Consume(queue.Name, "", true, false, false, false, nil)
	if err != nil {
		log.Fatalf("error consuming queue: %v", err)
	}

	for msg := range msgs {
		var order *pb.Order

		if err := json.Unmarshal(msg.Body, &order); err != nil {
			log.Fatalf("error unmarshalling order: %v", err)
		}
		// Update order status
		if order.Status != "order paid" {
			continue
		}
		err := g.gateway.UpdateOrder(ctx, &pb.Order{
			Status:  "order ready",
			OrderID: order.OrderID,
			Items:   order.Items,
		})
		if err != nil {
			log.Printf("error updating order status: %v", err)
			continue
		}
	}
}
