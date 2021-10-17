package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Command holds the schema definition for the Command entity.
type Command struct {
	ent.Schema
}

// Fields of the Command.
func (Command) Fields() []ent.Field {
	return []ent.Field{
		field.String("keyword").
			Unique().
			NotEmpty(),
		field.Time("created_at").
			Default(func() time.Time {
				return time.Now()
			}),
		field.Time("updated_at").
			Default(func() time.Time {
				return time.Now()
			}),
		field.String("creator").
			NotEmpty(),
		field.String("code"),
	}
}

// Edges of the Command.
func (Command) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("logs", Command.Type),
	}
}
