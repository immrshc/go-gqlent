package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	"github.com/immrshc/go-gqlent/ent/todo"
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
	if err := client.Schema.Create(ctx); err != nil {
		log.Fatal(err)
	}
	//task1, err := client.Todo.Create().SetText("Add GraphQL Example").Save(ctx)
	//if err != nil {
	//	log.Fatalf("failed creating a todo: %v", err)
	//}
	//fmt.Printf("%d: %q\n", task1.ID, task1.Text)
	//task2, err := client.Todo.Create().SetText("Add Tracing Example").Save(ctx)
	//if err != nil {
	//	log.Fatalf("failed creating a todo: %v", err)
	//}
	//fmt.Printf("%d: %q\n", task2.ID, task2.Text)
	//if err = task2.Update().SetParent(task1).Exec(ctx); err != nil {
	//	log.Fatalf("failed connecting todo2 to its parent: %v", err)
	//}
	items, err := client.Todo.Query().Where(todo.HasParent()).All(ctx)
	if err != nil {
		log.Fatalf("failed querying todos: %v", err)
	}
	for _, t := range items {
		fmt.Printf("%d: %q\n", t.ID, t.Text)
	}
}
