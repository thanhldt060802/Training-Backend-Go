package repository

import (
	"context"
	"training-project/internal/model"

	"github.com/uptrace/bun"
)

type invoiceRepository struct {
	db *bun.DB
}

type InvoiceRepository interface {
	ExistsById(ctx context.Context, it int64) (bool, error)
	FindAll(ctx context.Context) ([]model.Invoice, error)
	FindById(ctx context.Context, id int64) (*model.Invoice, error)
	FindAllByUserId(ctx context.Context, userId int64) ([]model.Invoice, error)
	Create(ctx context.Context, newInvoice *model.Invoice) error
	UpdateById(ctx context.Context, id int64, updatedInvoice *model.Invoice) error
	DeleteById(ctx context.Context, id int64) error
}

func NewInvoiceRepository(db *bun.DB) InvoiceRepository {
	return &invoiceRepository{db: db}
}

func (invoiceRepository *invoiceRepository) ExistsById(ctx context.Context, id int64) (bool, error) {
	var count int
	count, err := invoiceRepository.db.NewSelect().Model(&model.Invoice{}).Where("id = ?", id).Count(ctx)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func (invoiceRepository *invoiceRepository) FindAll(ctx context.Context) ([]model.Invoice, error) {
	var invoices []model.Invoice
	err := invoiceRepository.db.NewSelect().Model(&invoices).Scan(ctx)
	if err != nil {
		return nil, err
	}
	return invoices, nil
}

func (invoiceRepository *invoiceRepository) FindById(ctx context.Context, id int64) (*model.Invoice, error) {
	var invoice model.Invoice
	err := invoiceRepository.db.NewSelect().Model(&invoice).Where("id = ?", id).Scan(ctx)
	if err != nil {
		return nil, err
	}
	return &invoice, nil
}

func (invoiceRepository *invoiceRepository) FindAllByUserId(ctx context.Context, userId int64) ([]model.Invoice, error) {
	var invoices []model.Invoice
	err := invoiceRepository.db.NewSelect().Model(&invoices).Where("user_id = ?", userId).Scan(ctx)
	if err != nil {
		return nil, err
	}
	return invoices, nil
}

func (invoiceRepository *invoiceRepository) Create(ctx context.Context, newInvoice *model.Invoice) error {
	_, err := invoiceRepository.db.NewInsert().Model(newInvoice).Exec(ctx)
	return err
}

func (invoiceRepository *invoiceRepository) UpdateById(ctx context.Context, id int64, updatedInvoice *model.Invoice) error {
	_, err := invoiceRepository.db.NewUpdate().Model(updatedInvoice).Where("id = ?", id).Exec(ctx)
	return err
}

func (invoiceRepository *invoiceRepository) DeleteById(ctx context.Context, id int64) error {
	_, err := invoiceRepository.db.NewDelete().Model(&model.Invoice{}).Where("id = ?", id).Exec(ctx)
	return err
}
