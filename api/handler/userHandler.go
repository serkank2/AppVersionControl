package handler

import (
	"appVersionControl/api/models"
	"appVersionControl/api/services"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

type UserHandler struct {
	UserService *services.UserService
}

func (handler *UserHandler) HandleRegister(c *fiber.Ctx) error {
	var h = handler.UserService
	var registerModel models.RegisterModel
	err := c.BodyParser(&registerModel)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message":    "Body parse error",
			"httpStatus": http.StatusBadRequest,
			"data":       nil,
		})
	}
	//-----------------SerRegister----------------------
	dto, err := h.SerRegister(registerModel)
	//-----------------SerRegister----------------------
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message":    "User addition failed",
			"httpStatus": http.StatusBadRequest,
			"data":       nil,
		})
	}
	return c.Status(http.StatusCreated).JSON(fiber.Map{
		"message":    "Register success",
		"httpStatus": http.StatusCreated,
		"data":       dto,
	})
}

func NewUserHandler(userService *services.UserService) *UserHandler {
	return &UserHandler{UserService: userService}
}
