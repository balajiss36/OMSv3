package main

import (
	"database/sql"
	"log"
	"net"
	"os"

	"github.com/balajiss36/common"
	"github.com/go-sql-driver/mysql"
	"google.golang.org/grpc"
)

var (
	grpcAddr   = common.EnvString("GRPC_ADDR", ":30056")
	mqPort     = common.EnvString("MQ_ADDR", ":5672")
	mqHost     = common.EnvString("MQ_HOST", "localhost")
	mqUser     = common.EnvString("MQ_USER", "user")
	mqPassword = common.EnvString("MQ_PASSWORD", "password")
	dbUser     = common.EnvString("DBUSER", "root")
	dbPass     = common.EnvString("DBPASS", "password")
)

func main() {
	lis, err := net.Listen("tcp", grpcAddr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	defer lis.Close()
	grpcServer := grpc.NewServer()

	cfg := mysql.Config{
		User:   os.Getenv("DBUSER"),
		Passwd: os.Getenv("DBPASS"),
		Net:    "tcp",
		Addr:   "localhost:3306",
		DBName: "oms",
	}
	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}

	store := NewStore(db)
	svc := NewService(store)
	NewGRPCHandler(grpcServer, svc)
	// if err := svc.CreateOrder(context.Background()); err != nil {
	// 	log.Fatalf("error creating order: %v", err)
	// }

	log.Println("Starting server on", grpcAddr)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("error starting server: %v", err)
	}
}
