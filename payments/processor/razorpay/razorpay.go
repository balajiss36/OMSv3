package razorpay

import (
	"context"

	pb "github.com/balajiss36/common/api"
	razorpay "github.com/razorpay/razorpay-go"
)

type Stripe struct{}

func NewRazorProcessor() *Stripe {
	return &Stripe{}
}

func (s *Stripe) CreatePaymentLink(ctx context.Context, order *pb.Order) (string, error) {
	client := razorpay.NewClient("g_90ik34rtg", "gasf34t34t")

	data := map[string]interface{}{
		"upi_link":     "true",
		"amount":       1000,
		"currency":     "INR",
		"callback_url": "https://localhost:9097/webhook",
	}

	body, err := client.PaymentLink.Create(data, nil)
	if err != nil {
		return "", err
	}

	return body["id"].(string), nil
}
