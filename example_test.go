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

	fmt.Println(user1.Username)
	// Output:
	// anton
}
