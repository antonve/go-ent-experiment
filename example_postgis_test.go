package main_test

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/antonve/go-ent-experiment/ent"

	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	_ "github.com/jackc/pgx/v4/stdlib"
)

func ExamplePostGIS() {
	// Create an ent.Client with in-memory SQLite database.
	db, err := sql.Open("pgx", "host=postgis user=root dbname=experiment password=hunter2 sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	drv := entsql.OpenDB(dialect.Postgres, db)
	client := ent.NewClient(ent.Driver(drv))

	ctx := context.Background()
	// Run the automatic migration tool to create all schema resources.
	if err := client.Schema.Create(ctx); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	// Truncate data
	if _, err := client.User.Delete().Exec(ctx); err != nil {
		log.Fatalf("failed truncating users: %v", err)
	}
	if _, err := client.Book.Delete().Exec(ctx); err != nil {
		log.Fatalf("failed truncating users: %v", err)
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
	fmt.Println(res[0].Name)
	// Output:
	// The Phoenix Project
}
