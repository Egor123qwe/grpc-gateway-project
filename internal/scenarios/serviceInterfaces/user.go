package serviceInterfaces

import (
	"github.com/Egor123qwe/grpc-gateway-project/internal/models"
	"github.com/Egor123qwe/grpc-gateway-project/proto/api/generate/desc"

	"context"
)

type User interface {
	CreateUser(ctx context.Context, usr *models.User) (*models.User, error)
	GetUser(id string) (*models.User, error)
	DeleteUser(id string) error
	EditUser(usr *desc.UserData) error
	SubscribeUser(subscriberId string, userId string) error
	UnsubscribeUser(subscriberId string, userId string) error
}
