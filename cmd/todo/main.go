package main

import (
	"context"
	"database/sql"
	"log"
	"net/http"

	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	todo "github.com/immrshc/go-gqlent"
	"github.com/immrshc/go-gqlent/ent/migrate"
	_ "github.com/jackc/pgx/v5/stdlib"

	"github.com/immrshc/go-gqlent/ent"
)

func Open(dbURL string) *ent.Client {
	db, err := sql.Open("pgx", dbURL)
	if err != nil {
		log.Fatal(err)
	}
	drv := entsql.OpenDB(dialect.Postgres, db)
	return ent.NewClient(ent.Driver(drv))
}

func main() {
	client := Open("postgresql://postgres:postgres@127.0.0.1:5432/gqlent")
	ctx := context.Background()
	if err := client.Schema.Create(
		ctx,
		migrate.WithGlobalUniqueID(true),
	); err != nil {
		log.Fatal(err)
	}
	// Configure the server and start listening on :8081.
	srv := handler.NewDefaultServer(todo.NewSchema(client))
	http.Handle("/",
		playground.Handler("Todo", "/query"),
	)
	http.Handle("/query", srv)
	log.Println("listening on :8081")
	if err := http.ListenAndServe(":8081", nil); err != nil {
		log.Fatal("http server terminated", err)
	}
}
