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

func (h *grpcHandler) CreateOrder(ctx context.Context, p *pb.CreateOrderRequest) (*pb.Order, error) {
	log.Println("CreateOrder called")

	items, err := h.service.ValidateOrder(context.Background(), &pb.CreateOrderRequest{})
	if err != nil {
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

	queue, err := ch.QueueDeclare("order.created", false, false, false, false, nil)
	if err != nil {
		return nil, err
	}

	newOrder, err := h.service.CreateOrder(context.Background(), p, items)
	if err != nil {
		return nil, err
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

func (h *grpcHandler) UpdateOrder(ctx context.Context, p *pb.Order) (*pb.Order, error) {
	return h.service.UpdateOrder(ctx, p)
}

func (h *grpcHandler) GetOrder(ctx context.Context, p *pb.GetOrderRequest) (*pb.Order, error) {
	return h.service.GetOrder(ctx, p.OrderID, p.CustomerID)
}
