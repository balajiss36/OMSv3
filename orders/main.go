package main

import (
	"context"
	"log"
	"net"

	"github.com/balajiss36/common"
	"google.golang.org/grpc"
)

var (
	grpcAddr   = common.EnvString("GRPC_ADDR", ":30052")
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

	store := NewStore()
	svc := NewService(store)
	NewGRPCHandler(grpcServer, svc)
	if err := svc.CreateOrder(context.Background()); err != nil {
		log.Fatalf("error creating order: %v", err)
	}

	log.Println("Starting server on", grpcAddr)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("error starting server: %v", err)
	}
}