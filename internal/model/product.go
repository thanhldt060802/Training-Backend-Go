package model

import (
	"github.com/shopspring/decimal"
	"github.com/uptrace/bun"
)

type Product struct {
	bun.BaseModel `bun:"table:products"`

	Id          int64           `bun:"id,pk,autoincrement"`
	Name        string          `bun:"name"`
	Description string          `bun:"description"`
	Price       decimal.Decimal `bun:"price"`
	Stock       int32           `bun:"stock"`
	ImageURL    string          `bun:"image_url"`
	CategoryId  int64           `bun:"category_id"`
}
