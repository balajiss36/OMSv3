package main

import (
	"context"
	"log"
	"net"

	"github.com/balajiss36/omsv3/common"
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

	client, err := common.SetupMongoDB(context.Background(), config)
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v\n", err)
	}

	defer common.CloseConnection(context.Background(), client)

	store := NewStore(client)                // for db calls
	svc := NewService(store)                 // for main logic to handle store requests
	_, err = NewGRPCHandler(grpcServer, svc) // grpc handler for store service
	if err != nil {
		return
	}
}
