package gateway

import (
	"context"

	pb "github.com/balajiss36/omsv3/common/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type gateway struct{}

func NewGateway() *gateway {
	return &gateway{}
}

func (g *gateway) UpdateOrder(ctx context.Context, order *pb.Order) error {
	conn, err := grpc.NewClient("orders:30052", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return err
	}
	orderClient := pb.NewOrderServiceClient(conn)
	_, err = orderClient.UpdateOrder(ctx, order)
	return err
}
