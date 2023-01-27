package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

type Parameter struct {
	ent.Schema
}

func (Parameter) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New),
		field.String("name"),
		field.String("unit"),
		field.UUID("dataset_id", uuid.UUID{}).
			Optional(),
		field.Time("created_at").
			Default(time.Now),
	}
}

func (Parameter) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("datasets", Dataset.Type).
			Ref("parameters").
			Field("dataset_id").
			Unique(),
		edge.To("datapoints", Datapoint.Type),
	}
}
