package main

import (
	"context"
	"log"

	pb "github.com/balajiss36/omsv3/common/api"
	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/mgo.v2/bson"
)

type store struct {
	client *mongo.Client
}

func NewStore(client *mongo.Client) *store {
	return &store{
		client: client,
	}
}

func (s *store) GetItems(ctx context.Context, ids []string) ([]*pb.Item, error) {
	var stock []*pb.Item
	collection := s.client.Database("omsv3").Collection("stock")

	filter := bson.M{"ID": bson.M{"$in": ids}}
	cursor, err := collection.Find(context.Background(), filter)
	if err != nil {
		log.Printf("Error getting data order from cursor: %v", err)
	}

	err = cursor.All(ctx, &stock)
	if err != nil {
		log.Printf("Error getting order: %v", err)
	}

	defer cursor.Close(ctx)

	return stock, nil
}
