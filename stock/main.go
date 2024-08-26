package main

import (
	"log"
	"net"

	"github.com/balajiss36/common"
	"google.golang.org/grpc"
)

var grpcAddr = common.EnvString("GRPC_ADDR", ":30052")

func main() {
	lis, err := net.Listen("tcp", grpcAddr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	defer lis.Close()
	grpcServer := grpc.NewServer()

	store := NewStore()             // for db calls
	svc := NewService(store)        // for main logic to handle store requests
	NewGRPCHandler(grpcServer, svc) // grpc handler for store service
}
