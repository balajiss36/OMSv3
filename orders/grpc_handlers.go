package main

import (
	"context"
	"encoding/json"
	"log"

	pb "github.com/balajiss36/common/api"
	common "github.com/balajiss36/common/broker"
	"github.com/streadway/amqp"
	"google.golang.org/grpc"
)

type grpcHandler struct {
	// this truct implements Order Service interface
	pb.UnimplementedOrderServiceServer

	service OrdersService
}

func NewGRPCHandler(grpcServer *grpc.Server, service OrdersService) *grpcHandler {
	handler := &grpcHandler{
		service: service,
	}
	pb.RegisterOrderServiceServer(grpcServer, handler)
	return handler
}

func (h *grpcHandler) CreateOrder(context.Context, *pb.CreateOrderRequest) (*pb.Order, error) {
	log.Println("CreateOrder called")

	if err := h.service.ValidateOrder(context.Background(), &pb.CreateOrderRequest{}); err != nil {
		return nil, err
	}

	ch, err := common.ConnectMQ(&common.RabbitMQ{
		Host:     mqHost,
		User:     mqUser,
		Password: mqPassword,
		Port:     mqPort,
	})
	if err != nil {
		return nil, err
	}
	defer ch.Close()

	queue, err := ch.QueueDeclare("new.order", false, false, false, false, nil)
	if err != nil {
		return nil, err
	}

	newOrder := &pb.Order{
		OrderID: "368",
		Items: []*pb.Item{
			{
				ID:       "123",
				Name:     "item1",
				Quantity: 1,
			},
		},
	}
	newOrderBody, err := json.Marshal(newOrder)
	if err != nil {
		return nil, err
	}

	err = ch.Publish("", queue.Name, false, false, amqp.Publishing{
		ContentType: "text/plain",
		Body:        newOrderBody,
	})
	if err != nil {
		return nil, err
	}
	log.Println("Order created")

	return newOrder, nil
}
