package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Stock holds the schema definition for the Stock entity.
type Stock struct {
	ent.Schema
}

type Activity struct {
	Date          string `json:"date"`
	MessagesCount int    `json:"messages_count"`
	UsersCount    int    `json:"users_count"`
}

// Fields of the Stock.
func (Stock) Fields() []ent.Field {
	return []ent.Field{
		field.String("channel_id").Unique().Immutable().NotEmpty().Comment("Slack channel ID for the stock."),
		field.String("symbol").Unique().NotEmpty().Comment("Stock symbol, e.g., 'AAPL' for Apple Inc."),
		field.String("name").NotEmpty().Comment("Full name of the stock, e.g., 'Apple Inc.'"),
		field.JSON("activity", &Activity{}).Optional().Comment("Activity data for the stock, including date, messages count, and users count"),
	}
}

// Edges of the Stock.
func (Stock) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("stakes", Stake.Type).
			Comment("The stakes held in the stock by users"),
	}
}
