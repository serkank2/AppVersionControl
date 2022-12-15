package api

import (
	"appVersionControl/api/loader"
	"appVersionControl/api/router"
	"github.com/gofiber/fiber/v2"
)

func ApiStart() {

	loader.Loader()
	app := fiber.New()

	router.NewRouter(app)

	app.Listen(":" + loader.GetPort())

}
