package main

import (
	"log"
	"net/http"

	"github.com/balajiss36/omsv3/common"
	pb "github.com/balajiss36/omsv3/common/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// const httpAddr = ":8080"
func main() {
	config, err := common.LoadConfig(".")
	if err != nil {
		log.Fatalf("Failed to load config: %v\n", err)
	}

	// call order grpc service
	conn, err := grpc.NewClient(config.GRPCAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("could not connect to order service: %v", err)
	}
	defer conn.Close()

	log.Println("connected to grpc gateway")

	c := pb.NewOrderServiceClient(conn)

	mux := http.NewServeMux()
	handler := NewHandler(c)
	handler.ServeHTTP(mux)
	log.Printf("Starting server on %s", config.HTTPAddress)
	if err := http.ListenAndServe(config.HTTPAddress, mux); err != nil {
		log.Fatalf("error starting server: %v", err)
	}
}
