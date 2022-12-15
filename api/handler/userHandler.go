package handler

import (
	"appVersionControl/api/models"
	"appVersionControl/api/services"
	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	UserService *services.UserService
}

func (handler *UserHandler) HandleRegister(c *fiber.Ctx) error {
	var h = handler.UserService
	var registerModel models.RegisterModel
	err := c.BodyParser(&registerModel)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Bad Request",
		})
	}
	err = h.SerRegister()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Register error",
		})
	}
	return c.SendString("Register")
}

func NewUserHandler(userService *services.UserService) *UserHandler {
	return &UserHandler{UserService: userService}
}
