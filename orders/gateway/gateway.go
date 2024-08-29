package gateway

import (
	"context"

	pb "github.com/balajiss36/common/api"
)

type CheckStockGateway interface {
	CheckIfItemInStock(ctx context.Context, items []*pb.ItemsWithQuantity) (bool, []*pb.Item, error)
}
