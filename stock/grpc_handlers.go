package main

import (
	"context"

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

func (h *grpcHandler) GetItemsList(ctx context.Context, req *pb.GetItemsRequest) (*pb.GetItemsResponse, error) {
	items, err := h.service.GetItemsList(ctx, req.ItemIDs)
	if err != nil {
		return nil, err
	}

	return &pb.GetItemsResponse{Items: items}, nil
}

func (h *grpcHandler) CheckItems(ctx context.Context, req *pb.CheckItemsRequest) (*pb.CheckItemsResponse, error) {
	isStock, items, err := h.service.CheckIfItemAreInStock(ctx, req.Items)
	if err != nil {
		return nil, err
	}
	if !isStock {
		return &pb.CheckItemsResponse{IsStock: false}, nil
	}

	return &pb.CheckItemsResponse{IsStock: true, Items: items}, nil
}
