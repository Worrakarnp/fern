package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// Course holds the schema definition for the Course entity.
type Course struct {
	ent.Schema
}

// Fields of the Course.
func (Course) Fields() []ent.Field {
	return []ent.Field{

		field.String("Course_name").NotEmpty(),
	}
}

// Edges of the Course.
func (Course) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("InstructorInfo_id", InstructorInfo.Type).
			Ref("instructor").Unique(),

		edge.From("Department_id", Department.Type).
			Ref("department").Unique(),

		edge.From("Degree_id", Degree.Type).
			Ref("degree").Unique(),

		edge.From("Subject_id", Subject.Type).
			Ref("subject").Unique(),

		edge.From("Petition_id", Petition.Type).
			Ref("petition").Unique(),
	}
}
