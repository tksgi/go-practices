package main

import (
	_ "github.com/lib/pq"
	"github.com/tksgi/go-practices/web-with-echo/model"
)

func main() {
	defer model.CloseDatabase()

	e := newRouter()
	e.Logger.Fatal(e.Start(":8989"))
}
