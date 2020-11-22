package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// Personal holds the schema definition for the Personal entity.
type Personal struct {
	ent.Schema
}

// Fields of the Personal.
func (Personal) Fields() []ent.Field {
	return []ent.Field{
		field.String("PersonalName"),
	}
}

// Edges of the Personal.
func (Personal) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("jobtitle", Jobtitle.Type).Ref("personal").Unique(),
		edge.From("department", Department.Type).Ref("personal").Unique(),
		edge.From("gender", Gender.Type).Ref("personal").Unique(),
	}
}
