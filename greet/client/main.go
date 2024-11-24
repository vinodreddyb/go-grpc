package main

import (
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/vinodreddyb/go-grpc/greet/proto"
)

var add string = "localhost:50051"

func main() {
	conn, err := grpc.NewClient(add, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	defer conn.Close()

	client := pb.NewGreetServiceClient(conn)
	doGreet(client)
}
