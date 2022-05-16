package main_test

import (
	"context"
	"database/sql"
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
	if _, err := client.Restaurant.Delete().Exec(ctx); err != nil {
		log.Fatalf("failed restaurants users: %v", err)
	}
	// Output:
}
