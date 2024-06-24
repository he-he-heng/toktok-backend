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
			MaxLen(20).
			NotEmpty().
			Match(regexp.MustCompile("^[A-Za-z0-9!@#$%^&*]+$")),
		field.Enum("role").
			Values("admin", "user").
			Default("user"),
	}
}

func (User) Edges() []ent.Edge {
	return nil
}
