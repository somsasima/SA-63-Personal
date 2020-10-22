package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// Systemmember holds the schema definition for the Systemmember entity.
type Systemmember struct {
	ent.Schema
}

// Fields of the Systemmember.
func (Systemmember) Fields() []ent.Field {
	return []ent.Field{
		field.String("Mail").Unique(),
		field.String("Password"),
	}
}

// Edges of the Systemmember.
func (Systemmember) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("personal", Personal.Type).StorageKey(edge.Column("systemmember_id")),
	}
}
