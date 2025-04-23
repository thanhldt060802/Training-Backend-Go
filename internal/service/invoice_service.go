package service

import (
	"context"
	"fmt"
	"training-project/internal/dto"
	"training-project/internal/model"
	"training-project/internal/repository"
	"training-project/pkg/util"
)

type invoiceService struct {
	invoiceRepository repository.InvoiceRepository
}

type InvoiceService interface {
	IsInvoiceExistedById(ctx context.Context, id int64) (bool, error)
	FindAllInvoices(ctx context.Context) ([]model.Invoice, error)
	FindInvoiceById(ctx context.Context, id int64) (*model.Invoice, error)
	FindAllInvoicesByUserId(ctx context.Context, userId int64) ([]model.Invoice, error)
	CreateInvoice(ctx context.Context, newInvoice *dto.InvoiceCreateReqDTO) error
	UpdateInvoiceById(ctx context.Context, id int64, updatedInvoice *dto.InvoiceUpdateReqDTO) error
	DeleteInvoiceById(ctx context.Context, id int64) error
}

func NewInvoiceService(invoiceRepository repository.InvoiceRepository) InvoiceService {
	return &invoiceService{invoiceRepository: invoiceRepository}
}

func (invoiceService *invoiceService) IsInvoiceExistedById(ctx context.Context, id int64) (bool, error) {
	existed, err := invoiceService.invoiceRepository.ExistsById(ctx, id)
	if err != nil {
		return false, err
	}
	return existed, nil
}

func (invoiceService *invoiceService) FindAllInvoices(ctx context.Context) ([]model.Invoice, error) {
	return invoiceService.invoiceRepository.FindAll(ctx)
}

func (invoiceService *invoiceService) FindInvoiceById(ctx context.Context, id int64) (*model.Invoice, error) {
	return invoiceService.invoiceRepository.FindById(ctx, id)
}

func (invoiceService *invoiceService) FindAllInvoicesByUserId(ctx context.Context, userId int64) ([]model.Invoice, error) {
	return invoiceService.invoiceRepository.FindAllByUserId(ctx, userId)
}

func (invoiceService *invoiceService) CreateInvoice(ctx context.Context, invoiceCreateReqDTO *dto.InvoiceCreateReqDTO) error {
	newInvoice := model.Invoice{
		UserId:      invoiceCreateReqDTO.UserId,
		TotalAmount: invoiceCreateReqDTO.TotalAmount,
		Status:      invoiceCreateReqDTO.Status,
	}
	return invoiceService.invoiceRepository.Create(ctx, &newInvoice)
}

func (invoiceService *invoiceService) UpdateInvoiceById(ctx context.Context, id int64, updatedInvoice *dto.InvoiceUpdateReqDTO) error {
	foundInvoice, err := invoiceService.FindInvoiceById(ctx, id)
	if err != nil {
		return err
	}
	util.ApplyInvoiceUpdate(foundInvoice, updatedInvoice)
	return invoiceService.invoiceRepository.UpdateById(ctx, id, foundInvoice)
}

func (invoiceService *invoiceService) DeleteInvoiceById(ctx context.Context, id int64) error {
	existed, err := invoiceService.IsInvoiceExistedById(ctx, id)
	if err != nil {
		return err
	}
	if !existed {
		return fmt.Errorf("id of invoice is not valid")
	}
	return invoiceService.invoiceRepository.DeleteById(ctx, id)
}
