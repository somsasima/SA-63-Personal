package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// Jobtitle holds the schema definition for the Jobtitle entity.
type Jobtitle struct {
	ent.Schema
}

// Fields of the Jobtitle.
func (Jobtitle) Fields() []ent.Field {
	return []ent.Field{
		field.String("Jobtitlename").Unique(),
	}
}

// Edges of the Jobtitle.
func (Jobtitle) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("personal", Personal.Type).StorageKey(edge.Column("jobtitle_id")),
	}
}
