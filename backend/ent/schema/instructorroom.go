package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// InstructorRoom holds the schema definition for the InstructorRoom entity.
type InstructorRoom struct {
	ent.Schema
}

// Fields of the InstructorRoom.
func (InstructorRoom) Fields() []ent.Field {
	return []ent.Field{
		field.String("ROOM").Unique(),
		field.String("BUILDING"),
	}
}

// Edges of the InstructorRoom.
func (InstructorRoom) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("instructorinfos", InstructorInfo.Type).StorageKey(edge.Column("instructorroom_id")),
	}
}
