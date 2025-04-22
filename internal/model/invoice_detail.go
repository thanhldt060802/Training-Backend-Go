package model

import (
	"github.com/shopspring/decimal"
	"github.com/uptrace/bun"
)

type InvoiceDetail struct {
	bun.BaseModel `bun:"table:invoice_details"`

	Id                 int64           `bun:"id,pk,autoincrement"`
	InvoiceId          int64           `bun:"invoice_id"`
	ProductId          int64           `bun:"product_id"`
	Price              decimal.Decimal `bun:"price"`
	DiscountPercentage int32           `bun:"discount_percentage"`
	Quantity           int32           `bun:"quantity"`
	TotalPrice         decimal.Decimal `bun:"total_price"`
}
