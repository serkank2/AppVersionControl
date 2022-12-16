package dto

import "go.mongodb.org/mongo-driver/bson/primitive"

type LoginDto struct {
	Error          string        `json:"error"`
	HttpStatusCode int           `json:"httpStatus"`
	Result         *LoginDataDto `json:"result"`
}

type LoginDataDto struct {
	Id       string `json:"id"`
	Email    string `json:"email"`
	Username string `json:"username"`
}

type LoginDbResult struct {
	Id       primitive.ObjectID `bson:"_id"`
	Email    string             `bson:"email"`
	Username string             `bson:"username"`
	Password string             `bson:"password"`
}
