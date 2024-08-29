package main

import (
	"context"
	"database/sql"
	"log"

	pb "github.com/balajiss36/common/api"
)

type store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *store {
	return &store{db: db}
}

func (s *store) Create(ctx context.Context, order *pb.Order) error {
	_, err := s.db.Exec("INSERT INTO orders (ID, CustomerID, STATUS) VALUES (?, ?, ?)", order.OrderID, order.CustomerID, order.Status)
	if err != nil {
		return err
	}
	log.Println("Order Created")
	return nil
}

func (s *store) Update(ctx context.Context, id string, order *pb.Order) (*pb.Order, error) {
	_, err := s.db.Exec("UPDATE orders SET STATUS = ?, PaymentLink = ? WHERE ID = ?", order.Status, order.PaymentLink, id)
	if err != nil {
		return nil, err
	}

	log.Println("Order Updated")

	return nil, nil
}

func (s *store) Get(ctx context.Context, orderid, customerID string) (*pb.Order, error) {
	var order pb.Order
	err := s.db.QueryRow("SELECT ID, STATUS, PaymentLink FROM orders WHERE ID = ?, CustomerID = ?", orderid, customerID).Scan(&order.OrderID, &order.Status, &order.PaymentLink)
	if err != nil {
		return nil, err
	}

	log.Println("Order Retrieved")

	return &order, nil
}
