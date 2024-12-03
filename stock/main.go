package main

import (
	"log"
	"net"

	"github.com/balajiss36/common"
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

	store := NewStore()                      // for db calls
	svc := NewService(store)                 // for main logic to handle store requests
	_, err = NewGRPCHandler(grpcServer, svc) // grpc handler for store service
	if err != nil {
		return
	}
}
