package grpcHandlers

import (
	"github.com/Egor123qwe/grpc-gateway-project/internal/models"
	"github.com/Egor123qwe/grpc-gateway-project/proto/api/generate/desc"
	"google.golang.org/protobuf/types/known/emptypb"

	"context"
)

func (h *Handlers) SubscribeUser(ctx context.Context, usr *desc.UserRequest) (*emptypb.Empty, error) {
	if err := h.scenarios.SubscribeUser(ctx, &models.SubscribeEvent{
		SubscriberId: "65525fc2d1c6dce25fcfd9e7",
		ListenerId:   usr.GetId(),
	}); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (h *Handlers) UnsubscribeUser(ctx context.Context, usr *desc.UserRequest) (*emptypb.Empty, error) {
	if err := h.scenarios.UnsubscribeUser(ctx, &models.SubscribeEvent{
		SubscriberId: "65525fc2d1c6dce25fcfd9e7",
		ListenerId:   usr.GetId(),
	}); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
