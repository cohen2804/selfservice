package main

import (
	"context"
	"log"
	"net"

	"github.com/cohen2804/selfservice/generated"
	"google.golang.org/grpc"
)

// should implement the interface myPkgName.InvoicerServer
type myGRPCServer struct {
	generated.UnimplementedGreetingServiceServer
}

func (s *myGRPCServer) SayHello(ctx context.Context, req *generated.HelloMessage) (*generated.HelloMessage, error) {
	myMessage := &generated.HelloMessage{Title: "My title:" + req.Title, Body: "My body:" + req.Body}
	return myMessage, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	opts := []grpc.ServerOption{}
	grpcServer := grpc.NewServer(opts...)
	generated.RegisterGreetingServiceServer(grpcServer, &myGRPCServer{})
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

//go run ./server/server.go
