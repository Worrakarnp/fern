package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// Degree holds the schema definition for the Degree entity.
type Degree struct {
	ent.Schema
}

// Fields of the Degree.
func (Degree) Fields() []ent.Field {
	return []ent.Field{
		
		field.String("Degree_name").Unique(),
	}
}

// Edges of the Degree.
func (Degree) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("degree", Course.Type).StorageKey(edge.Column("Degree_id")),
	}
}
