package dto

import (
	"training-project/internal/model"

	"github.com/shopspring/decimal"
)

type InvoiceDetailResDTO struct {
	Id                 int64           `json:"id"`
	InvoiceId          int64           `json:"invoice_id"`
	ProductId          int64           `json:"product_id"`
	Price              decimal.Decimal `json:"price"`
	DiscountPercentage int32           `json:"discount_percentage"`
	Quantity           int32           `json:"quantity"`
	TotalPrice         decimal.Decimal `json:"total_price"`
}

type InvoiceDetailCreateReqDTO struct {
	InvoiceId int64 `json:"invoice_id"`
	ProductId int64 `json:"product_id"`
	Quantity  int32 `json:"quantity"`
}

type InvoiceDetailDeleteReqDTO struct {
	Id int64 `json:"id"`
}

func ToInvoiceDetailResDTO(invoiceDetail *model.InvoiceDetail) *InvoiceDetailResDTO {
	return &InvoiceDetailResDTO{
		Id:                 invoiceDetail.Id,
		InvoiceId:          invoiceDetail.InvoiceId,
		ProductId:          invoiceDetail.ProductId,
		Price:              invoiceDetail.Price,
		DiscountPercentage: invoiceDetail.DiscountPercentage,
		Quantity:           invoiceDetail.Quantity,
		TotalPrice:         invoiceDetail.TotalPrice,
	}
}

func ToInvoiceDetailResDTOs(invoiceDetails []model.InvoiceDetail) []InvoiceDetailResDTO {
	invoiceDetailResDTOs := make([]InvoiceDetailResDTO, len(invoiceDetails))
	for i, invoiceDetail := range invoiceDetails {
		invoiceDetailResDTOs[i] = *ToInvoiceDetailResDTO(&invoiceDetail)
	}
	return invoiceDetailResDTOs
}
