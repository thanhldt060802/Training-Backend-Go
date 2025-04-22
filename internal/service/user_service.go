package service

import (
	"context"
	"training-project/internal/model"
	"training-project/internal/repository"
)

type userService struct {
	userRepo repository.UserRepository
}

type UserService interface {
	GetAllUsers(ctx context.Context) ([]model.User, error)
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{userRepo: userRepo}
}

func (userServ *userService) GetAllUsers(ctx context.Context) ([]model.User, error) {
	return userServ.userRepo.GetAllUsers(ctx)
}
