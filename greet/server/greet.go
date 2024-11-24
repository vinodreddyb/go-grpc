package main

import (
	"context"

	pb "github.com/vinodreddyb/go-grpc/greet/proto"
)

func (s *Server) Greet(ctx context.Context, in *pb.GreetRequest) (*pb.GreetResponse, error) {
	return &pb.GreetResponse{Result: "Hello " + in.FirstName}, nil
}
