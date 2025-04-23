package dto

import (
	"time"
	"training-project/internal/model"

	"github.com/shopspring/decimal"
)

type InvoiceResDTO struct {
	Id          int64           `json:"id"`
	UserId      int64           `json:"user_id"`
	TotalAmount decimal.Decimal `json:"total_amount"`
	Status      string          `json:"status"`
	CreatedAt   time.Time       `json:"created_at"`
}

type InvoiceCreateReqDTO struct {
	UserId      int64           `json:"user_id"`
	TotalAmount decimal.Decimal `json:"total_amount"`
	Status      string          `json:"status"`
}

type InvoiceUpdateReqDTO struct {
	Status *string `json:"status"`
}

type InvoiceDeleteReqDTO struct {
	Id int64 `json:"id"`
}

func ToInvoiceResDTO(invoice *model.Invoice) *InvoiceResDTO {
	return &InvoiceResDTO{
		Id:          invoice.Id,
		UserId:      invoice.UserId,
		TotalAmount: invoice.TotalAmount,
		Status:      invoice.Status,
		CreatedAt:   invoice.CreatedAt,
	}
}

func ToInvoiceResDTOs(invoices []model.Invoice) []InvoiceResDTO {
	invoiceResDTOs := make([]InvoiceResDTO, len(invoices))
	for i, invoice := range invoices {
		invoiceResDTOs[i] = *ToInvoiceResDTO(&invoice)
	}
	return invoiceResDTOs
}
