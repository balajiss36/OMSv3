package gateway

import (
	"context"

	pb "github.com/balajiss36/omsv3/common/api"
)

type KitchenGateway interface {
	UpdateOrder(context.Context, *pb.Order) error
}
