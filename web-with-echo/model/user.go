package model

import (
	"context"
	"fmt"

	"github.com/uptrace/bun"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	bun.BaseModel `bun:"table:users,alias:u"`

	Id   int64  `bun:"id,pk,autoincrement"`
	Name string `bun:"name,notnull"`
  Password string `bun:"password,notnull"`
}

func (user *User) String() string {
  return fmt.Sprintf("Id=%d, Name=%s", user.Id, user.Name)
}

func ListAllUsers() (*[]User, error) {
	var users []User
	ctx := context.Background()
	err := db.NewSelect().Model(&users).Scan(ctx)

	return &users, err
}

func CreateUser(user *User) error {
	ctx := context.Background()
  password, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
  if err != nil {
    return err
  }
  user.Password = string(password)

  _, err = db.NewInsert().Model(user).Exec(ctx)
  if err != nil {
    return err
  }

  // LastInsertId is not supported by this driver
  // userId, err := res.LastInsertId()
  // if err != nil {
  //   return err
  // }

  err = db.NewSelect().Model(user).Where("name = ?", user.Name).Scan(ctx)
  return err
}
