package main

import (
	"context"
	"fmt"
	"log"

	pb "github.com/balajiss36/common/api"
	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/mgo.v2/bson"
)

type store struct {
	client *mongo.Client
}

func NewStore(clientDB *mongo.Client) *store {
	return &store{client: clientDB}
}

func (s *store) Create(ctx context.Context, order *pb.Order) error {
	collection := s.client.Database("omsv3").Collection("orders")

	_, err := collection.InsertOne(context.Background(), order)
	if err != nil {
		fmt.Errorf("Error inserting order: %v", err)
	}

	log.Println("Order Created")
	return nil
}

func (s *store) Update(ctx context.Context, id string, order *pb.Order) (*pb.Order, error) {
	collection := s.client.Database("omsv3").Collection("orders")

	filter := bson.M{"ID": id}
	update := bson.M{"$set": order}
	mongoResult := collection.FindOneAndUpdate(context.Background(), filter, update)
	if mongoResult.Err() != nil {
		fmt.Errorf("Error updating order: %v", mongoResult.Err())
	}

	log.Println("Order Updated")

	return nil, nil
}

func (s *store) Get(ctx context.Context, orderid, customerID string) (*pb.Order, error) {
	var order pb.Order

	collection := s.client.Database("omsv3").Collection("orders")

	filter := bson.M{"ID": customerID}
	cursor, err := collection.Find(context.Background(), filter)
	if err != nil {
		fmt.Errorf("Error getting data order from cursor: %v", err)
	}

	err = cursor.All(context.Background(), &order)
	if err != nil {
		fmt.Errorf("Error getting order: %v", err)
	}
	defer cursor.Close(context.Background())

	log.Println("Order Retrieved")

	return &order, nil
}
