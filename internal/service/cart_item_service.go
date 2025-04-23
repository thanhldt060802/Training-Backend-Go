package service

import (
	"context"
	"fmt"
	"training-project/internal/dto"
	"training-project/internal/model"
	"training-project/internal/repository"
	"training-project/pkg/util"
)

type cartItemService struct {
	cartItemRepository repository.CartItemRepository
}

type CartItemService interface {
	IsCartItemExistedById(ctx context.Context, id int64) (bool, error)
	FindAllCartItems(ctx context.Context) ([]model.CartItem, error)
	FindCartItemById(ctx context.Context, id int64) (*model.CartItem, error)
	FindAllCartItemsByUserId(ctx context.Context, userId int64) ([]model.CartItem, error)
	CreateCartItem(ctx context.Context, newCartItem *dto.CartItemCreateReqDTO) error
	UpdateCartItemById(ctx context.Context, id int64, updatedCartItem *dto.CartItemUpdateReqDTO) error
	DeleteCartItemById(ctx context.Context, id int64) error
}

func NewCartItemService(cartItemRepository repository.CartItemRepository) CartItemService {
	return &cartItemService{cartItemRepository: cartItemRepository}
}

func (cartItemService *cartItemService) IsCartItemExistedById(ctx context.Context, id int64) (bool, error) {
	existed, err := cartItemService.cartItemRepository.ExistsById(ctx, id)
	if err != nil {
		return false, err
	}
	return existed, nil
}

func (cartItemService *cartItemService) FindAllCartItems(ctx context.Context) ([]model.CartItem, error) {
	return cartItemService.cartItemRepository.FindAll(ctx)
}

func (cartItemService *cartItemService) FindCartItemById(ctx context.Context, id int64) (*model.CartItem, error) {
	return cartItemService.cartItemRepository.FindById(ctx, id)
}

func (cartItemService *cartItemService) FindAllCartItemsByUserId(ctx context.Context, userId int64) ([]model.CartItem, error) {
	return cartItemService.cartItemRepository.FindAllByUserId(ctx, userId)
}

func (cartItemService *cartItemService) CreateCartItem(ctx context.Context, cartItemCreateReqDTO *dto.CartItemCreateReqDTO) error {
	newCartItem := model.CartItem{
		UserId:    cartItemCreateReqDTO.UserId,
		ProductId: cartItemCreateReqDTO.ProductId,
		Quantity:  1,
	}
	return cartItemService.cartItemRepository.Create(ctx, &newCartItem)
}

func (cartItemService *cartItemService) UpdateCartItemById(ctx context.Context, id int64, updatedCartItem *dto.CartItemUpdateReqDTO) error {
	foundCartItem, err := cartItemService.FindCartItemById(ctx, id)
	if err != nil {
		return err
	}
	util.ApplyCartItemUpdate(foundCartItem, updatedCartItem)
	return cartItemService.cartItemRepository.UpdateById(ctx, id, foundCartItem)
}

func (cartItemService *cartItemService) DeleteCartItemById(ctx context.Context, id int64) error {
	existed, err := cartItemService.IsCartItemExistedById(ctx, id)
	if err != nil {
		return err
	}
	if !existed {
		return fmt.Errorf("id of cart item is not valid")
	}
	return cartItemService.cartItemRepository.DeleteById(ctx, id)
}
