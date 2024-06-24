package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Attempt holds the schema definition for the Attempt entity.
type Attempt struct {
	ent.Schema
}

// Fields of the Attempt.
func (Attempt) Fields() []ent.Field {
	return []ent.Field{
		field.String("phone").Unique(),
		field.Int("authcode"),
		field.Int("cnt").Default(0),
		field.Bool("break").Default(false),
		field.Time("timestamp").Default(time.Now),
	}
}

// Edges of the Attempt.
func (Attempt) Edges() []ent.Edge {
	return nil
}
