package repsInterfaces

import (
	"context"
	"github.com/Egor123qwe/grpc-gateway-project/internal/models"
)

type User interface {
	Create(ctx context.Context, usr *models.User) (*models.User, error)
	Get(ctx context.Context, id string) (*models.User, error)
	Delete(ctx context.Context, id string) error
}
