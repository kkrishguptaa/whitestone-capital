// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"github.com/kkrishguptaa/whitestone-capital/apps/api/ent/stake"
	"github.com/kkrishguptaa/whitestone-capital/apps/api/ent/stock"
	"github.com/kkrishguptaa/whitestone-capital/apps/api/ent/user"
)

// Stake is the model entity for the Stake schema.
type Stake struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Amount of stake in the stock
	Amount int `json:"amount,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the StakeQuery when eager-loading is set.
	Edges        StakeEdges `json:"edges"`
	stock_stakes *int
	user_stakes  *uuid.UUID
	selectValues sql.SelectValues
}

// StakeEdges holds the relations/edges for other nodes in the graph.
type StakeEdges struct {
	// The user who owns the stake in the stock
	User *User `json:"user,omitempty"`
	// The stock in which the stake is held
	Stock *Stock `json:"stock,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// UserOrErr returns the User value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e StakeEdges) UserOrErr() (*User, error) {
	if e.User != nil {
		return e.User, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: user.Label}
	}
	return nil, &NotLoadedError{edge: "user"}
}

// StockOrErr returns the Stock value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e StakeEdges) StockOrErr() (*Stock, error) {
	if e.Stock != nil {
		return e.Stock, nil
	} else if e.loadedTypes[1] {
		return nil, &NotFoundError{label: stock.Label}
	}
	return nil, &NotLoadedError{edge: "stock"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Stake) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case stake.FieldID, stake.FieldAmount:
			values[i] = new(sql.NullInt64)
		case stake.ForeignKeys[0]: // stock_stakes
			values[i] = new(sql.NullInt64)
		case stake.ForeignKeys[1]: // user_stakes
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Stake fields.
func (s *Stake) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case stake.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			s.ID = int(value.Int64)
		case stake.FieldAmount:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field amount", values[i])
			} else if value.Valid {
				s.Amount = int(value.Int64)
			}
		case stake.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field stock_stakes", value)
			} else if value.Valid {
				s.stock_stakes = new(int)
				*s.stock_stakes = int(value.Int64)
			}
		case stake.ForeignKeys[1]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field user_stakes", values[i])
			} else if value.Valid {
				s.user_stakes = new(uuid.UUID)
				*s.user_stakes = *value.S.(*uuid.UUID)
			}
		default:
			s.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Stake.
// This includes values selected through modifiers, order, etc.
func (s *Stake) Value(name string) (ent.Value, error) {
	return s.selectValues.Get(name)
}

// QueryUser queries the "user" edge of the Stake entity.
func (s *Stake) QueryUser() *UserQuery {
	return NewStakeClient(s.config).QueryUser(s)
}

// QueryStock queries the "stock" edge of the Stake entity.
func (s *Stake) QueryStock() *StockQuery {
	return NewStakeClient(s.config).QueryStock(s)
}

// Update returns a builder for updating this Stake.
// Note that you need to call Stake.Unwrap() before calling this method if this Stake
// was returned from a transaction, and the transaction was committed or rolled back.
func (s *Stake) Update() *StakeUpdateOne {
	return NewStakeClient(s.config).UpdateOne(s)
}

// Unwrap unwraps the Stake entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (s *Stake) Unwrap() *Stake {
	_tx, ok := s.config.driver.(*txDriver)
	if !ok {
		panic("ent: Stake is not a transactional entity")
	}
	s.config.driver = _tx.drv
	return s
}

// String implements the fmt.Stringer.
func (s *Stake) String() string {
	var builder strings.Builder
	builder.WriteString("Stake(")
	builder.WriteString(fmt.Sprintf("id=%v, ", s.ID))
	builder.WriteString("amount=")
	builder.WriteString(fmt.Sprintf("%v", s.Amount))
	builder.WriteByte(')')
	return builder.String()
}

// Stakes is a parsable slice of Stake.
type Stakes []*Stake
