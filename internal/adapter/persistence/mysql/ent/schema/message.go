package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Message holds the schema definition for the Message entity.
type Message struct {
	ent.Schema
}

// Fields of the Message.
func (Message) Fields() []ent.Field {
	return []ent.Field{
		field.Enum("state").
			Values(
				"check",
				"uncheck",
			),

		field.String("content").
			NotEmpty(),

		field.Time("entered_at").
			Default(time.Now),
	}
}

// Edges of the Message.
func (Message) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("relation", Relation.Type).
			Ref("messages").
			Unique(),

		edge.From("avatar", Avatar.Type).
			Ref("messages").
			Unique(),
	}
}

func (Message) Mixin() []ent.Mixin {
	return []ent.Mixin{
		SoftDeleteMixin{},
		TimeMixin{},
	}
}
