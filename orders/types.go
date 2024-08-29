package main

import (
	"context"

	pb "github.com/balajiss36/common/api"
)

type OrdersService interface {
	GetOrder(ctx context.Context, orderid, customerID string) (*pb.Order, error)
	CreateOrder(ctx context.Context, p *pb.CreateOrderRequest, items []*pb.Item) (*pb.Order, error)
	ValidateOrder(context.Context, *pb.CreateOrderRequest) ([]*pb.Item, error)
	UpdateOrder(context.Context, *pb.Order) (*pb.Order, error)
}

type OrdersStore interface {
	Get(ctx context.Context, orderid, customerID string) (*pb.Order, error)
	Create(ctx context.Context, order *pb.Order) error
	Update(ctx context.Context, id string, order *pb.Order) (*pb.Order, error)
}
