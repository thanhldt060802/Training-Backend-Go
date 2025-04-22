package model

import (
	"time"

	"github.com/shopspring/decimal"
	"github.com/uptrace/bun"
)

type Invoice struct {
	bun.BaseModel `bun:"table:invoices"`

	Id          int64           `bun:"id,pk,autoincrement"`
	UserId      int64           `bun:"user_id"`
	TotalAmount decimal.Decimal `bun:"total_amount"`
	Status      string          `bun:"status"`
	CreatedAt   time.Time       `bun:"created_at,default:current_timestamp"`
}
