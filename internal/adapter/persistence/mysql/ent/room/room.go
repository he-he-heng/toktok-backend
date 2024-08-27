// Code generated by ent, DO NOT EDIT.

package room

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the room type in the database.
	Label = "room"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldDeletedAt holds the string denoting the deleted_at field in the database.
	FieldDeletedAt = "deleted_at"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// EdgeAvatar holds the string denoting the avatar edge name in mutations.
	EdgeAvatar = "avatar"
	// EdgeFriend holds the string denoting the friend edge name in mutations.
	EdgeFriend = "friend"
	// EdgeMessages holds the string denoting the messages edge name in mutations.
	EdgeMessages = "messages"
	// Table holds the table name of the room in the database.
	Table = "rooms"
	// AvatarTable is the table that holds the avatar relation/edge.
	AvatarTable = "rooms"
	// AvatarInverseTable is the table name for the Relation entity.
	// It exists in this package in order to avoid circular dependency with the "relation" package.
	AvatarInverseTable = "relations"
	// AvatarColumn is the table column denoting the avatar relation/edge.
	AvatarColumn = "relation_avatar_rooms"
	// FriendTable is the table that holds the friend relation/edge.
	FriendTable = "rooms"
	// FriendInverseTable is the table name for the Relation entity.
	// It exists in this package in order to avoid circular dependency with the "relation" package.
	FriendInverseTable = "relations"
	// FriendColumn is the table column denoting the friend relation/edge.
	FriendColumn = "relation_friend_rooms"
	// MessagesTable is the table that holds the messages relation/edge.
	MessagesTable = "messages"
	// MessagesInverseTable is the table name for the Message entity.
	// It exists in this package in order to avoid circular dependency with the "message" package.
	MessagesInverseTable = "messages"
	// MessagesColumn is the table column denoting the messages relation/edge.
	MessagesColumn = "room_messages"
)

// Columns holds all SQL columns for room fields.
var Columns = []string{
	FieldID,
	FieldDeletedAt,
	FieldCreatedAt,
	FieldUpdatedAt,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "rooms"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"relation_avatar_rooms",
	"relation_friend_rooms",
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
)

// OrderOption defines the ordering options for the Room queries.
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

// ByAvatarField orders the results by avatar field.
func ByAvatarField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newAvatarStep(), sql.OrderByField(field, opts...))
	}
}

// ByFriendField orders the results by friend field.
func ByFriendField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newFriendStep(), sql.OrderByField(field, opts...))
	}
}

// ByMessagesField orders the results by messages field.
func ByMessagesField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newMessagesStep(), sql.OrderByField(field, opts...))
	}
}
func newAvatarStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(AvatarInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2O, true, AvatarTable, AvatarColumn),
	)
}
func newFriendStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(FriendInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2O, true, FriendTable, FriendColumn),
	)
}
func newMessagesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(MessagesInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2O, false, MessagesTable, MessagesColumn),
	)
}
