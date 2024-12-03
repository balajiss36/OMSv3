package main

import (
	"context"
	"log"
	"net"

	"github.com/balajiss36/common"
	"github.com/balajiss36/kitchen/gateway"
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

	gw := gateway.NewGateway()
	conn := NewConsumer(&config, gw)

	go conn.Listen(context.Background())
	grpcServer := grpc.NewServer()
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal("failed to serve", err)
	}
}
