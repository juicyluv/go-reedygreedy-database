package main

import (
	"context"
	"fmt"
	"go.uber.org/zap"
	"rgdb/rgdb"
	"rgdb/rgdbmsg"
)

func main() {
	l, err := zap.NewProduction()

	if err != nil {
		panic(err)
	}

	driver, err := rgdb.New(l, &rgdb.Config{ConnString: "postgres://web:qwerty@localhost:5432/reedygreedy?sslmode=disable"})

	if err != nil {
		panic(err)
	}

	id, total, err := driver.GetBooks(context.Background(), &rgdbmsg.GetBooksRequest{})

	if err != nil {
		panic(err)
	}

	fmt.Println(total, len(id))
}
