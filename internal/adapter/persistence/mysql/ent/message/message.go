// Code generated by ent, DO NOT EDIT.

package message

import (
	"fmt"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the message type in the database.
	Label = "message"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldDeletedAt holds the string denoting the deleted_at field in the database.
	FieldDeletedAt = "deleted_at"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldState holds the string denoting the state field in the database.
	FieldState = "state"
	// FieldContent holds the string denoting the content field in the database.
	FieldContent = "content"
	// FieldEnteredAt holds the string denoting the entered_at field in the database.
	FieldEnteredAt = "entered_at"
	// EdgeAvatar holds the string denoting the avatar edge name in mutations.
	EdgeAvatar = "avatar"
	// EdgeRoom holds the string denoting the room edge name in mutations.
	EdgeRoom = "room"
	// Table holds the table name of the message in the database.
	Table = "messages"
	// AvatarTable is the table that holds the avatar relation/edge.
	AvatarTable = "messages"
	// AvatarInverseTable is the table name for the Avatar entity.
	// It exists in this package in order to avoid circular dependency with the "avatar" package.
	AvatarInverseTable = "avatars"
	// AvatarColumn is the table column denoting the avatar relation/edge.
	AvatarColumn = "avatar_messages"
	// RoomTable is the table that holds the room relation/edge.
	RoomTable = "messages"
	// RoomInverseTable is the table name for the Room entity.
	// It exists in this package in order to avoid circular dependency with the "room" package.
	RoomInverseTable = "rooms"
	// RoomColumn is the table column denoting the room relation/edge.
	RoomColumn = "room_messages"
)

// Columns holds all SQL columns for message fields.
var Columns = []string{
	FieldID,
	FieldDeletedAt,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldState,
	FieldContent,
	FieldEnteredAt,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "messages"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"avatar_messages",
	"room_messages",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

// Note that the variables below are initialized by the runtime
// package on the initialization of the application. Therefore,
// it should be imported in the main as follows:
//
//	import _ "toktok-backend/internal/adapter/persistence/mysql/ent/runtime"
var (
	Hooks        [1]ent.Hook
	Interceptors [1]ent.Interceptor
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
	// ContentValidator is a validator for the "content" field. It is called by the builders before save.
	ContentValidator func(string) error
	// DefaultEnteredAt holds the default value on creation for the "entered_at" field.
	DefaultEnteredAt func() time.Time
)

// State defines the type for the "state" enum field.
type State string

// State values.
const (
	StateCheck   State = "check"
	StateUncheck State = "uncheck"
)

func (s State) String() string {
	return string(s)
}

// StateValidator is a validator for the "state" field enum values. It is called by the builders before save.
func StateValidator(s State) error {
	switch s {
	case StateCheck, StateUncheck:
		return nil
	default:
		return fmt.Errorf("message: invalid enum value for state field: %q", s)
	}
}

// OrderOption defines the ordering options for the Message queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByDeletedAt orders the results by the deleted_at field.
func ByDeletedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDeletedAt, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByUpdatedAt orders the results by the updated_at field.
func ByUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedAt, opts...).ToFunc()
}

// ByState orders the results by the state field.
func ByState(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldState, opts...).ToFunc()
}

// ByContent orders the results by the content field.
func ByContent(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldContent, opts...).ToFunc()
}

// ByEnteredAt orders the results by the entered_at field.
func ByEnteredAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldEnteredAt, opts...).ToFunc()
}

// ByAvatarField orders the results by avatar field.
func ByAvatarField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newAvatarStep(), sql.OrderByField(field, opts...))
	}
}

// ByRoomField orders the results by room field.
func ByRoomField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newRoomStep(), sql.OrderByField(field, opts...))
	}
}
func newAvatarStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(AvatarInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, AvatarTable, AvatarColumn),
	)
}
func newRoomStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(RoomInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2O, true, RoomTable, RoomColumn),
	)
}
