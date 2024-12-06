package main

import (
	"context"

	pb "github.com/balajiss36/omsv3/common/api"
)

type service struct {
	store Store
}

func NewService(store Store) *service {
	return &service{store}
}

func (s *service) GetItemsList(ctx context.Context, ids []string) ([]*pb.Item, error) {
	return s.store.GetItems(ctx, ids)
}

func (s *service) CheckIfItemAreInStock(ctx context.Context, items []*pb.ItemsWithQuantity) (bool, []*pb.Item, error) {
	itemIDs := make([]string, 0)
	for _, item := range items {
		itemIDs = append(itemIDs, item.ID)
	}
	itemList, err := s.store.GetItems(ctx, itemIDs)
	if err != nil {
		return false, nil, err
	}
	newItems := []*pb.Item{}
	for _, item := range itemList {
		for _, itemWithQuantity := range items {
			if item.ID == itemWithQuantity.ID {
				newItems = append(newItems, item)
			}
		}
	}
	return false, newItems, nil
}
