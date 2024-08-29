package processor

import (
	"context"

	pb "github.com/balajiss36/common/api"
)

type PaymentProcessor interface {
	CreatePaymentLink(ctx context.Context, order *pb.Order) (string, error)
}
