package mongoDb

import (
	"fmt"
	"github.com/Egor123qwe/grpc-gateway-project/internal/storage"
	"github.com/Egor123qwe/grpc-gateway-project/internal/storage/mongoDb/repositories"
	"github.com/Egor123qwe/grpc-gateway-project/internal/storage/repsInterfaces"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"context"
)

type MongoDb struct {
	db   *mongo.Collection
	user repsInterfaces.User
}

func New(database *mongo.Database) storage.Storage {
	return configure(database.Collection(NewConfig().MongoCollection))
}

func NewClient(ctx context.Context) (db *mongo.Database, err error) {
	var mongoDBURL string
	config := NewConfig()

	mongoDBURL = fmt.Sprintf("mongodb://%s:%s", config.MongoHost, config.MongoPort)
	clientOptions := options.Client().ApplyURI(mongoDBURL)

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, err
	}
	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, err
	}
	return client.Database(config.MongoDBName), nil
}

func configure(db *mongo.Collection) storage.Storage {
	return &MongoDb{
		db:   db,
		user: repositories.NewUserRep(db),
	}
}

func (s *MongoDb) User() repsInterfaces.User { return s.user }
