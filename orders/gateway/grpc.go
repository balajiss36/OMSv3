package gateway

import (
	"context"
	"log"

	pb "github.com/balajiss36/omsv3/common/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type gateway struct{}

func NewGateway() *gateway {
	return &gateway{}
}

func (g *gateway) CheckIfItemInStock(ctx context.Context, items []*pb.ItemsWithQuantity) (bool, []*pb.Item, error) {
	// grpc client call to stock service
	conn, err := grpc.NewClient("stock:30056", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return false, nil, err
	}

	defer conn.Close()

	c := pb.NewStockServiceClient(conn)

	list, err := c.CheckItems(ctx, &pb.CheckItemsRequest{Items: items})
	if err != nil {
		log.Fatalf("error checking items from stock: %v", err)
		return false, nil, err
	}

	return list.IsStock, list.Items, nil
}
