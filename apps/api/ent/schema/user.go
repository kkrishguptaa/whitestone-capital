package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.New()).Comment("Unique identifier for the user, generated as a UUID"),
		field.String("slack_id").Unique().Immutable().NotEmpty().Comment("Slack user ID, unique and immutable"),
		field.String("name").NotEmpty().Comment("Full name of the user, e.g., 'John Doe'"),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("stakes", Stake.Type).
			Comment("The stakes held by the user in various stocks"),
	}
}
