package repository

import (
	"context"
	"training-project/internal/model"

	"github.com/uptrace/bun"
)

type cartItemRepository struct {
	db *bun.DB
}

type CartItemRepository interface {
	ExistsById(ctx context.Context, id int64) (bool, error)
	FindAll(ctx context.Context) ([]model.CartItem, error)
	FindById(ctx context.Context, id int64) (*model.CartItem, error)
	FindAllByUserId(ctx context.Context, userId int64) ([]model.CartItem, error)
	Create(ctx context.Context, newCartItem *model.CartItem) error
	UpdateById(ctx context.Context, id int64, updatedCartItem *model.CartItem) error
	DeleteById(ctx context.Context, id int64) error
}

func NewCartItemRepository(db *bun.DB) CartItemRepository {
	return &cartItemRepository{db: db}
}

func (cartItemRepository *cartItemRepository) ExistsById(ctx context.Context, id int64) (bool, error) {
	var count int
	count, err := cartItemRepository.db.NewSelect().Model(&model.CartItem{}).Where("id = ?", id).Count(ctx)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func (cartItemRepository *cartItemRepository) FindAll(ctx context.Context) ([]model.CartItem, error) {
	var cartItems []model.CartItem
	err := cartItemRepository.db.NewSelect().Model(&cartItems).Scan(ctx)
	if err != nil {
		return nil, err
	}
	return cartItems, nil
}

func (cartItemRepository *cartItemRepository) FindById(ctx context.Context, id int64) (*model.CartItem, error) {
	var cartItem model.CartItem
	err := cartItemRepository.db.NewSelect().Model(&cartItem).Where("id = ?", id).Scan(ctx)
	if err != nil {
		return nil, err
	}
	return &cartItem, nil
}

func (cartItemRepository *cartItemRepository) FindAllByUserId(ctx context.Context, userId int64) ([]model.CartItem, error) {
	var cartItems []model.CartItem
	err := cartItemRepository.db.NewSelect().Model(&cartItems).Where("user_id = ?", userId).Scan(ctx)
	if err != nil {
		return nil, err
	}
	return cartItems, nil
}

func (cartItemRepository *cartItemRepository) Create(ctx context.Context, newCartItem *model.CartItem) error {
	_, err := cartItemRepository.db.NewInsert().Model(newCartItem).Exec(ctx)
	return err
}

func (cartItemRepository *cartItemRepository) UpdateById(ctx context.Context, id int64, updatedCartItem *model.CartItem) error {
	_, err := cartItemRepository.db.NewUpdate().Model(updatedCartItem).Where("id = ?", id).Exec(ctx)
	return err
}

func (cartItemRepository *cartItemRepository) DeleteById(ctx context.Context, id int64) error {
	_, err := cartItemRepository.db.NewDelete().Model(&model.CartItem{}).Where("id = ?", id).Exec(ctx)
	return err
}
