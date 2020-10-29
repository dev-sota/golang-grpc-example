package main

import (
	"context"
	"log"
	"net"

	"github.com/dev-sota/golang-grpc-example/pinger"
	"google.golang.org/grpc"
)

type server struct{}

func (s *server) Ping(ctx context.Context, req *pinger.Request) (*pinger.Pong, error) {
	res := &pinger.Pong{
		Message: "pong",
	}
	log.Printf("New Request: %v", req.Message)
	return res, nil
}

func main() {
	listener, err := net.Listen("tcp", ":5300")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcSrv := grpc.NewServer()
	pinger.RegisterPingerServer(grpcSrv, &server{})
	log.Printf("Pinger service is running")
	err = grpcSrv.Serve(listener)
	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
