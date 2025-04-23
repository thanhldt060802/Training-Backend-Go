package repository

import (
	"context"
	"training-project/internal/model"

	"github.com/uptrace/bun"
)

type invoiceDetailRepository struct {
	db *bun.DB
}

type InvoiceDetailRepository interface {
	ExistsById(ctx context.Context, id int64) (bool, error)
	FindAll(ctx context.Context) ([]model.InvoiceDetail, error)
	FindById(ctx context.Context, id int64) (*model.InvoiceDetail, error)
	FindAllByInvoiceId(ctx context.Context, invoiceId int64) ([]model.InvoiceDetail, error)
	Create(ctx context.Context, newInvoiceDetail *model.InvoiceDetail) error
	DeteleById(ctx context.Context, invoiceId int64) error
}

func NewInvoiceDetailRepository(db *bun.DB) InvoiceDetailRepository {
	return &invoiceDetailRepository{db: db}
}

func (invoiceDetailRepository *invoiceDetailRepository) ExistsById(ctx context.Context, id int64) (bool, error) {
	var count int
	count, err := invoiceDetailRepository.db.NewSelect().Model(&model.InvoiceDetail{}).Where("id = ?", id).Count(ctx)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func (invoiceDetailRepository *invoiceDetailRepository) FindAll(ctx context.Context) ([]model.InvoiceDetail, error) {
	var invoiceDetails []model.InvoiceDetail
	err := invoiceDetailRepository.db.NewSelect().Model(&invoiceDetails).Scan(ctx)
	if err != nil {
		return nil, err
	}
	return invoiceDetails, nil
}

func (invoiceDetailRepository *invoiceDetailRepository) FindById(ctx context.Context, id int64) (*model.InvoiceDetail, error) {
	var invoiceDetail model.InvoiceDetail
	err := invoiceDetailRepository.db.NewSelect().Model(&invoiceDetail).Where("id = ?", id).Scan(ctx)
	if err != nil {
		return nil, err
	}
	return &invoiceDetail, nil
}

func (invoiceDetailRepository *invoiceDetailRepository) FindAllByInvoiceId(ctx context.Context, invoiceId int64) ([]model.InvoiceDetail, error) {
	var invoiceDetails []model.InvoiceDetail
	err := invoiceDetailRepository.db.NewSelect().Model(&invoiceDetails).Where("invoice_id = ?", invoiceId).Scan(ctx)
	if err != nil {
		return nil, err
	}
	return invoiceDetails, nil
}

func (invoiceDetailRepository *invoiceDetailRepository) Create(ctx context.Context, newInvoiceDetail *model.InvoiceDetail) error {
	_, err := invoiceDetailRepository.db.NewInsert().Model(newInvoiceDetail).Exec(ctx)
	return err
}

func (invoiceDetailRepository *invoiceDetailRepository) DeteleById(ctx context.Context, invoiceId int64) error {
	_, err := invoiceDetailRepository.db.NewDelete().Model(&model.InvoiceDetail{}).Where("invoice_id = ?", invoiceId).Exec(ctx)
	return err
}
