package repository

import (
	"context"
	"training-project/internal/model"

	"github.com/uptrace/bun"
)

type userRepository struct {
	db *bun.DB
}

type UserRepository interface {
	GetAllUsers(ctx context.Context) ([]model.User, error)
}

func NewUserRepository(db *bun.DB) UserRepository {
	return &userRepository{db: db}
}

func (userRepo *userRepository) GetAllUsers(ctx context.Context) ([]model.User, error) {
	users := []model.User{}
	err := userRepo.db.NewSelect().Model(&users).Scan(ctx)
	if err != nil {
		return nil, err
	}
	return users, nil
}
