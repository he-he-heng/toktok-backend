// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/he-he-heng/toktok-backend/ent/attempt"
)

// Attempt is the model entity for the Attempt schema.
type Attempt struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Cnt holds the value of the "cnt" field.
	Cnt int `json:"cnt,omitempty"`
	// Break holds the value of the "break" field.
	Break bool `json:"break,omitempty"`
	// Timestamp holds the value of the "timestamp" field.
	Timestamp    time.Time `json:"timestamp,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Attempt) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case attempt.FieldBreak:
			values[i] = new(sql.NullBool)
		case attempt.FieldID, attempt.FieldCnt:
			values[i] = new(sql.NullInt64)
		case attempt.FieldTimestamp:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Attempt fields.
func (a *Attempt) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case attempt.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			a.ID = int(value.Int64)
		case attempt.FieldCnt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field cnt", values[i])
			} else if value.Valid {
				a.Cnt = int(value.Int64)
			}
		case attempt.FieldBreak:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field break", values[i])
			} else if value.Valid {
				a.Break = value.Bool
			}
		case attempt.FieldTimestamp:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field timestamp", values[i])
			} else if value.Valid {
				a.Timestamp = value.Time
			}
		default:
			a.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Attempt.
// This includes values selected through modifiers, order, etc.
func (a *Attempt) Value(name string) (ent.Value, error) {
	return a.selectValues.Get(name)
}

// Update returns a builder for updating this Attempt.
// Note that you need to call Attempt.Unwrap() before calling this method if this Attempt
// was returned from a transaction, and the transaction was committed or rolled back.
func (a *Attempt) Update() *AttemptUpdateOne {
	return NewAttemptClient(a.config).UpdateOne(a)
}

// Unwrap unwraps the Attempt entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (a *Attempt) Unwrap() *Attempt {
	_tx, ok := a.config.driver.(*txDriver)
	if !ok {
		panic("ent: Attempt is not a transactional entity")
	}
	a.config.driver = _tx.drv
	return a
}

// String implements the fmt.Stringer.
func (a *Attempt) String() string {
	var builder strings.Builder
	builder.WriteString("Attempt(")
	builder.WriteString(fmt.Sprintf("id=%v, ", a.ID))
	builder.WriteString("cnt=")
	builder.WriteString(fmt.Sprintf("%v", a.Cnt))
	builder.WriteString(", ")
	builder.WriteString("break=")
	builder.WriteString(fmt.Sprintf("%v", a.Break))
	builder.WriteString(", ")
	builder.WriteString("timestamp=")
	builder.WriteString(a.Timestamp.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Attempts is a parsable slice of Attempt.
type Attempts []*Attempt
