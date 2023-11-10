package server

import (
	"github.com/Egor123qwe/grpc-gateway-project/internal/app/server/grpcGwServer"
	"github.com/Egor123qwe/grpc-gateway-project/internal/app/server/grpcServer"
	"github.com/Egor123qwe/grpc-gateway-project/internal/config"
	"github.com/Egor123qwe/grpc-gateway-project/internal/scenarios"
	"github.com/Egor123qwe/grpc-gateway-project/proto/api/generate/desc"
	"golang.org/x/exp/slog"
	"golang.org/x/sync/errgroup"

	"context"
)

type server struct {
	config    *config.Config
	scenarios scenarios.Scenarios
}

func New() *server {
	return &server{
		config: config.New(),
	}
}

func (s *server) Start() error {
	ctx := context.Background()
	gr, ctx := errgroup.WithContext(ctx)

	gr.Go(func() error {
		return grpcGwServer.Start(ctx, s.config)
	})
	gr.Go(func() error {
		return grpcServer.Start(ctx, s.scenarios, s.config)
	})

	if err := gr.Wait(); err != nil {
		slog.Error(err.Error())
	}

	return nil
}

type tmp struct {
	desc.UnimplementedUserServiceServer
}
