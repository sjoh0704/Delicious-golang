package main

import (
    "context"
    "log"

    "google.golang.org/grpc"
    "grpc-test-client/chat"
)

const (
    address     = "localhost:9000"
)

func main() {
    conn, err := grpc.Dial(address, grpc.WithInsecure())
    if err != nil {
        log.Fatalf("did not connect: %v", err)
    }
    defer conn.Close()
	c := chat.NewGreeterClient(conn)

    name := "World"
    r, err := c.SayHello(context.Background(), &chat.HelloRequest{Name: name})
    if err != nil {
        log.Fatalf("could not greet: %v", err)
    }
    log.Printf("Greeting: %s", r.Message)
}