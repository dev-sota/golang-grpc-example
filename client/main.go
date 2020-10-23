package main

import (
	"context"
	"log"

	pb "github.com/dev-sota/golang-grpc-example/helloworld"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

func main() {
	conn, err := grpc.Dial(port, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn)

	res, err := c.SayHello(context.Background(), &pb.HelloRequest{Name: "dev-sota"})
	if err != nil {
		log.Fatalf("RPC error: %v", err)
	}
	log.Printf("Greeting: %s", res.Message)
}
