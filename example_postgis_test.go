package main_test

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/antonve/go-ent-experiment/ent"
	"github.com/antonve/go-ent-experiment/ent/schema"

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

	// Create restaurants
	_ = createRestaurant(client, "CoCo Ichibanya Ebisu", 35.64699825984844, 139.71194575396922)
	_ = createRestaurant(client, "SHAKE SHACK Ebisu", 35.64669248211187, 139.70949784477963)
	_ = createRestaurant(client, "Ichiran Ramen Shinjuku", 35.69079988476277, 139.70286473414785)
	_ = createRestaurant(client, "Torikizoku Shinjuku", 35.68918337273537, 139.70249991934935)

	// List all restaurants
	restaurants, err := client.Restaurant.Query().All(ctx)
	if err != nil {
		log.Fatalf("could not fetch restaurant list: %v", err)
	}
	fmt.Println("restaurants length", len(restaurants))

	ebisu := &schema.Point{35.64699709191131, 139.71000533635765}

	query := client.Restaurant.Query().
		Where(func(s *entsql.Selector) {
			s.Where(entsql.ExprP("ST_Distance(location, GeomFromEWKB($1), false) < $2", ebisu, 300.0))
		})

	ebisuRestaurants := query.AllX(ctx)

	for _, r := range ebisuRestaurants {
		fmt.Println(r.Name)
	}

	// Output:
	// restaurants length 4
	// CoCo Ichibanya Ebisu
	// SHAKE SHACK Ebisu
}

func createRestaurant(client *ent.Client, name string, long, lat float64) *ent.Restaurant {
	r := client.Restaurant.Create().
		SetName(name).
		SetLocation(&schema.Point{long, lat}).
		SaveX(context.Background())

	return r
}
