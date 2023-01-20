package proto

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func CallGRPCServer(msg string) {
	fmt.Println("Calling gRPC server")
	callServer(msg)
}

func callServer(msg string) {
	conn, _ := grpc.Dial("localhost:9000", grpc.WithTransportCredentials(insecure.NewCredentials()))

	// Code removed for brevity

	client := NewTutorialClient(conn)

	// Note how we are calling the SayHello method on the server
	// This is available to us through the auto-generated code
	_, _ = client.SayHello(context.Background(), &HelloRequest{Name: msg})

}
