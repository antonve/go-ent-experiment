package main_test

import (
	"context"
	"fmt"
	"log"

	"github.com/antonve/go-ent-experiment/ent"

	"entgo.io/ent/dialect"
	_ "github.com/mattn/go-sqlite3"
)

func ExampleEnt() {
	// Create an ent.Client with in-memory SQLite database.
	client, err := ent.Open(dialect.SQLite, "file:ent?mode=memory&cache=shared&_fk=1")
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	defer client.Close()
	ctx := context.Background()
	// Run the automatic migration tool to create all schema resources.
	if err := client.Schema.Create(ctx); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	user1, err := client.User.Create().
		SetUsername("anton").
		SetDisplayName("io").
		Save(ctx)
	if err != nil {
		log.Fatalf("failed to create a user: %v", err)
	}

	_, err = client.Book.Create().
		SetName("The Phoenix Project").
		SetUser(user1).
		Save(ctx)
	if err != nil {
		log.Fatalf("failed to create a book: %v", err)
	}

	_, err = client.Book.Create().
		SetName("The Unicorn Project").
		Save(ctx)
	if err != nil {
		log.Fatalf("failed to create a book: %v", err)
	}

	res, err := user1.QueryBooks().All(ctx)
	if err != nil {
		log.Fatalf("failed to query books: %v", err)
	}
	fmt.Println(res)
	// Output:
	// [Book(id=1, name=The Phoenix Project)]
}
