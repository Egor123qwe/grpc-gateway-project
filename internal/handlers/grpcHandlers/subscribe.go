package grpcHandlers

import (
	"github.com/Egor123qwe/grpc-gateway-project/proto/api/generate/desc"

	"context"
)

func (h *Handlers) SubscribeUser(context.Context, *desc.UserRequest) (*desc.Empty, error) {
	return nil, nil
}

func (h *Handlers) UnsubscribeUser(context.Context, *desc.UserRequest) (*desc.Empty, error) {
	return nil, nil
}
