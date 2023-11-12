package repositories

import (
	"context"
	"github.com/Egor123qwe/grpc-gateway-project/internal/models"
	"github.com/Egor123qwe/grpc-gateway-project/internal/storage/repsInterfaces"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Users struct {
	db *mongo.Collection
}

func NewUserRep(db *mongo.Collection) repsInterfaces.User {
	return &Users{db: db}
}

func (s *Users) Create(ctx context.Context, usr *models.User) (*models.User, error) {
	usr.Id = primitive.NewObjectID()
	id, err := s.db.InsertOne(ctx, usr)
	if err != nil {
		return nil, err
	}
	usr.Id = id.InsertedID.(primitive.ObjectID)
	return usr, nil
}
