package schema

import "github.com/facebookincubator/ent"

// Request holds the schema definition for the Request entity.
type Request struct {
	ent.Schema
}

// Fields of the Request.
func (Request) Fields() []ent.Field {
	return nil
}

// Edges of the Request.
func (Request) Edges() []ent.Edge {
	return nil
}
