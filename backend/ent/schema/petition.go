package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// Petition holds the schema definition for the Petition entity.
type Petition struct {
	ent.Schema
}

// Fields of the Petition.
func (Petition) Fields() []ent.Field {
	return []ent.Field{

		field.String("Petition_name").Unique(),
	}
}

// Edges of the Petition.
func (Petition) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("petition", Course.Type).StorageKey(edge.Column("Petition_id")),
	}
}
