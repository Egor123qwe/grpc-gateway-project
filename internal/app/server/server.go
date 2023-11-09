package server

import (
	"github.com/Egor123qwe/grpc-gateway-project/internal/app/server/grpcGwServer"
	"github.com/Egor123qwe/grpc-gateway-project/internal/app/server/grpcServer"
	"golang.org/x/exp/slog"
	"golang.org/x/sync/errgroup"

	"context"
)

func Start() error {
	ctx := context.Background()
	gr, ctx := errgroup.WithContext(ctx)

	gr.Go(func() error {
		return grpcGwServer.Start(ctx)
	})
	gr.Go(func() error {
		return grpcServer.Start(ctx)
	})

	if err := gr.Wait(); err != nil {
		slog.Error(err.Error())
	}

	return nil
}
