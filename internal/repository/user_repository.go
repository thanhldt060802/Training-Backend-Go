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
	ExistsById(ctx context.Context, id int64) (bool, error)
	FindAll(ctx context.Context) ([]model.User, error)
	FindById(ctx context.Context, id int64) (*model.User, error)
	FindByUsername(ctx context.Context, username string) (*model.User, error)
	Create(ctx context.Context, newUser *model.User) error
	UpdateById(ctx context.Context, id int64, updatedUser *model.User) error
	DeleteById(ctx context.Context, id int64) error
}

func NewUserRepository(db *bun.DB) UserRepository {
	return &userRepository{db: db}
}

func (userRepository *userRepository) ExistsById(ctx context.Context, id int64) (bool, error) {
	var count int
	count, err := userRepository.db.NewSelect().Model(&model.User{}).Where("id = ?", id).Count(ctx)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func (userRepository *userRepository) FindAll(ctx context.Context) ([]model.User, error) {
	var users []model.User
	err := userRepository.db.NewSelect().Model(&users).Scan(ctx)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (userRepository *userRepository) FindById(ctx context.Context, id int64) (*model.User, error) {
	var user model.User
	err := userRepository.db.NewSelect().Model(&user).Where("id = ?", id).Scan(ctx)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (userRepository *userRepository) FindByUsername(ctx context.Context, username string) (*model.User, error) {
	var user model.User
	err := userRepository.db.NewSelect().Model(&user).Where("username = ?", username).Scan(ctx)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (userRepository *userRepository) Create(ctx context.Context, newUser *model.User) error {
	_, err := userRepository.db.NewInsert().Model(newUser).Exec(ctx)
	return err
}

func (userRepository *userRepository) UpdateById(ctx context.Context, id int64, updatedUser *model.User) error {
	_, err := userRepository.db.NewUpdate().Model(updatedUser).Where("id = ?", id).Exec(ctx)
	return err
}

func (userRepository *userRepository) DeleteById(ctx context.Context, id int64) error {
	_, err := userRepository.db.NewDelete().Model(&model.User{}).Where("id = ?", id).Exec(ctx)
	return err
}
