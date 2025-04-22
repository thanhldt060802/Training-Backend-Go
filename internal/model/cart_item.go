package model

import "github.com/uptrace/bun"

type CartItem struct {
	bun.BaseModel `bun:"table:cart_items"`

	Id        int64 `bun:"id,pk,autoincrement"`
	UserId    int64 `bun:"user_id"`
	ProductId int64 `bun:"product_id"`
	Quantity  int32 `bun:"quantity"`
}
