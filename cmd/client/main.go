package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/DylanGraham/run-grpc/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
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
	gc.conn, err = grpc.Dial(":8081", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("could not connect to backend: %v", err)
	}
	defer gc.conn.Close()

	gc.client = pb.NewRunClient(gc.conn)

	stream, err := gc.client.Chat(ctx)
	if err != nil {
		log.Fatal(err)
	}

	for {
		note, err := stream.Recv()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("Response: %s\n", note.Msg)
		time.Sleep(5 * time.Second)
	}
}
