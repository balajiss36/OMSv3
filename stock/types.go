package main

import "context"

type Store interface {
	GetItems(context.Context) error
}

type StoreService interface {
	GetItemsList(context.Context) error
}
