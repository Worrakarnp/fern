package schema

import "github.com/facebookincubator/ent"

// Academic holds the schema definition for the Academic entity.
type Academic struct {
	ent.Schema
}

// Fields of the Academic.
func (Academic) Fields() []ent.Field {
	return nil
}

// Edges of the Academic.
func (Academic) Edges() []ent.Edge {
	return nil
}
