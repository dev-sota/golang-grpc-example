package main

import (
	"context"
	"fmt"
	"os"

	"github.com/dev-sota/golang-grpc-example/pinger"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial(":5300", grpc.WithInsecure())
	if err != nil {
		fmt.Fprintf(os.Stderr, "grpc.Dial: %v\n", err)
		return
	}
	defer conn.Close()

	client := pinger.NewPingerClient(conn)
	req := &pinger.Request{
		Message: "ping",
	}
	res, err := client.Ping(context.Background(), req)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Ping: %v\n", err)
		return
	}
	fmt.Fprintf(os.Stdout, "Response: %s\n", res.Message)
}
