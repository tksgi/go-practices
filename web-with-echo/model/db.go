package model

import (
	"context"
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
)

var db *bun.DB
var _sqldb *sql.DB

func init() {
	var err error
	_sqldb, err = sql.Open("postgres", "host=db port=5432 user=user password=password sslmode=disable")
	if err != nil {
		log.Fatal(err)
		panic("failed to open database")
	}

	db = bun.NewDB(_sqldb, pgdialect.New())

	ctx := context.Background()
	_, err = db.NewCreateTable().Model((*User)(nil)).IfNotExists().Exec(ctx)
	if err != nil {
		log.Fatal(err)
	}

}

func CloseDatabase() {
	db.Close()
	_sqldb.Close()
}
