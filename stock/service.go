package main

import "context"

type service struct {
	store Store
}

func NewService(store Store) *service {
	return &service{store: store}
}

func (s *service) GetItemsList(ctx context.Context) error {
	return nil
}
