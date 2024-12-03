package main

import (
	"context"
	"encoding/json"
	"log"

	common "github.com/balajiss36/common"
	pb "github.com/balajiss36/common/api"
	broker "github.com/balajiss36/common/broker"
	"github.com/streadway/amqp"
	"google.golang.org/grpc"
)

type grpcHandler struct {
	// this truct implements Order Service interface
	pb.UnimplementedOrderServiceServer
	config  *common.Config
	service OrdersService
}

func NewGRPCHandler(config *common.Config, grpcServer *grpc.Server, service OrdersService) *grpcHandler {
	handler := &grpcHandler{
		config:  config,
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

	ch, err := broker.ConnectMQ(&broker.RabbitMQ{
		Host:     h.config.RABBIT_MQ_HOST,
		User:     h.config.RABBIT_MQ_USER,
		Password: h.config.RABBIT_MQ_PASSWORD,
		Port:     h.config.RABBIT_MQ_PORT,
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
