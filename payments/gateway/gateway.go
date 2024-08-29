package gateway

import (
	"context"
)

type PaymentGateway interface {
	UpdateOrderAfterPayment(ctx context.Context, orderID, paymentLink string) error
}
