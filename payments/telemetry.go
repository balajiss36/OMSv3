package main

import (
	"context"
	"fmt"

	pb "github.com/balajiss36/common/api"
	"go.opentelemetry.io/otel/trace"
)

type TelemetryMiddleware struct {
	next Payments
}

func NewTelemetryMiddleware(next Payments) Payments {
	return &TelemetryMiddleware{next}
}

func (s *TelemetryMiddleware) CreatePayment(ctx context.Context, o *pb.Order) (string, error) {
	span := trace.SpanFromContext(ctx)
	span.AddEvent(fmt.Sprintf("CreatePayment: %v", o))

	return s.next.CreatePayment(ctx, o)
}
