package main

import (
	"context"

	pb "github.com/balajiss36/common/api"
)

type Store interface {
	GetItems(ctx context.Context, ids []string) ([]*pb.Item, error)
}

type StoreService interface {
	CheckIfItemAreInStock(context.Context, []*pb.ItemsWithQuantity) (bool, []*pb.Item, error)
	GetItemsList(ctx context.Context, ids []string) ([]*pb.Item, error)
}
