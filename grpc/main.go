package main

import (
	"context"
	"grpc-test/chat"
	"log"
	"net"

	"google.golang.org/grpc"
)

type server struct{
	chat.UnimplementedGreeterServer
}

func (s *server) SayHello(ctx context.Context, in *chat.HelloRequest) (*chat.HelloReply, error) {
    log.Printf("Received: %v", in.Name)
    return &chat.HelloReply{Message: "Hello " + in.Name}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	chat.RegisterGreeterServer(grpcServer, &server{})

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

