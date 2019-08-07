package main

import (
	"log"
	"net"

	"github.com/thetommytwitch/simple-cache-service/backends/memory"
	"github.com/thetommytwitch/simple-cache-service/config"
	pb "github.com/thetommytwitch/simple-cache-service/proto/go"
	"google.golang.org/grpc"
)

const (
	port = ":7771"
)

var server pb.CacheServer

func init() {
	c, err := config.Get()
	if err != nil {
		panic(err)
	}

	switch c.Backend {
	case "memory":
		server = memory.NewServer()
	default:
		panic("backend not set there is a problem with the config")
	}
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	log.Printf("listening on: %s", port)
	pb.RegisterCacheServer(grpcServer, server)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
