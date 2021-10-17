package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// ResultLog holds the schema definition for the ResultLog entity.
type ResultLog struct {
	ent.Schema
}

// Fields of the ResultLog.
func (ResultLog) Fields() []ent.Field {
	return []ent.Field{
		field.String("level"),
		field.String("log"),
	}
}

// Edges of the ResultLog.
func (ResultLog) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("command", Command.Type).
			Ref("logs"),
	}
}
