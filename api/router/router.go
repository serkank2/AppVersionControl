package router

import "github.com/gofiber/fiber/v2"

type Router struct {
	App *fiber.App
}

func (r *Router) UserRouter() {
	router := r.App
	router.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})
}

func NewRouter(App *fiber.App) *Router {
	return &Router{App: App}

}
