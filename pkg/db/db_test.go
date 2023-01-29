package db_test

import (
	"context"
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/jackc/pgx/v4/pgxpool"
)

func TestDB(t *testing.T) {
	var err error

	pgHost := "localhost"
	if host := os.Getenv("PGHOST"); host != "" {
		pgHost = host
	}
	dbPool, err := pgxpool.Connect(context.Background(),
		fmt.Sprintf("postgres://%v:%v@%v:%v/%v", "postgres", "pass", pgHost, 5432, "testdb"),
	)
	if err != nil {
		t.Fatal("Fatal error while connecting to postgres: ", err)
	}

	// test db connection
	var greeting string
	err = dbPool.QueryRow(context.Background(), "select 'Hello, world!'").Scan(&greeting)
	if err != nil {
		t.Fatal("Error while making test select statement in postgres: ", err)
	}

	if greeting != "Hello, world!" {
		t.Fatal("Error on simple Postgres smoke test, incorrect result, got: ", greeting)
	}

	log.Println("Successfully connected to postgres!")
}
