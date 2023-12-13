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

	client := pb.NewGreetingServiceClient(conn) // pb.NewGreetServiceClient(conn)
	req := &pb.HelloMessage{Title: "Ran", Body: "Cohen"}

	response, err := client.SayHello(context.Background(), req)
	if err != nil {
		log.Fatalf("Failed to call SayHello: %v", err)
	}

	log.Printf("Response from server: %s %s", response.Title, response.Body)
}

//go run ./client/client.go
