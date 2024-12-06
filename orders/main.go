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

	store := NewStore(client)
	svc := NewService(store)
	NewGRPCHandler(&config, grpcServer, svc)
	// if err := svc.CreateOrder(context.Background()); err != nil {
	// 	log.Fatalf("error creating order: %v", err)
	// }

	log.Println("Starting server on", config.GRPCAddress)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("error starting server: %v", err)
	}
}
