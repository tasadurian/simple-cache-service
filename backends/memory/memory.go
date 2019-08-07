package memory

import (
	"context"
	"errors"
	"sync"

	pb "github.com/thetommytwitch/simple-cache-service/proto/go"
)

// Check that Server implements the CacheServer interface
var _ pb.CacheServer = (*Server)(nil)

// Server in the in memory server type.
type Server struct {
	store sync.Map
}

// NewServer returns a new in memory server.
func NewServer() *Server {
	return &Server{}
}

// Get is the method to retrieve the value at the given key.
func (s *Server) Get(ctx context.Context, in *pb.GetRequest) (*pb.GetReply, error) {
	value, err := s.get(in.Key)
	return &pb.GetReply{Key: in.Key, Value: value, Error: err.Error()}, err
}

func (s *Server) get(key string) ([]byte, error) {
	result, ok := s.store.Load(key)
	if ok {
		return result.([]byte), nil
	}

	return nil, errors.New("key not found")
}

// Set stores the the given value at the given key.
func (s *Server) Set(ctx context.Context, in *pb.SetRequest) (*pb.SetReply, error) {
	err := s.set(in.Key, in.Value)
	return &pb.SetReply{Error: err.Error()}, err
}

func (s *Server) set(key string, value []byte) error {
	s.store.Store(key, value)
	return nil
}

// Del deletes the key and value from the store.
func (s *Server) Del(ctx context.Context, in *pb.DelRequest) (*pb.DelReply, error) {
	err := s.del(in.Key)
	return &pb.DelReply{Error: err.Error()}, err
}

func (s *Server) del(key string) error {
	s.store.Delete(key)
	return nil
}
