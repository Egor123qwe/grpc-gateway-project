package grpcServer

import (
	"context"
	"github.com/Egor123qwe/grpc-gateway-project/proto/api/generate/desc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

type server struct {
	desc.UnimplementedQuoteServiceServer
}

func Start(ctx context.Context) error {
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln("Failed to listen grpc server: ", err)
	}
	s := grpc.NewServer()
	reflection.Register(s)
	desc.RegisterQuoteServiceServer(s, &server{})
	log.Println("serving gRPC on localhost:8080")

	errCh := make(chan error)
	go func() {
		defer close(errCh)
		errCh <- s.Serve(lis)
	}()
	select {
	case err = <-errCh:
		return err
	case <-ctx.Done():
		s.GracefulStop()
	}
	return nil
}
