package main

import (
	"context"
	"log"
	"time"

	pb "github.com/thetommytwitch/simple-cache-service/proto/go"
	"google.golang.org/grpc"
)

const (
	address = "localhost:7771"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewCacheClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	_, err = c.Get(ctx, &pb.GetRequest{Key: "test_key"})
	if err != nil {
		log.Fatalf("could not get: %v", err)
	}
}
