package services

import (
	"appVersionControl/api/repository"
	"fmt"
)

type UserService struct {
	UserRepository *repository.UserRepository
}

func (service *UserService) SerRegister() error {
	var s = service.UserRepository
	err := s.RepoRegister()
	if err != nil {
		fmt.Printf("Register error %v", err)
	}
	return err
}

func NewUserService(userRepository *repository.UserRepository) *UserService {
	return &UserService{UserRepository: userRepository}
}
