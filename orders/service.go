package main

import (
	"context"
	"errors"

	pb "github.com/balajiss36/common/api"
)

type service struct {
	store OrdersStore
}

func NewService(store OrdersStore) *service {
	return &service{store: store}
}

// implement the CreateOrder Method defined in types.go
func (s *service) CreateOrder(ctx context.Context) error {
	return nil
}

func (s *service) ValidateOrder(ctx context.Context, pb *pb.CreateOrderRequest) error {
	if len(pb.Items) == 0 {
		return errors.New("items must not be empty")
	}
	itemQuantities := mergeItemsQuantities(pb.Items)
	for _, item := range itemQuantities {
		if item.ID == "" {
			return errors.New("item ID must not be empty")
		}
	}
	return nil
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
