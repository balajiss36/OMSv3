package main

import (
	"context"
	"log"
	"net"

	"github.com/balajiss36/common"
	"github.com/balajiss36/kitchen/gateway"
	"google.golang.org/grpc"
)

var (
	grpcAddr = common.EnvString("GRPC_ADDR", ":30058")

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

	gw := gateway.NewGateway()
	conn := NewConsumer(gw)

	go conn.Listen(context.Background())
	grpcServer := grpc.NewServer()
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal("failed to serve", err)
	}
}
