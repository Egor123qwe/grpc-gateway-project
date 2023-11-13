package repositories

import (
	"context"
	"errors"
	"github.com/Egor123qwe/grpc-gateway-project/internal/models"
	"github.com/Egor123qwe/grpc-gateway-project/internal/storage/repsInterfaces"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var userNotFoundErr = errors.New("user not found")

type Users struct {
	db *mongo.Collection
}

func NewUserRep(db *mongo.Collection) repsInterfaces.User {
	return &Users{db: db}
}

func (s *Users) Create(ctx context.Context, usr *models.User) (*models.User, error) {
	id, err := s.db.InsertOne(ctx, usr)
	if err != nil {
		return nil, err
	}
	usr.Id = id.InsertedID.(primitive.ObjectID).Hex()
	return usr, nil
}

func (s *Users) Get(ctx context.Context, id string) (*models.User, error) {
	userId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	var usr models.User
	err = s.db.FindOne(ctx, bson.D{{"_id", userId}}).Decode(&usr)
	if err != nil {
		return nil, err
	}

	return &usr, nil
}

func (s *Users) Delete(ctx context.Context, id string) error {
	userId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	res, err := s.db.DeleteOne(ctx, bson.D{{"_id", userId}})
	if err != nil {
		return err
	}
	if res.DeletedCount == 0 {
		return userNotFoundErr
	}
	return nil
}

func (s *Users) SubscribeUser(ctx context.Context, ids *models.SubscribeEvent) error {
	_, err := primitive.ObjectIDFromHex(ids.UserId)
	if err != nil {
		return err
	}

	//filter := bson.M{"_id": userId}
	//update := bson.M{"$addToSet": bson.M{"subscriptions": ids.SubscriberId}}

	return nil
}
