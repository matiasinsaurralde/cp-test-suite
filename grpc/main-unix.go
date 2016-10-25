package main

import (
	"log"
	"net"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"github.com/TykTechnologies/tyk/coprocess"
)

const (
	sock = "/tmp/lol.sock"
)

type server struct{}

func (s *server) Dispatch(ctx context.Context, object *coprocess.Object) (*coprocess.Object, error) {
  log.Println("Dispatch")
  object.Request.SetHeaders = make(map[string]string)
  object.Request.SetHeaders["grpcheader"] = "grpcvalue"
  return object, nil
}

func (s *server) DispatchEvent(ctx context.Context, object *coprocess.Event) (*coprocess.EventReply, error) {
  return &coprocess.EventReply{}, nil
}
func main() {
	lis, err := net.Listen("unix", sock)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	coprocess.RegisterDispatcherServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
