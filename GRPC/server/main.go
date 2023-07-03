package main

import (
	"log"
	"net"

	pb "github.com/ankit/GRPC/proto"
	"google.golang.org/grpc"
)

const (
	port = ":8080"
)

type helloServer struct {
	pb.RegisterGreetServiceServer
}

func main() {
	lis, err := net.Listen("tcp", port)

	if err != nil {
		log.Fatalf("Failed to start the server %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterGreetServiceServer(grpcServer, &helloServer{})
	if err := grpcServer.serve(lis); err != nil {
		log.Fatalf("failled to start: %v", err)
	}
}
