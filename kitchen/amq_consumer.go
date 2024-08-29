package main

import (
	"context"
	"encoding/json"
	"log"

	pb "github.com/balajiss36/common/api"
	common "github.com/balajiss36/common/broker"
	"github.com/balajiss36/kitchen/gateway"
)

type Consumer struct {
	gateway gateway.KitchenGateway
}

func NewConsumer(gateway gateway.KitchenGateway) *Consumer {
	return &Consumer{gateway}
}

func (g *Consumer) Listen(ctx context.Context) {
	ch, err := common.ConnectMQ(&common.RabbitMQ{
		Host:     mqHost,
		User:     mqUser,
		Password: mqPassword,
		Port:     mqPort,
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
