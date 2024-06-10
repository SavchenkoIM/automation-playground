package main

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"os"
)

func main() {

	pool, _ := pgxpool.New(context.Background(), "postgresql://postgres:postgres@postgres/postgres?sslmode=disable")

	_, err := pool.Query(context.Background(), "SELECT * from information_schema.tables")
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}

	println("Hello World")
}
