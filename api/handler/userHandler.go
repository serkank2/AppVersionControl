package handler

import (
	"appVersionControl/api/dto"
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
		return c.Status(http.StatusBadRequest).JSON(dto.RegisterDto{
			HttpStatusCode: http.StatusBadRequest,
			Error:          "body parse error",
			Result:         nil,
		})
	}
	//-----------------SerRegister----------------------
	dto := h.SerRegister(registerModel)
	//-----------------SerRegister----------------------
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message":    "User addition failed",
			"httpStatus": http.StatusBadRequest,
			"data":       nil,
		})
	}
	return c.Status(dto.HttpStatusCode).JSON(dto)
}
func (handler *UserHandler) HandleLogin(c *fiber.Ctx) error {
	var h = handler.UserService
	var loginModel models.LoginModel
	err := c.BodyParser(&loginModel)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(dto.RegisterDto{
			HttpStatusCode: http.StatusBadRequest,
			Error:          "body parse error",
			Result:         nil,
		})
	}
	dto := h.SerLogin(loginModel)
	return c.Status(dto.HttpStatusCode).JSON(dto)
}
func NewUserHandler(userService *services.UserService) *UserHandler {
	return &UserHandler{UserService: userService}
}
