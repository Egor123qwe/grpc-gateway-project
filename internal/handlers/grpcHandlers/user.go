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
		Id:    newUser.Id,
	}, nil
}

func (h *Handlers) GetUser(ctx context.Context, usr *desc.UserRequest) (*desc.User, error) {
	user, err := h.scenarios.GetUser(ctx, usr.GetId())
	if err != nil {
		return nil, err
	}

	return &desc.User{
		Id: user.Id,
		User: &desc.UserData{
			Email: user.Email,
			Name:  user.Name,
			Age:   user.Age,
		},
	}, nil
}

func (h *Handlers) DeleteUser(ctx context.Context, in *emptypb.Empty) (*emptypb.Empty, error) {
	if err := h.scenarios.DeleteUser(ctx, "65513ee8b2600562e33e9b2b"); err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}
