package repository

import (
	"appVersionControl/api/dto"
	"appVersionControl/api/models"
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type UserRepository struct {
	Collection *mongo.Collection
}

func (collection *UserRepository) RepoRegister(registerModel models.RegisterModel) (dto.RegisterDto, error) {
	var c = collection.Collection
	ctx, _ := context.WithTimeout(context.Background(), 20*time.Second)
	id, err := c.InsertOne(ctx, registerModel)

	return dto.RegisterDto{
		Id:       id.InsertedID.(primitive.ObjectID).Hex(),
		Name:     registerModel.Name,
		Surname:  registerModel.Surname,
		Username: registerModel.Username,
		Email:    registerModel.Email,
	}, err
}

func NewUserRepository(collection *mongo.Collection) *UserRepository {
	return &UserRepository{Collection: collection}
}
