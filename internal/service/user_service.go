package service

import (
	"context"
	"fmt"
	"training-project/internal/dto"
	"training-project/internal/model"
	"training-project/internal/repository"
	"training-project/pkg/util"
)

type userService struct {
	userRepository repository.UserRepository
}

type UserService interface {
	IsUserExistedById(ctx context.Context, id int64) (bool, error)
	FindAllUsers(ctx context.Context) ([]model.User, error)
	FindUserById(ctx context.Context, id int64) (*model.User, error)
	FindUserByUsername(ctx context.Context, username string) (*model.User, error)
	CreateUser(ctx context.Context, newUser *dto.UserCreateReqDTO) error
	UpdateUserById(ctx context.Context, id int64, updatedUser *dto.UserUpdateReqDTO) error
	DeleteUserById(ctx context.Context, id int64) error
}

func NewUserService(userRepository repository.UserRepository) UserService {
	return &userService{userRepository: userRepository}
}

func (userService *userService) IsUserExistedById(ctx context.Context, id int64) (bool, error) {
	existed, err := userService.userRepository.ExistsById(ctx, id)
	if err != nil {
		return false, err
	}
	return existed, nil
}

func (userService *userService) FindAllUsers(ctx context.Context) ([]model.User, error) {
	return userService.userRepository.FindAll(ctx)
}

func (userService *userService) FindUserById(ctx context.Context, id int64) (*model.User, error) {
	return userService.userRepository.FindById(ctx, id)
}

func (userService *userService) FindUserByUsername(ctx context.Context, username string) (*model.User, error) {
	return userService.userRepository.FindByUsername(ctx, username)
}

func (userService *userService) CreateUser(ctx context.Context, userCreateReqDTO *dto.UserCreateReqDTO) error {
	newUser := model.User{
		Name:     userCreateReqDTO.Name,
		Email:    userCreateReqDTO.Email,
		Username: userCreateReqDTO.Username,
		Password: userCreateReqDTO.Password,
		Address:  userCreateReqDTO.Address,
		Role:     "ROLE_CUSTOMER",
	}
	return userService.userRepository.Create(ctx, &newUser)
}

func (userService *userService) UpdateUserById(ctx context.Context, id int64, userUpdateReqDTO *dto.UserUpdateReqDTO) error {
	foundUser, err := userService.FindUserById(ctx, id)
	if err != nil {
		return err
	}
	util.ApplyUserUpdate(foundUser, userUpdateReqDTO)
	return userService.userRepository.UpdateById(ctx, id, foundUser)
}

func (userService *userService) DeleteUserById(ctx context.Context, id int64) error {
	existed, err := userService.IsUserExistedById(ctx, id)
	if err != nil {
		return err
	}
	if !existed {
		return fmt.Errorf("id of user is not valid")
	}
	return userService.userRepository.DeleteById(ctx, id)
}
