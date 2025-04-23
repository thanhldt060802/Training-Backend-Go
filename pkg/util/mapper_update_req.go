package util

import (
	"training-project/internal/dto"
	"training-project/internal/model"
)

func ApplyUserUpdate(user *model.User, dto *dto.UserUpdateReqDTO) {
	if dto.Name != nil {
		user.Name = *dto.Name
	}
	if dto.Email != nil {
		user.Email = *dto.Email
	}
	if dto.Password != nil {
		user.Password = *dto.Password
	}
	if dto.Address != nil {
		user.Address = *dto.Address
	}
}

func ApplyCategoryUpdate(category *model.Category, dto *dto.CategoryUpdateReqDTO) {
	if dto.Name != nil {
		category.Name = *dto.Name
	}
	if dto.Description != nil {
		category.Description = *dto.Description
	}
}

func ApplyProductUpdate(product *model.Product, dto *dto.ProductUpdateReqDTO) {
	if dto.Name != nil {
		product.Name = *dto.Name
	}
	if dto.Description != nil {
		product.Description = *dto.Description
	}
	if dto.Price != nil {
		product.Price = *dto.Price
	}
	if dto.DiscountPercentage != nil {
		product.DiscountPercentage = *dto.DiscountPercentage
	}
	if dto.Stock != nil {
		product.Stock = *dto.Stock
	}
	if dto.ImageURL != nil {
		product.ImageURL = *dto.ImageURL
	}
	if dto.CategoryId != nil {
		product.CategoryId = *dto.CategoryId
	}
}

func ApplyCartItemUpdate(cartItem *model.CartItem, dto *dto.CartItemUpdateReqDTO) {
	if dto.Quantity != nil {
		cartItem.Quantity = *dto.Quantity
	}
}

func ApplyInvoiceUpdate(invoice *model.Invoice, dto *dto.InvoiceUpdateReqDTO) {
	if dto.Status != nil {
		invoice.Status = *dto.Status
	}
}
