package main

import (
	"fmt"
	"log"
	"net"
	"time"

	"github.com/DylanGraham/run-grpc/pb"
	"google.golang.org/grpc"
)

type runServer struct {
	pb.UnimplementedRunServer
}

func (r *runServer) Chat(stream pb.Run_ChatServer) error {
	resp := pb.ChatNote{
		Msg: "Hello",
	}

	for {
		err := stream.Send(&resp)
		if err != nil {
			return err
		}
		time.Sleep(5 * time.Second)
	}
}

func main() {
	const port = ":8081"
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
