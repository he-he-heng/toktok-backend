package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("device"),
		field.String("phone"),
		field.String("email").Nillable(),
		field.Enum("sex").Values("male", "female"),
		field.Time("birthday"),
		field.String("mbti").Nillable(),
		field.String("avatar"),
		field.String("nickname"),
		field.String("introduce").Nillable(),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return nil
}
