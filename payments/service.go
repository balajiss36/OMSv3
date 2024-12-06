package main

import (
	"context"

	pb "github.com/balajiss36/omsv3/common/api"
	"github.com/balajiss36/omsv3/payments/gateway"
	"github.com/balajiss36/omsv3/payments/processor"
)

type service struct {
	orders    gateway.PaymentGateway
	processor processor.PaymentProcessor
}

func NewService(orders gateway.PaymentGateway, processor processor.PaymentProcessor) *service {
	return &service{orders: orders, processor: processor}
}

func (s *service) CreatePayment(ctx context.Context, order *pb.Order) (string, error) {
	link, err := s.processor.CreatePaymentLink(ctx, order)
	if err != nil {
		return "", err
	}

	err = s.orders.UpdateOrderAfterPayment(ctx, order.OrderID, link)
	if err != nil {
		return "", err
	}

	return link, nil
}
