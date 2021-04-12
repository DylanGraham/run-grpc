package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/DylanGraham/run-grpc/pb"
	"google.golang.org/grpc"
)

type runServer struct {
	pb.UnimplementedRunServer
}

func (r *runServer) Hello(ctx context.Context, hr *pb.HelloRequest) (*pb.HelloResponse, error) {
	resp := pb.HelloResponse{
		Msg: "Hello " + hr.Name,
	}

	return &resp, nil
}

func main() {
	const port = ":8080"
	s := runServer{}
	grpcServer := grpc.NewServer()
	pb.RegisterRunServer(grpcServer, &s)
	l, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("could not listen to %s: %v", port, err)
	}

	fmt.Printf("Listening on %s\n", port)
	grpcServer.Serve(l)
}
