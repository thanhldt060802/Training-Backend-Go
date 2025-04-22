package model

import "github.com/uptrace/bun"

type User struct {
	bun.BaseModel `bun:"table:users"`

	Id       int64  `bun:"id,pk,autoincrement"`
	Name     string `bun:"name"`
	Email    string `bun:"email"`
	Username string `bun:"username"`
	Password string `bun:"password"`
	Address  string `bun:"address"`
	Role     string `bun:"role"`
}
