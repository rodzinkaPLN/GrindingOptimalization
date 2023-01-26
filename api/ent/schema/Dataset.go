package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

type Dataset struct {
	ent.Schema
}

func (Dataset) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New),
		field.String("name"),
		field.Time("created_at").
			Default(time.Now),
	}
}

func (Dataset) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("parameters", Parameter.Type),
	}
}
