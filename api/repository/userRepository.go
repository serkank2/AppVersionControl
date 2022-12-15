package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type UserRepository struct {
	Collection *mongo.Collection
}

func (collection *UserRepository) RepoRegister() error {
	var c = collection.Collection
	ctx, _ := context.WithTimeout(context.Background(), 20*time.Second)
	_, err := c.InsertOne(ctx, bson.D{{"name", "test"}})
	return err
}

func NewUserRepository(collection *mongo.Collection) *UserRepository {
	return &UserRepository{Collection: collection}
}
