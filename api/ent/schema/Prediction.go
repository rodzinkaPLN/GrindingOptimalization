package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

type Prediction struct {
	ent.Schema
}

func (Prediction) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New),
		field.Time("data_time"),
		field.JSON("vals", DataT{}),
		field.Time("created_at").
			Default(time.Now),
	}
}

func (Prediction) Edges() []ent.Edge {
	return []ent.Edge{}
}
