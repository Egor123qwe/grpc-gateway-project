package grpcGwServer

import (
	"context"
	quotepb "github.com/Egor123qwe/grpc-gateway-project/proto/api/generate/desc"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"golang.org/x/exp/slog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net/http"
)

func Start(ctx context.Context) error {
	conn, err := grpc.DialContext(
		context.Background(),
		"localhost:8080",
		grpc.WithBlock(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalln(err)
	}

	gwMux := runtime.NewServeMux()
	err = quotepb.RegisterQuoteServiceHandler(context.Background(), gwMux, conn)
	if err != nil {
		log.Fatalln("Failed to register gateway : ", err)
	}
	gwServer := &http.Server{
		Addr:    ":8000",
		Handler: gwMux,
	}
	log.Println("serving grpc-gateway on http://localhost:8000")

	gwServer.RegisterOnShutdown(func() {
		err = conn.Close()
		slog.Error(err.Error())
	})
	errCh := make(chan error)
	go func() {
		defer close(errCh)
		errCh <- gwServer.ListenAndServe()
	}()
	select {
	case err = <-errCh:
	case <-ctx.Done():
		err = gwServer.Shutdown(ctx)
		slog.Error(err.Error())
	}
	return err
}
