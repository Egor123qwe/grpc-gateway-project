package scenarios

import (
	"context"
	"errors"
	"github.com/Egor123qwe/grpc-gateway-project/internal/models"
	"github.com/Egor123qwe/grpc-gateway-project/internal/scenarios/serviceInterfaces"
	"github.com/Egor123qwe/grpc-gateway-project/internal/services"
	"github.com/Egor123qwe/grpc-gateway-project/internal/storage"
)

var userNotFoundErr = errors.New("user not found")

type Scenarios struct {
	storage storage.Storage
}

func New(storage storage.Storage) serviceInterfaces.User {
	return &Scenarios{
		storage: storage,
	}
}

func (s *Scenarios) CreateUser(ctx context.Context, usr *models.User) (*models.User, error) {
	token, err := services.CreateToken(usr.Email)
	if err != nil {
		return nil, err
	}
	usr.Token = token
	usr.Subscribers = []string{}
	usr.Subscriptions = []string{}

	newUser, err := s.storage.User().Create(ctx, usr)
	if err != nil {
		return nil, err
	}
	return newUser, nil
}

func (s *Scenarios) GetUser(ctx context.Context, id string) (*models.User, error) {
	usr, err := s.storage.User().Get(ctx, id)
	if err != nil {
		return nil, userNotFoundErr
	}

	return usr, nil
}

func (s *Scenarios) DeleteUser(ctx context.Context, id string) error {
	if err := s.storage.User().Delete(ctx, id); err != nil {
		return userNotFoundErr
	}

	return nil
}

func (s *Scenarios) SubscribeUser(ctx context.Context, ids *models.SubscribeEvent) error {
	return s.storage.User().AddSubscribeEvent(ctx, ids)
}

func (s *Scenarios) UnsubscribeUser(ctx context.Context, ids *models.SubscribeEvent) error {
	return s.storage.User().StealSubscribeEvent(ctx, ids)
}
