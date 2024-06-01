package main

import (
	"context"
	"database/sql"
	"log"
	"time"

	"github.com/immrshc/go-gqlgent/ent"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/stdlib"

	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	_ "github.com/jackc/pgx/v5/stdlib"
)

// Open new connection
func Open(databaseURL string) (*ent.Client, error) {
	cfg, err := pgx.ParseConfig(databaseURL)
	if err != nil {
		return nil, err
	}
	db, err := sql.Open("pgx", stdlib.RegisterConnConfig(cfg))
	if err != nil {
		return nil, err
	}
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(10)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(30 * time.Minute)

	// Create an ent.Driver from `db`.
	drv := entsql.OpenDB(dialect.Postgres, db)
	return ent.NewClient(ent.Driver(drv)), nil
}

func main() {
	const databaseURL = "postgresql://user:password@127.0.0.1:5432/database?sslmode=disable"
	client, err := Open(databaseURL)

	// Your code. For example:
	ctx := context.Background()
	if err := client.Schema.Create(ctx); err != nil {
		log.Fatal(err)
	}
	users, err := client.User.Query().All(ctx)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(users)
}
