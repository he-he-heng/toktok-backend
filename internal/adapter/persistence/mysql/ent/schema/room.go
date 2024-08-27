package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
)

// Room holds the schema definition for the Room entity.
type Room struct {
	ent.Schema
}

// Fields of the Room.
func (Room) Fields() []ent.Field {
	return nil
}

// Edges of the Room.
func (Room) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("avatar", Relation.Type).
			Ref("avatar_rooms").
			Unique(),

		edge.From("friend", Relation.Type).
			Ref("friend_rooms").
			Unique(),

		edge.To("messages", Message.Type).
			Unique(),
	}
}

func (Room) Mixin() []ent.Mixin {
	return []ent.Mixin{
		SoftDeleteMixin{},
		TimeMixin{},
	}
}
