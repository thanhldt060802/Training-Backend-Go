package model

import "github.com/uptrace/bun"

type Category struct {
	bun.BaseModel `bun:"table:categories"`

	Id          int64  `bun:"id,pk,autoincrement"`
	Name        string `bun:"name"`
	Description string `bun:"description"`
}
