package main

import (
	"context"
	"encoding/json"
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

	driver, err := rgdb.New(l, &rgdb.Config{
		ConnString: "postgres://web:qwerty@localhost:5432/reedygreedy?sslmode=disable",
	})

	if err != nil {
		panic(err)
	}

	users, total, err := driver.GetUsers(context.Background(), &rgdbmsg.GetUsersRequest{Sort: []string{"+test"}})

	if err != nil {
		panic(err)
	}

	b, err := json.MarshalIndent(&users, "", "\t")

	if err != nil {
		panic(err)
	}

	fmt.Println(string(b))

	fmt.Println(total)

	fmt.Println(len(users))
}
