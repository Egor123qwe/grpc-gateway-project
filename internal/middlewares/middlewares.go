package middlewares

import (
	"context"
	"github.com/Egor123qwe/grpc-gateway-project/internal/storage"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"google.golang.org/grpc"
)

type Middlewares struct {
	storage storage.Storage
}

func New(storage storage.Storage) grpc.UnaryServerInterceptor {
	m := &Middlewares{
		storage: storage,
	}
	return configureMainInterceptor(m.configureMethodInterceptors())
}

func (m *Middlewares) configureMethodInterceptors() map[string]grpc.UnaryServerInterceptor {
	return map[string]grpc.UnaryServerInterceptor{
		"/package.Service/MethodA": grpc_middleware.ChainUnaryServer(),
		"/package.Service/MethodB": grpc_middleware.ChainUnaryServer(),
	}
}

func configureMainInterceptor(methodInterceptors map[string]grpc.UnaryServerInterceptor) grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (interface{}, error) {
		if interceptor, ok := methodInterceptors[info.FullMethod]; ok {
			return interceptor(ctx, req, info, handler)
		}
		return handler(ctx, req)
	}
}
