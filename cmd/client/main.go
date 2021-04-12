package main

import (
	"context"
	"fmt"
	"log"

	"github.com/DylanGraham/run-grpc/pb"
	"google.golang.org/grpc"
)

type grpcContext struct {
	conn   *grpc.ClientConn
	client pb.RunClient
}

func main() {
	ctx := context.Background()
	// ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	// defer cancel()

	gc := grpcContext{}
	var err error
	gc.conn, err = grpc.Dial(":8080", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect to backend: %v", err)
	}
	defer gc.conn.Close()

	gc.client = pb.NewRunClient(gc.conn)

	msg := pb.HelloRequest{
		Name: "Dylan",
	}

	resp, err := gc.client.Hello(ctx, &msg)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Response: %v\n", resp.Msg)
}
