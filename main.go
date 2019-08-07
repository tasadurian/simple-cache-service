package main

import (
	"log"
	"net"

	"github.com/thetommytwitch/simple-cache-service/backends/memory"
	pb "github.com/thetommytwitch/simple-cache-service/proto/go"
	"google.golang.org/grpc"
)

const (
	port = ":7771"
)

var server pb.CacheServer

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	server = memory.NewServer()

	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	log.Printf("listening on: %s", port)
	pb.RegisterCacheServer(grpcServer, server)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
