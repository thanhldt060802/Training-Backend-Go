package dto

import "training-project/internal/model"

type CartItemResDTO struct {
	Id        int64 `json:"id"`
	UserId    int64 `json:"user_id"`
	ProductId int64 `json:"product_id"`
	Quantity  int32 `json:"quantity"`
}

type CartItemCreateReqDTO struct {
	UserId    int64 `json:"user_id"`
	ProductId int64 `json:"product_id"`
}

type CartItemUpdateReqDTO struct {
	Quantity *int32 `json:"quantity"`
}

type CartItemDeleteReqDTO struct {
	Id int64 `json:"id"`
}

func ToCartItemResDTO(cartItem *model.CartItem) *CartItemResDTO {
	return &CartItemResDTO{
		Id:        cartItem.Id,
		UserId:    cartItem.UserId,
		ProductId: cartItem.ProductId,
		Quantity:  cartItem.Quantity,
	}
}

func ToCartItemResDTOs(cartItems []model.CartItem) []CartItemResDTO {
	cartItemResDTOs := make([]CartItemResDTO, len(cartItems))
	for i, cartItem := range cartItems {
		cartItemResDTOs[i] = *ToCartItemResDTO(&cartItem)
	}
	return cartItemResDTOs
}
