// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"
	"toktok-backend/internal/adapter/persistence/database/ent/avatar"
	"toktok-backend/internal/adapter/persistence/database/ent/relation"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// Relation is the model entity for the Relation schema.
type Relation struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// DeletedAt holds the value of the "deleted_at" field.
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
	// State holds the value of the "state" field.
	State relation.State `json:"state,omitempty"`
	// AlterState holds the value of the "alterState" field.
	AlterState relation.AlterState `json:"alterState,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the RelationQuery when eager-loading is set.
	Edges                   RelationEdges `json:"edges"`
	avatar_avatar_relations *int
	avatar_friend_relations *int
	selectValues            sql.SelectValues
}

// RelationEdges holds the relations/edges for other nodes in the graph.
type RelationEdges struct {
	// Avatar holds the value of the avatar edge.
	Avatar *Avatar `json:"avatar,omitempty"`
	// Friend holds the value of the friend edge.
	Friend *Avatar `json:"friend,omitempty"`
	// Messages holds the value of the messages edge.
	Messages []*Message `json:"messages,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
}

// AvatarOrErr returns the Avatar value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e RelationEdges) AvatarOrErr() (*Avatar, error) {
	if e.Avatar != nil {
		return e.Avatar, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: avatar.Label}
	}
	return nil, &NotLoadedError{edge: "avatar"}
}

// FriendOrErr returns the Friend value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e RelationEdges) FriendOrErr() (*Avatar, error) {
	if e.Friend != nil {
		return e.Friend, nil
	} else if e.loadedTypes[1] {
		return nil, &NotFoundError{label: avatar.Label}
	}
	return nil, &NotLoadedError{edge: "friend"}
}

// MessagesOrErr returns the Messages value or an error if the edge
// was not loaded in eager-loading.
func (e RelationEdges) MessagesOrErr() ([]*Message, error) {
	if e.loadedTypes[2] {
		return e.Messages, nil
	}
	return nil, &NotLoadedError{edge: "messages"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Relation) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case relation.FieldID:
			values[i] = new(sql.NullInt64)
		case relation.FieldState, relation.FieldAlterState:
			values[i] = new(sql.NullString)
		case relation.FieldCreatedAt, relation.FieldUpdatedAt, relation.FieldDeletedAt:
			values[i] = new(sql.NullTime)
		case relation.ForeignKeys[0]: // avatar_avatar_relations
			values[i] = new(sql.NullInt64)
		case relation.ForeignKeys[1]: // avatar_friend_relations
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Relation fields.
func (r *Relation) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case relation.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			r.ID = int(value.Int64)
		case relation.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				r.CreatedAt = value.Time
			}
		case relation.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				r.UpdatedAt = value.Time
			}
		case relation.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				r.DeletedAt = new(time.Time)
				*r.DeletedAt = value.Time
			}
		case relation.FieldState:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field state", values[i])
			} else if value.Valid {
				r.State = relation.State(value.String)
			}
		case relation.FieldAlterState:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field alterState", values[i])
			} else if value.Valid {
				r.AlterState = relation.AlterState(value.String)
			}
		case relation.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field avatar_avatar_relations", value)
			} else if value.Valid {
				r.avatar_avatar_relations = new(int)
				*r.avatar_avatar_relations = int(value.Int64)
			}
		case relation.ForeignKeys[1]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field avatar_friend_relations", value)
			} else if value.Valid {
				r.avatar_friend_relations = new(int)
				*r.avatar_friend_relations = int(value.Int64)
			}
		default:
			r.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Relation.
// This includes values selected through modifiers, order, etc.
func (r *Relation) Value(name string) (ent.Value, error) {
	return r.selectValues.Get(name)
}

// QueryAvatar queries the "avatar" edge of the Relation entity.
func (r *Relation) QueryAvatar() *AvatarQuery {
	return NewRelationClient(r.config).QueryAvatar(r)
}

// QueryFriend queries the "friend" edge of the Relation entity.
func (r *Relation) QueryFriend() *AvatarQuery {
	return NewRelationClient(r.config).QueryFriend(r)
}

// QueryMessages queries the "messages" edge of the Relation entity.
func (r *Relation) QueryMessages() *MessageQuery {
	return NewRelationClient(r.config).QueryMessages(r)
}

// Update returns a builder for updating this Relation.
// Note that you need to call Relation.Unwrap() before calling this method if this Relation
// was returned from a transaction, and the transaction was committed or rolled back.
func (r *Relation) Update() *RelationUpdateOne {
	return NewRelationClient(r.config).UpdateOne(r)
}

// Unwrap unwraps the Relation entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (r *Relation) Unwrap() *Relation {
	_tx, ok := r.config.driver.(*txDriver)
	if !ok {
		panic("ent: Relation is not a transactional entity")
	}
	r.config.driver = _tx.drv
	return r
}

// String implements the fmt.Stringer.
func (r *Relation) String() string {
	var builder strings.Builder
	builder.WriteString("Relation(")
	builder.WriteString(fmt.Sprintf("id=%v, ", r.ID))
	builder.WriteString("created_at=")
	builder.WriteString(r.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(r.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	if v := r.DeletedAt; v != nil {
		builder.WriteString("deleted_at=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", ")
	builder.WriteString("state=")
	builder.WriteString(fmt.Sprintf("%v", r.State))
	builder.WriteString(", ")
	builder.WriteString("alterState=")
	builder.WriteString(fmt.Sprintf("%v", r.AlterState))
	builder.WriteByte(')')
	return builder.String()
}

// Relations is a parsable slice of Relation.
type Relations []*Relation
