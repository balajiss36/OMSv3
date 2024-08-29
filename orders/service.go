package main

import (
	"context"
	"errors"

	pb "github.com/balajiss36/common/api"
	"github.com/balajiss36/orders/gateway"
)

type service struct {
	store   OrdersStore
	gateway gateway.CheckStockGateway
}

func NewService(store OrdersStore) *service {
	return &service{store: store}
}

// implement the CreateOrder Method defined in types.go
func (s *service) CreateOrder(ctx context.Context, p *pb.CreateOrderRequest, items []*pb.Item) (*pb.Order, error) {
	err := s.store.Create(ctx, &pb.Order{
		Items:      items,
		CustomerID: p.CustomerID,
	})
	if err != nil {
		return nil, err
	}

	return &pb.Order{
		Items:      items,
		Status:     "order pending",
		CustomerID: p.CustomerID,
	}, nil
}

func (s *service) GetOrder(ctx context.Context, orderid, customerID string) (*pb.Order, error) {
	order, err := s.store.Get(ctx, orderid, customerID)
	if err != nil {
		return nil, err
	}

	return order, nil
}

func (s *service) ValidateOrder(ctx context.Context, pb *pb.CreateOrderRequest) ([]*pb.Item, error) {
	if len(pb.Items) == 0 {
		return nil, errors.New("items must not be empty")
	}
	itemQuantities := mergeItemsQuantities(pb.Items)
	for _, item := range itemQuantities {
		if item.ID == "" {
			return nil, errors.New("item ID must not be empty")
		}
	}

	isStock, items, err := s.gateway.CheckIfItemInStock(ctx, itemQuantities)
	if err != nil {
		return nil, err
	}
	if !isStock {
		return nil, errors.New("item not in stock")
	}
	return items, nil
}

func mergeItemsQuantities(items []*pb.ItemsWithQuantity) []*pb.ItemsWithQuantity {
	itemQuantities := make([]*pb.ItemsWithQuantity, 0)
	for _, item := range items {
		found := false
		for _, finalItem := range itemQuantities {
			if finalItem.ID == item.ID {
				finalItem.Quantity += item.Quantity
				found = true
				break
			}
		}
		if !found {
			itemQuantities = append(itemQuantities, item)
		}
	}
	return itemQuantities
}

func (s *service) UpdateOrder(ctx context.Context, items *pb.Order) (*pb.Order, error) {
	_, err := s.store.Update(ctx, items.OrderID, items)
	if err != nil {
		return nil, err
	}

	return nil, nil
}
