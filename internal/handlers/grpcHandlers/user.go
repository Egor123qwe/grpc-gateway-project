package grpcHandlers

import (
	"github.com/Egor123qwe/grpc-gateway-project/internal/models"
	"github.com/Egor123qwe/grpc-gateway-project/internal/scenarios/serviceInterfaces"
	"github.com/Egor123qwe/grpc-gateway-project/proto/api/generate/desc"
	"google.golang.org/protobuf/types/known/emptypb"

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

func (h *Handlers) CreateUser(ctx context.Context, usr *desc.UserData) (*desc.UserAccessInfo, error) {
	newUser, err := h.scenarios.CreateUser(
		ctx,
		&models.User{
			Email: usr.GetEmail(),
			Name:  usr.GetName(),
			Age:   usr.GetAge(),
		},
	)
	if err != nil {
		return nil, err
	}

	return &desc.UserAccessInfo{
		Token: newUser.Token,
		Id:    newUser.Id.Hex(),
	}, nil
}

func (h *Handlers) GetUser(context.Context, *desc.UserRequest) (*desc.User, error) {
	return nil, nil
}

func (h *Handlers) DeleteUser(context.Context, *emptypb.Empty) (*emptypb.Empty, error) {
	return nil, nil
}

func (h *Handlers) EditUser(context.Context, *desc.UserData) (*emptypb.Empty, error) {
	return nil, nil
}
