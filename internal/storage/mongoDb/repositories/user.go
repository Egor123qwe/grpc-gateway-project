package repositories

import (
	"github.com/Egor123qwe/grpc-gateway-project/internal/storage/repsInterfaces"
	"go.mongodb.org/mongo-driver/mongo"
)

type Users struct {
	db *mongo.Collection
}

func NewUserRep(db *mongo.Collection) repsInterfaces.User {
	return &Users{db: db}
}
