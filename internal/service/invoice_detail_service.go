package service

import (
	"context"
	"fmt"
	"training-project/internal/dto"
	"training-project/internal/model"
	"training-project/internal/repository"

	"github.com/shopspring/decimal"
)

type invoiceDetailService struct {
	invoiceDetailRepository repository.InvoiceDetailRepository
	invoiceService          InvoiceService
	productService          ProductService
}

type InvoiceDetailService interface {
	IsInvoiceDetailExistedById(ctx context.Context, id int64) (bool, error)
	FindAllInvoiceDetails(ctx context.Context) ([]model.InvoiceDetail, error)
	FindInvoiceDetailById(ctx context.Context, id int64) (*model.InvoiceDetail, error)
	FindAllInvoiceDetailsByInvoiceId(ctx context.Context, invoiceId int64) ([]model.InvoiceDetail, error)
	CreateInvoiceDetail(ctx context.Context, newInvoiceDetail *dto.InvoiceDetailCreateReqDTO) error
	DeleteInvoiceDetailById(ctx context.Context, id int64) error
}

func NewInvoiceDetailService(invoiceDetailRepository repository.InvoiceDetailRepository, invoiceService InvoiceService, productService ProductService) InvoiceDetailService {
	return &invoiceDetailService{
		invoiceDetailRepository: invoiceDetailRepository,
		invoiceService:          invoiceService,
		productService:          productService,
	}
}

func (invoiceDetailService *invoiceDetailService) IsInvoiceDetailExistedById(ctx context.Context, id int64) (bool, error) {
	existed, err := invoiceDetailService.invoiceDetailRepository.ExistsById(ctx, id)
	if err != nil {
		return false, err
	}
	return existed, nil
}

func (invoiceDetailService *invoiceDetailService) FindAllInvoiceDetails(ctx context.Context) ([]model.InvoiceDetail, error) {
	return invoiceDetailService.invoiceDetailRepository.FindAll(ctx)
}

func (invoiceDetailService *invoiceDetailService) FindInvoiceDetailById(ctx context.Context, id int64) (*model.InvoiceDetail, error) {
	return invoiceDetailService.invoiceDetailRepository.FindById(ctx, id)
}

func (invoiceDetailService *invoiceDetailService) FindAllInvoiceDetailsByInvoiceId(ctx context.Context, invoiceId int64) ([]model.InvoiceDetail, error) {
	return invoiceDetailService.invoiceDetailRepository.FindAllByInvoiceId(ctx, invoiceId)
}

func (invoiceDetailService *invoiceDetailService) CreateInvoiceDetail(ctx context.Context, invoiceDetailCreateReqDTO *dto.InvoiceDetailCreateReqDTO) error {
	foundProduct, err := invoiceDetailService.productService.FindProductById(ctx, invoiceDetailCreateReqDTO.ProductId)
	if err != nil {
		return err
	}
	totalPrice := decimal.NewFromInt32(invoiceDetailCreateReqDTO.Quantity).Mul(foundProduct.Price.Mul(decimal.NewFromFloat(float64(100-foundProduct.DiscountPercentage) / 100)))
	newInvoiceDetail := model.InvoiceDetail{
		InvoiceId:          invoiceDetailCreateReqDTO.InvoiceId,
		ProductId:          invoiceDetailCreateReqDTO.ProductId,
		Price:              foundProduct.Price,
		DiscountPercentage: foundProduct.DiscountPercentage,
		Quantity:           invoiceDetailCreateReqDTO.Quantity,
		TotalPrice:         totalPrice,
	}
	return invoiceDetailService.invoiceDetailRepository.Create(ctx, &newInvoiceDetail)
}

func (invoiceDetailService *invoiceDetailService) DeleteInvoiceDetailById(ctx context.Context, id int64) error {
	existed, err := invoiceDetailService.IsInvoiceDetailExistedById(ctx, id)
	if err != nil {
		return err
	}
	if !existed {
		return fmt.Errorf("id of invoice detail is not valid")
	}
	return invoiceDetailService.invoiceDetailRepository.DeteleById(ctx, id)
}
