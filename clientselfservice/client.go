package main

import (
	"context"
	"log"

	pb "github.com/cohen2804/selfservice/generated"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewSessionRecordsServiceClient(conn)

	req := &pb.ListActionsRq{}
	response, err := client.ListActions(context.Background(), req)
	if err != nil {
		log.Fatalf("Failed to call SayHello: %v", err)
	}

	log.Printf("Response from server: %v", response.GetRecords())
}

//go run ./client/client.go
