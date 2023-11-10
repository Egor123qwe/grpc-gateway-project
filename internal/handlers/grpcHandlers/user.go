package grpcHandlers

import (
	"github.com/Egor123qwe/grpc-gateway-project/internal/scenarios/serviceInterfaces"
	"github.com/Egor123qwe/grpc-gateway-project/proto/api/generate/desc"

	"context"
)

type Handlers struct {
	scenarios serviceInterfaces.User
	desc.UnimplementedUserServiceServer
}

func New(scenarios serviceInterfaces.User) *Handlers {
	return &Handlers{
		scenarios: scenarios,
	}
}

func (h *Handlers) CreateUser(context.Context, *desc.Empty) (*desc.Empty, error) {
	return nil, nil
}

func (h *Handlers) GetUser(context.Context, *desc.UserRequest) (*desc.Empty, error) {
	return nil, nil
}

func (h *Handlers) DeleteUser(context.Context, *desc.UserRequest) (*desc.Empty, error) {
	return nil, nil
}

func (h *Handlers) EditUser(context.Context, *desc.UserRequest) (*desc.Empty, error) {
	return nil, nil
}
