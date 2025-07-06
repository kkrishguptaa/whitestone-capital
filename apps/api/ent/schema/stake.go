package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Stake holds the schema definition for the Stake entity.
type Stake struct {
	ent.Schema
}

// Fields of the Stake.
func (Stake) Fields() []ent.Field {
	return []ent.Field{
		field.Int("amount").Positive().Comment("Amount of stake in the stock"),
	}
}

// Edges of the Stake.
func (Stake) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).
			Ref("stakes").
			Unique().
			Required().
			Comment("The user who owns the stake in the stock"),
		edge.From("stock", Stock.Type).
			Ref("stakes").
			Unique().
			Required().
			Comment("The stock in which the stake is held"),
	}
}
