package common

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func SetupMongoDB(ctx context.Context, config Config) (*mongo.Client, error) {
	uri := fmt.Sprintf("mongodb://%s:%s@%s.%s.svc.cluster.local%s", config.MONGO_USER, config.MONGO_PASSWORD, config.MONGO_SRV, config.MONGO_NAMESPACE, config.MONGO_PORT)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		log.Panicf("unable to connect to MongoDB, err - %s", err)
		return nil, err
	}
	err = client.Ping(ctx, readpref.SecondaryPreferred())
	if err != nil {
		log.Panicf("unable to ping to MongoDB, err - %s", err)
		return nil, err
	}
	return client, nil
}

// Close the connection
func CloseConnection(context context.Context, client *mongo.Client) {
	defer func() {
		if err := client.Disconnect(context); err != nil {
			log.Panic(err)
		}
		fmt.Println("Close connection is called")
	}()
}
