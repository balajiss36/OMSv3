package main

import (
	"context"

	pb "github.com/balajiss36/common/api"
)

type Payments interface {
	CreatePayment(context.Context, *pb.Order) (string, error)
}
