package main

import (
	"log"
	"net"
	"net/http"

	"github.com/balajiss36/common"
	common1 "github.com/balajiss36/common/broker"
	"github.com/balajiss36/payments/gateway"
	"github.com/balajiss36/payments/processor/razorpay"
	"google.golang.org/grpc"
)

var (
	grpcAddr   = common.EnvString("GRPC_ADDR", ":30055")
	mqPort     = common.EnvString("MQ_ADDR", ":5672")
	mqHost     = common.EnvString("MQ_HOST", "localhost")
	mqUser     = common.EnvString("MQ_USER", "user")
	mqPassword = common.EnvString("MQ_PASSWORD", "password")
)

func main() {
	lis, err := net.Listen("tcp", grpcAddr)
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

	ch, err := common1.ConnectMQ(&common1.RabbitMQ{
		Host:     mqHost,
		User:     mqUser,
		Password: mqPassword,
		Port:     mqPort,
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
