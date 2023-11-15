package grpcHandlers

import (
	"github.com/Egor123qwe/grpc-gateway-project/internal/models"
	"github.com/Egor123qwe/grpc-gateway-project/proto/api/generate/desc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"

	"context"
)

func (h *Handlers) SubscribeUser(ctx context.Context, usr *desc.UserRequest) (*emptypb.Empty, error) {
	if err := h.scenarios.SubscribeUser(ctx, &models.SubscribeEvent{
		SubscriberId: "655325e8f6344ad8f0d9119e",
		ListenerId:   usr.GetId(),
	}); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return &emptypb.Empty{}, nil
}

func (h *Handlers) UnsubscribeUser(ctx context.Context, usr *desc.UserRequest) (*emptypb.Empty, error) {
	if err := h.scenarios.UnsubscribeUser(ctx, &models.SubscribeEvent{
		SubscriberId: "655325e8f6344ad8f0d9119e",
		ListenerId:   usr.GetId(),
	}); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return &emptypb.Empty{}, nil
}
