package schema

import (
	"regexp"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

type User struct {
	ent.Schema
}

func (User) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}

func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("uid").
			MaxLen(20).
			NotEmpty().
			Unique().
			Match(regexp.MustCompile("^[A-Za-z0-9]+$")),
		field.String("password").
			NotEmpty(),
		field.String("email").
			MaxLen(300).
			Nillable().
			Optional().
			Match(regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)),
		field.Enum("role").
			Values("admin", "user").
			Default("user"),
	}
}

func (User) Edges() []ent.Edge {
	return nil
}
