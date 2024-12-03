package main

import (
	"log"
	"net"
	"net/http"

	"github.com/balajiss36/common"
	broker "github.com/balajiss36/common/broker"
	"github.com/balajiss36/payments/gateway"
	"github.com/balajiss36/payments/processor/razorpay"
	"google.golang.org/grpc"
)

func main() {
	config, err := common.LoadConfig(".")
	if err != nil {
		log.Fatalf("Failed to load config: %v\n", err)
	}

	lis, err := net.Listen("tcp", config.GRPCAddress)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	defer lis.Close()
	grpcServer := grpc.NewServer()

	gw := gateway.NewGateway()
	razor := razorpay.NewRazorProcessor()

	svc := NewService(gw, razor)
	// NewGRPCHandler(grpcServer, svc)

	NewTelemetryMiddleware(svc)

	ch, err := broker.ConnectMQ(&broker.RabbitMQ{
		Host:     config.RABBIT_MQ_HOST,
		User:     config.RABBIT_MQ_USER,
		Password: config.RABBIT_MQ_PASSWORD,
		Port:     config.RABBIT_MQ_PORT,
	})
	if err != nil {
		log.Printf("error connecting to rabbitmq: %v", err)
	}
	defer ch.Close()

	h := NewPaymentHTTPHandler(ch)

	http.HandleFunc("/webhook", h.handleWebhook)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("error starting server: %v", err)
	}
}
