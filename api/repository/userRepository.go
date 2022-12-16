package repository

import (
	"appVersionControl/api/dto"
	"appVersionControl/api/models"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
	"time"
)

type UserRepository struct {
	Collection *mongo.Collection
}

func (collection *UserRepository) RepoRegister(registerModel models.RegisterModel) dto.RegisterDto {
	var c = collection.Collection
	ctx, _ := context.WithTimeout(context.Background(), 20*time.Second)
	id, err := c.InsertOne(ctx, registerModel)
	if err != nil {
		return dto.RegisterDto{
			Error:          "User addition failed",
			HttpStatusCode: http.StatusBadRequest,
			Result:         nil,
		}
	}
	return dto.RegisterDto{
		Error:          "",
		HttpStatusCode: http.StatusCreated,
		Result: &dto.RegisterDataDto{
			Id:       id.InsertedID.(primitive.ObjectID).Hex(),
			Name:     registerModel.Name,
			Surname:  registerModel.Surname,
			Username: registerModel.Username,
			Email:    registerModel.Email,
		},
	}
}

func (collection *UserRepository) RepoUserCheck(registerModel models.RegisterModel) bool {
	var c = collection.Collection
	ctx, _ := context.WithTimeout(context.Background(), 20*time.Second)
	filter := bson.D{
		{"$or", bson.A{
			bson.D{{"username", registerModel.Username}},
			bson.D{{"email", registerModel.Email}},
		}},
	}
	err := c.FindOne(ctx, filter)
	if err.Err() == mongo.ErrNoDocuments {
		return false
		// user not found --> return false
	}
	return true
	// user find success --> return true
}

func (collection *UserRepository) RepoLogin(loginModel models.LoginModel) dto.LoginDto {
	var c = collection.Collection
	ctx, _ := context.WithTimeout(context.Background(), 20*time.Second)
	filter := bson.D{
		{"$or", bson.A{
			bson.D{{"username", loginModel.EmailOrUsername}},
			bson.D{{"email", loginModel.EmailOrUsername}},
		}},
	}

	var LoginDbResult dto.LoginDbResult
	err := c.FindOne(ctx, filter).Decode(&LoginDbResult)
	if err != nil {
		return dto.LoginDto{
			Error:          "User not found",
			HttpStatusCode: http.StatusBadRequest,
			Result:         nil,
		}
	}
	if LoginDbResult.Password != loginModel.Password {
		return dto.LoginDto{
			Error:          "Password is incorrect",
			HttpStatusCode: http.StatusBadRequest,
			Result:         nil,
		}
	}
	return dto.LoginDto{
		Error:          "",
		HttpStatusCode: http.StatusOK,
		Result: &dto.LoginDataDto{
			Id:       LoginDbResult.Id.Hex(),
			Email:    LoginDbResult.Email,
			Username: LoginDbResult.Username,
		},
	}
}

func NewUserRepository(collection *mongo.Collection) *UserRepository {
	return &UserRepository{Collection: collection}
}
