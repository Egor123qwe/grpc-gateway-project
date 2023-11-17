package serviceInterfaces

import (
	"context"
	"github.com/Egor123qwe/grpc-gateway-project/internal/models"
)

type User interface {
	CreateUser(ctx context.Context, usr *models.User) (*models.User, error)
	GetUser(ctx context.Context, id string) (*models.User, error)
	DeleteUser(ctx context.Context, id string) error
	SubscribeUser(ctx context.Context, ids *models.SubscribeEvent) error
	UnsubscribeUser(ctx context.Context, ids *models.SubscribeEvent) error
	GetUserByToken(ctx context.Context, token string) (*models.User, error)
}
