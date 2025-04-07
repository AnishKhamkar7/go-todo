package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type Todo struct {
	ent.Schema
}

var defaultValues = []string{"homework", "office", "reading", "chores"}

func (Todo) Fields() []ent.Field {
	return []ent.Field{
		field.Strings("types").Optional().Default(defaultValues),
		field.String("title"),
		field.Enum("status").
			Values("pending", "in_progress", "completed", "cancelled").
			Default("pending"),
		field.Enum("priority").Values("low", "medium", "high").Default("medium"),
	}
}

func (Todo) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).Ref("todos").Unique().Required(),
	}
}
