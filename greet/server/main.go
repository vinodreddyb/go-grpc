package main

import (
	"context"
	"fmt"
	"log/slog"
	"net"
	"os"
	"os/signal"

	pb "github.com/vinodreddyb/go-grpc/greet/proto"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
)

var add string = "0.0.0.0:50051"

type Server struct {
	pb.GreetServiceServer
}

func main() {

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill)
	defer cancel()

	if err := run(ctx); err != nil {
		slog.Error("Failed to create signal context: %v", err)
		os.Exit(1)
	}

}

func run(ctx context.Context) error {

	g, ctx := errgroup.WithContext(ctx)
	s := grpc.NewServer()
	//Register the server
	pb.RegisterGreetServiceServer(s, &Server{})

	g.Go(func() error {
		lis, err := net.Listen("tcp", add)

		if err != nil {
			//log.Fatalf("Failed to listen: %v", err)
			return fmt.Errorf("failed to listen %q : %v", add, err)
		}
		slog.Info("starting grpc server on address", slog.String("address", add))
		if err = s.Serve(lis); err != nil {
			return fmt.Errorf("failed to serve: %v", err)
		}
		return nil
	})

	g.Go(func() error {
		<-ctx.Done()
		slog.Info("Shutting down the server")
		s.GracefulStop()
		return nil
	})

	return g.Wait()
}
