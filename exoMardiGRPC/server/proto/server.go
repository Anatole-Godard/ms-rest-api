package proto

import (
	"context"
	"fmt"
)

type Server struct {
	UnimplementedTutorialServer
}

func (s *Server) SayHello(ctx context.Context, in *HelloRequest) (*HelloReply, error) {
	fmt.Printf("Received message:", in.Name)
	return &HelloReply{Message: "Hello, " + in.GetName()}, nil
}
