package model

import (
	"context"

	"github.com/uptrace/bun"
)

type User struct {
	bun.BaseModel `bun:"table:users,alias:u"`

	Id   int64  `bun:"id,pk,autoincrement"`
	Name string `bun:"name,notnull"`
}

func ListAllUsers() ([]User, error) {
	var users []User
	ctx := context.Background()
	err := db.NewSelect().Model(&users).Scan(ctx)

	return users, err
}
