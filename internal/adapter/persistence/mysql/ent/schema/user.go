package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("uid").
			MinLen(6).
			MaxLen(18).
			NotEmpty().
			Unique(),

		field.String("password").
			NotEmpty(),

		field.String("email").
			Nillable().
			Optional().
			Unique(),

		field.Enum("role").
			Values(
				"admin",
				"user",
			).
			Default("user"),

		field.Enum("ban_state").
			Values(
				"ban",
				"unban",
			).Default("unban"),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("avatar", Avatar.Type).
			Unique(),
	}
}

func (User) Mixin() []ent.Mixin {
	return []ent.Mixin{
		SoftDeleteMixin{},
		TimeMixin{},
	}
}
