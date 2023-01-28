package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

type Userinput struct {
	ent.Schema
}

func (Userinput) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New),
		field.Float("min"),
		field.Float("max"),
		field.Float("step"),
		field.Float("defaultval"),
		field.String("name"),
	}
}

func (Userinput) Edges() []ent.Edge {
	return []ent.Edge{}
}
