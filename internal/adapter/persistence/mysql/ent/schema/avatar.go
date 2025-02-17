package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Avatar holds the schema definition for the Avatar entity.
type Avatar struct {
	ent.Schema
}

// Fields of the Avatar.
func (Avatar) Fields() []ent.Field {
	return []ent.Field{
		field.Enum("sex").
			Values(
				"male",
				"female",
			),

		field.String("birthday").
			MinLen(8).
			MaxLen(8).
			NotEmpty(),

		field.String("mbti").
			MinLen(4).
			MaxLen(4).
			Nillable().
			Optional(),

		field.Enum("picture").
			Values(
				"lightRed",
				"milkRed",
				"lightOrange",
				"deepYellow",
				"deepPurple",
				"bluePurple",
				"lightPink",
				"deepSky",
			).Default("lightRed"),

		field.String("nickname").
			NotEmpty().
			Unique(),

		field.String("introduce").
			Nillable().
			Optional(),

		field.Enum("state").
			Values(
				"online",
				"offline",
			).
			Default("online"),
	}
}

// Edges of the Avatar.
func (Avatar) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).
			Ref("avatar").
			Unique().
			Required(),

		edge.To("avatar_relations", Relation.Type),
		edge.To("friend_relations", Relation.Type),

		edge.To("messages", Message.Type),
	}
}

func (Avatar) Mixin() []ent.Mixin {
	return []ent.Mixin{
		SoftDeleteMixin{},
		TimeMixin{},
	}
}
