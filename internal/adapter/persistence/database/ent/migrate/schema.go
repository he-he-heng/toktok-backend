// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// AvatarsColumns holds the columns for the "avatars" table.
	AvatarsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
		{Name: "sex", Type: field.TypeEnum, Enums: []string{"male", "female"}},
		{Name: "birthday", Type: field.TypeString, Size: 8},
		{Name: "mbti", Type: field.TypeString, Nullable: true, Size: 4},
		{Name: "picture", Type: field.TypeString},
		{Name: "nickname", Type: field.TypeString},
		{Name: "introduce", Type: field.TypeString, Nullable: true},
		{Name: "state", Type: field.TypeEnum, Enums: []string{"online", "offline"}, Default: "online"},
		{Name: "user_avatar", Type: field.TypeInt, Unique: true},
	}
	// AvatarsTable holds the schema information for the "avatars" table.
	AvatarsTable = &schema.Table{
		Name:       "avatars",
		Columns:    AvatarsColumns,
		PrimaryKey: []*schema.Column{AvatarsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "avatars_users_avatar",
				Columns:    []*schema.Column{AvatarsColumns[11]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// MessagesColumns holds the columns for the "messages" table.
	MessagesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
		{Name: "state", Type: field.TypeEnum, Enums: []string{"check", "uncheck"}},
		{Name: "content", Type: field.TypeString},
		{Name: "entered_at", Type: field.TypeTime},
		{Name: "avatar_messages", Type: field.TypeInt, Nullable: true},
		{Name: "relation_messages", Type: field.TypeInt, Nullable: true},
	}
	// MessagesTable holds the schema information for the "messages" table.
	MessagesTable = &schema.Table{
		Name:       "messages",
		Columns:    MessagesColumns,
		PrimaryKey: []*schema.Column{MessagesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "messages_avatars_messages",
				Columns:    []*schema.Column{MessagesColumns[7]},
				RefColumns: []*schema.Column{AvatarsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "messages_relations_messages",
				Columns:    []*schema.Column{MessagesColumns[8]},
				RefColumns: []*schema.Column{RelationsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// RelationsColumns holds the columns for the "relations" table.
	RelationsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
		{Name: "state", Type: field.TypeEnum, Enums: []string{"accepted", "pending", "declined", "removed"}, Default: "pending"},
		{Name: "alter_state", Type: field.TypeEnum, Enums: []string{"allow", "deny"}, Default: "allow"},
		{Name: "avatar_avatar_relations", Type: field.TypeInt, Nullable: true},
		{Name: "avatar_friend_relations", Type: field.TypeInt, Nullable: true},
	}
	// RelationsTable holds the schema information for the "relations" table.
	RelationsTable = &schema.Table{
		Name:       "relations",
		Columns:    RelationsColumns,
		PrimaryKey: []*schema.Column{RelationsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "relations_avatars_avatarRelations",
				Columns:    []*schema.Column{RelationsColumns[6]},
				RefColumns: []*schema.Column{AvatarsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "relations_avatars_friendRelations",
				Columns:    []*schema.Column{RelationsColumns[7]},
				RefColumns: []*schema.Column{AvatarsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
		{Name: "uid", Type: field.TypeString, Size: 18},
		{Name: "password", Type: field.TypeString, Size: 32},
		{Name: "role", Type: field.TypeEnum, Enums: []string{"admin", "user"}, Default: "user"},
		{Name: "is_ban", Type: field.TypeBool, Default: false},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		AvatarsTable,
		MessagesTable,
		RelationsTable,
		UsersTable,
	}
)

func init() {
	AvatarsTable.ForeignKeys[0].RefTable = UsersTable
	MessagesTable.ForeignKeys[0].RefTable = AvatarsTable
	MessagesTable.ForeignKeys[1].RefTable = RelationsTable
	RelationsTable.ForeignKeys[0].RefTable = AvatarsTable
	RelationsTable.ForeignKeys[1].RefTable = AvatarsTable
}
