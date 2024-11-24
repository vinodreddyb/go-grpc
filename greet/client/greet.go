package main

import (
	"context"
	"log"

	pb "github.com/vinodreddyb/go-grpc/greet/proto"
)

func doGreet(c pb.GreetServiceClient) {
	// Create a request
	req := &pb.GreetRequest{
		FirstName: "Vinod",
	}

	// Call the service
	res, err := c.Greet(context.Background(), req)
	if err != nil {
		log.Fatalf("Failed to call Greet: %v", err)
	}

	// Print the response
	log.Printf("Response: %s", res.Result)
}
