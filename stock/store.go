package main

import (
	"context"

	pb "github.com/balajiss36/common/api"
)

type store struct{}

func NewStore() *store {
	return &store{}
}

func (s *store) GetItems(ctx context.Context, ids []string) ([]*pb.Item, error) {
	return nil, nil
}
