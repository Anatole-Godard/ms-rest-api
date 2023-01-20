package main

import (
	"log"
	"main/proto"
	"net"

	"google.golang.org/grpc"
)

func main() {
	println("gRPC server proto in Go")

	listener, err := net.Listen("tcp", ":9000")
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()
	proto.RegisterTutorialServer(s, &proto.Server{})
	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
