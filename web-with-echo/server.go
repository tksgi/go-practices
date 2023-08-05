package main

import (
	"context"
	"database/sql"
	"log"
	"net/http"
  "fmt"

	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
)

func main() {
  sqldb, err := sql.Open("postgres", "host=db port=5432 user=user password=password sslmode=disable")
  if err != nil {
    log.Fatal(err)
  }
  defer sqldb.Close()

  db := bun.NewDB(sqldb, pgdialect.New())
  defer db.Close()

  ctx := context.Background()
  _, err = db.NewCreateTable().Model((*User)(nil)).IfNotExists().Exec(ctx)
  if err != nil {
    log.Fatal(err)
  }

  e := echo.New()
  e.GET("/", func(c echo.Context) error {
    var users []User
    ctx := context.Background()
    err := db.NewSelect().Model(&users).Scan(ctx)
    if err != nil {
      e.Logger.Error(err)
      return c.String(http.StatusBadRequest, "cannot get Users")
    }
    return c.String(http.StatusOK, fmt.Sprintf("%+v", users))
  })
  e.Logger.Fatal(e.Start(":8989"))
}

type User struct {
  bun.BaseModel `bun:"table:users,alias:u"`

  Id  int64 `bun:"id,pk,autoincrement"`
  Name string `bun:"name,notnull"`
}
