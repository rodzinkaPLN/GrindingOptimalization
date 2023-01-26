package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

type Datapoint struct {
	ent.Schema
}

func (Datapoint) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New),
		field.Time("data_time"),
		field.UUID("parameter_id", uuid.UUID{}).
			Optional(),
		field.Float("val"),
		field.Time("created_at").
			Default(time.Now),
	}
}

func (Datapoint) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("parameters", Parameter.Type).
			Ref("datapoints").
			Field("parameter_id").
			Unique(),
	}
}
