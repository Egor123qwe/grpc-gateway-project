package grpcServer

import (
	"context"
	"fmt"
	"github.com/Egor123qwe/grpc-gateway-project/internal/config"
	"github.com/Egor123qwe/grpc-gateway-project/proto/api/generate/desc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

func Start(ctx context.Context, service desc.UserServiceServer, config *config.Config) error {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", config.GrpcPort))
	if err != nil {
		log.Fatalln("Failed to listen grpc server: ", err)
	}

	s := grpc.NewServer()
	reflection.Register(s)
	desc.RegisterUserServiceServer(s, service)
	log.Printf("serving gRPC on http://localhost:%d\n", config.GrpcPort)

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
