package main

import (
	"context"
	"log"
	"net"

	pb "github.com/thetommytwitch/simple-cache-service/proto/go"
	"google.golang.org/grpc"
)

const (
	port = ":7771"
)

type server struct{}

func (s *server) Get(ctx context.Context, in *pb.GetRequest) (*pb.GetReply, error) {
	log.Print("Get")
	return nil, nil
}

func (s *server) Set(ctx context.Context, in *pb.SetRequest) (*pb.SetReply, error) {
	log.Print("Set")
	return nil, nil
}

func (s *server) Del(ctx context.Context, in *pb.DelRequest) (*pb.DelReply, error) {
	log.Print("Del")
	return nil, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	log.Printf("listening on: %s", port)
	pb.RegisterCacheServer(grpcServer, &server{})
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
