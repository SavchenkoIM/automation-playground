package main

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {

	pool, _ := pgxpool.New(context.Background(), "postgresql://localhost:5432/postgres?user=postgres&password=postgres")

	pool.Query(context.Background(), "SELECT * from information_schema.tables")

	println("Hello World")
}
