package main

import "context"

type store struct{}

func NewStore() *store {
	return &store{}
}

func (s *store) GetItems(context.Context) error {
	return nil
}
