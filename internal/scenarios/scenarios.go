package scenarios

import (
	"github.com/Egor123qwe/grpc-gateway-project/internal/models"
	"github.com/Egor123qwe/grpc-gateway-project/internal/scenarios/serviceInterfaces"
	"github.com/Egor123qwe/grpc-gateway-project/internal/services"
	"github.com/Egor123qwe/grpc-gateway-project/internal/storage"
	"github.com/Egor123qwe/grpc-gateway-project/proto/api/generate/desc"

	"context"
)

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

	newUser, err := s.storage.User().Create(ctx, usr)
	if err != nil {
		return nil, err
	}

	return newUser, nil
}

func (s *Scenarios) GetUser(id string) (*models.User, error) {
	return nil, nil
}

func (s *Scenarios) DeleteUser(id string) error {
	return nil
}

func (s *Scenarios) EditUser(usr *desc.UserData) error {
	return nil
}

func (s *Scenarios) SubscribeUser(subscriberId string, userId string) error {
	return nil
}

func (s *Scenarios) UnsubscribeUser(subscriberId string, userId string) error {
	return nil
}
