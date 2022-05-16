package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/field"
)

// Restaurant holds the schema definition for the Restaurant entity.
type Restaurant struct {
	ent.Schema
}

// Fields of the Restaurant.
func (Restaurant) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.Other("location", &Point{}).
			SchemaType(map[string]string{
				dialect.Postgres: "geometry(point, 4326)",
			}),
	}
}

// Edges of the Restaurant.
func (Restaurant) Edges() []ent.Edge {
	return nil
}
