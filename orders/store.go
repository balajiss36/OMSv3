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

func NewStore(clientDB *mongo.Client) *store {
	return &store{client: clientDB}
}

func (s *store) Create(ctx context.Context, order *pb.Order) error {
	collection := s.client.Database("omsv3").Collection("orders")

	_, err := collection.InsertOne(context.Background(), order)
	if err != nil {
		log.Fatalf("Error inserting order: %v", err)
		return err
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
		log.Fatalf("Error updating order: %v", mongoResult.Err())
		return nil, mongoResult.Err()
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
		log.Fatalf("Error getting data order from cursor: %v", err)
		return nil, err
	}

	err = cursor.All(context.Background(), &order)
	if err != nil {
		log.Fatalf("Error getting order: %v", err)
		return nil, err
	}
	defer cursor.Close(context.Background())

	log.Println("Order Retrieved")

	return &order, nil
}
