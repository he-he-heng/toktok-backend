package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Relation holds the schema definition for the Relation entity.
type Relation struct {
	ent.Schema
}

// Fields of the Relation.
func (Relation) Fields() []ent.Field {
	return []ent.Field{
		field.Enum("state").
			Values(
				"request-friend",
				"pending",
				"friend",
				"decline",
				"remove",
			).
			Default("pending"),

		field.Enum("alertState").
			Values(
				"allow",
				"deny",
			).
			Default("allow"),
	}
}

// Edges of the Relation.
func (Relation) Edges() []ent.Edge {
	return []ent.Edge{

		edge.From("avatar", Avatar.Type).
			Ref("avatarRelations").
			Unique(),

		edge.From("friend", Avatar.Type).
			Ref("friendRelations").
			Unique(),

		edge.To("messages", Message.Type),
	}
}

func (Relation) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}
