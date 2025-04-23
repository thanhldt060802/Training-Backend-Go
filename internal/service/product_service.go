package service

import (
	"context"
	"fmt"
	"training-project/internal/dto"
	"training-project/internal/model"
	"training-project/internal/repository"
	"training-project/pkg/util"
)

type productService struct {
	productRepository repository.ProductRepository
}

type ProductService interface {
	IsProductExistedById(ctx context.Context, id int64) (bool, error)
	FindAllProducts(ctx context.Context) ([]model.Product, error)
	FindProductById(ctx context.Context, id int64) (*model.Product, error)
	CreateProduct(ctx context.Context, newProduct *dto.ProductCreateReqDTO) error
	UpdateProductById(ctx context.Context, id int64, updatedProduct *dto.ProductUpdateReqDTO) error
	DeleteProductById(ctx context.Context, id int64) error
}

func NewProductService(productRepository repository.ProductRepository) ProductService {
	return &productService{productRepository: productRepository}
}

func (productService *productService) IsProductExistedById(ctx context.Context, id int64) (bool, error) {
	existed, err := productService.productRepository.ExistsById(ctx, id)
	if err != nil {
		return false, err
	}
	return existed, nil
}

func (productService *productService) FindAllProducts(ctx context.Context) ([]model.Product, error) {
	return productService.productRepository.FindAll(ctx)
}

func (productService *productService) FindProductById(ctx context.Context, id int64) (*model.Product, error) {
	return productService.productRepository.FindById(ctx, id)
}

func (productService *productService) CreateProduct(ctx context.Context, productCreateReqDTO *dto.ProductCreateReqDTO) error {
	newProduct := model.Product{
		Name:               productCreateReqDTO.Name,
		Description:        productCreateReqDTO.Description,
		Price:              productCreateReqDTO.Price,
		DiscountPercentage: productCreateReqDTO.DiscountPercentage,
		Stock:              productCreateReqDTO.Stock,
		ImageURL:           productCreateReqDTO.ImageURL,
		CategoryId:         productCreateReqDTO.CategoryId,
	}
	return productService.productRepository.Create(ctx, &newProduct)
}

func (productService *productService) UpdateProductById(ctx context.Context, id int64, updatedProduct *dto.ProductUpdateReqDTO) error {
	foundProduct, err := productService.FindProductById(ctx, id)
	if err != nil {
		return err
	}
	util.ApplyProductUpdate(foundProduct, updatedProduct)
	return productService.productRepository.UpdateById(ctx, id, foundProduct)
}

func (productService *productService) DeleteProductById(ctx context.Context, id int64) error {
	existed, err := productService.IsProductExistedById(ctx, id)
	if err != nil {
		return err
	}
	if !existed {
		return fmt.Errorf("id of product is not valid")
	}
	return productService.productRepository.DeleteById(ctx, id)
}
