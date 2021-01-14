package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// InstructorInfo holds the schema definition for the InstructorInfo entity.
type InstructorInfo struct {
	ent.Schema
}

// Fields of the InstructorInfo.
func (InstructorInfo) Fields() []ent.Field {
	return []ent.Field{
		field.String("NAME").Unique(),
		field.String("PHONENUMBER").Unique(),
		field.String("EMAIL").Unique(),
		field.String("PASSWORD"),
	}
}

// Edges of the InstructorInfo.
func (InstructorInfo) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("title", Title.Type).Ref("instructorinfos").Unique(),
		edge.From("instructorroom", InstructorRoom.Type).Ref("instructorinfos").Unique(),
		edge.From("department", Department.Type).Ref("instructorinfos").Unique(),
		edge.To("instructor", Course.Type).StorageKey(edge.Column("InstructorInfo_id")),
	}
}
