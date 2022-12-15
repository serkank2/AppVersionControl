package api

import (
	"appVersionControl/api/config"
	"appVersionControl/api/handler"
	"appVersionControl/api/loader"
	"appVersionControl/api/repository"
	"appVersionControl/api/services"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
)

func ApiStart() {
	loader.Loader()
	var DB *mongo.Client = config.ConnectDB(loader.GetMongoUri())
	collection := config.GetCollection(DB, "users")
	app := fiber.New()

	//-----------------User-----------------
	UserRepo := repository.NewUserRepository(collection)
	UserService := services.NewUserService(UserRepo)
	UserHandler := handler.NewUserHandler(UserService)

	app.Post("/register", UserHandler.HandleRegister)
	//-----------------User-----------------
	app.Listen(":" + loader.GetPort())

}
