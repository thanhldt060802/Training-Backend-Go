package dto

import (
	"training-project/internal/model"

	"github.com/shopspring/decimal"
)

type ProductResDTO struct {
	Id                 int64           `json:"id"`
	Name               string          `json:"name"`
	Description        string          `json:"description"`
	Price              decimal.Decimal `json:"price"`
	DiscountPercentage int32           `json:"discount_percentage"`
	Stock              int32           `json:"stock"`
	ImageURL           string          `json:"image_url"`
	CategoryId         int64           `json:"category_id"`
}

type ProductCreateReqDTO struct {
	Name               string          `json:"name"`
	Description        string          `json:"description"`
	Price              decimal.Decimal `json:"price"`
	DiscountPercentage int32           `json:"discount_percentage"`
	Stock              int32           `json:"stock"`
	ImageURL           string          `json:"image_url"`
	CategoryId         int64           `json:"category_id"`
}

type ProductUpdateReqDTO struct {
	Name               *string          `json:"name"`
	Description        *string          `json:"description"`
	Price              *decimal.Decimal `json:"price"`
	DiscountPercentage *int32           `json:"discount_percentage"`
	Stock              *int32           `json:"stock"`
	ImageURL           *string          `json:"image_url"`
	CategoryId         *int64           `json:"category_id"`
}

type ProductDeleteReqDTO struct {
	Id int64 `json:"id"`
}

func ToProductResDTO(product *model.Product) *ProductResDTO {
	return &ProductResDTO{
		Id:                 product.Id,
		Name:               product.Name,
		Description:        product.Description,
		Price:              product.Price,
		DiscountPercentage: product.DiscountPercentage,
		Stock:              product.Stock,
		ImageURL:           product.ImageURL,
		CategoryId:         product.CategoryId,
	}
}

func ToProductResDTOs(products []model.Product) []ProductResDTO {
	productResDTOs := make([]ProductResDTO, len(products))
	for i, product := range products {
		productResDTOs[i] = *ToProductResDTO(&product)
	}
	return productResDTOs
}
