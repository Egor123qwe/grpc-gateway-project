package grpcHandlers

import (
	"github.com/Egor123qwe/grpc-gateway-project/proto/api/generate/desc"
	"google.golang.org/protobuf/types/known/emptypb"

	"context"
)

func (h *Handlers) SubscribeUser(context.Context, *desc.UserRequest) (*emptypb.Empty, error) {
	return nil, nil
}

func (h *Handlers) UnsubscribeUser(context.Context, *desc.UserRequest) (*emptypb.Empty, error) {
	return nil, nil
}
