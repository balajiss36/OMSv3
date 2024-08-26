package main

import (
	pb "github.com/balajiss36/common/api"
	"google.golang.org/grpc"
)

type grpcHandler struct {
	// this truct implements Order Service interface
	// pb.UnimplementedOrderServiceServer

	service StoreService

	pb.UnimplementedStockServiceServer
}

func NewGRPCHandler(grpcServer *grpc.Server, service StoreService) (*grpcHandler, error) {
	handler := &grpcHandler{
		service: service,
	}
	// pb.RegisterOrderServiceServer(grpcServer, handler)
	// return handler

	pb.RegisterStockServiceServer(grpcServer, handler)

	return nil, nil
}
