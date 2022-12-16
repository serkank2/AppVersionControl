package services

import (
	"appVersionControl/api/dto"
	"appVersionControl/api/models"
	"appVersionControl/api/repository"
)

type UserService struct {
	UserRepository *repository.UserRepository
}

func (service *UserService) SerRegister(registerModel models.RegisterModel) dto.RegisterDto {
	var s = service.UserRepository

	// ---------Check if the user exists----------
	check := s.RepoUserCheck(registerModel)
	if check {
		return dto.RegisterDto{
			Error:          "User already exists",
			HttpStatusCode: 400,
			Result:         nil,
		}
	}

	// ---------Check if the user exists----------

	//-----------------RepoRegister----------------------
	dto := s.RepoRegister(registerModel)
	//-----------------RepoRegister----------------------
	return dto
}
func (service *UserService) SerLogin(loginModel models.LoginModel) dto.LoginDto {
	var s = service.UserRepository
	return s.RepoLogin(loginModel)
}

func NewUserService(userRepository *repository.UserRepository) *UserService {
	return &UserService{UserRepository: userRepository}
}
