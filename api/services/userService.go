package services

import (
	"appVersionControl/api/dto"
	"appVersionControl/api/models"
	"appVersionControl/api/repository"
	"fmt"
)

type UserService struct {
	UserRepository *repository.UserRepository
}

func (service *UserService) SerRegister(registerModel models.RegisterModel) (dto.RegisterDto, error) {
	var s = service.UserRepository
	//-----------------RepoRegister----------------------
	dto, err := s.RepoRegister(registerModel)
	//-----------------RepoRegister----------------------
	if err != nil {
		fmt.Printf("Register error %v", err)
	}
	return dto, err
}

func NewUserService(userRepository *repository.UserRepository) *UserService {
	return &UserService{UserRepository: userRepository}
}
