// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"
	"toktok-backend/internal/adapter/persistence/mysql/ent/avatar"
	"toktok-backend/internal/adapter/persistence/mysql/ent/message"
	"toktok-backend/internal/adapter/persistence/mysql/ent/predicate"
	"toktok-backend/internal/adapter/persistence/mysql/ent/relation"
	"toktok-backend/internal/adapter/persistence/mysql/ent/user"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// AvatarUpdate is the builder for updating Avatar entities.
type AvatarUpdate struct {
	config
	hooks    []Hook
	mutation *AvatarMutation
}

// Where appends a list predicates to the AvatarUpdate builder.
func (au *AvatarUpdate) Where(ps ...predicate.Avatar) *AvatarUpdate {
	au.mutation.Where(ps...)
	return au
}

// SetDeletedAt sets the "deleted_at" field.
func (au *AvatarUpdate) SetDeletedAt(t time.Time) *AvatarUpdate {
	au.mutation.SetDeletedAt(t)
	return au
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (au *AvatarUpdate) SetNillableDeletedAt(t *time.Time) *AvatarUpdate {
	if t != nil {
		au.SetDeletedAt(*t)
	}
	return au
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (au *AvatarUpdate) ClearDeletedAt() *AvatarUpdate {
	au.mutation.ClearDeletedAt()
	return au
}

// SetUpdatedAt sets the "updated_at" field.
func (au *AvatarUpdate) SetUpdatedAt(t time.Time) *AvatarUpdate {
	au.mutation.SetUpdatedAt(t)
	return au
}

// SetSex sets the "sex" field.
func (au *AvatarUpdate) SetSex(a avatar.Sex) *AvatarUpdate {
	au.mutation.SetSex(a)
	return au
}

// SetNillableSex sets the "sex" field if the given value is not nil.
func (au *AvatarUpdate) SetNillableSex(a *avatar.Sex) *AvatarUpdate {
	if a != nil {
		au.SetSex(*a)
	}
	return au
}

// SetBirthday sets the "birthday" field.
func (au *AvatarUpdate) SetBirthday(s string) *AvatarUpdate {
	au.mutation.SetBirthday(s)
	return au
}

// SetNillableBirthday sets the "birthday" field if the given value is not nil.
func (au *AvatarUpdate) SetNillableBirthday(s *string) *AvatarUpdate {
	if s != nil {
		au.SetBirthday(*s)
	}
	return au
}

// SetMbti sets the "mbti" field.
func (au *AvatarUpdate) SetMbti(s string) *AvatarUpdate {
	au.mutation.SetMbti(s)
	return au
}

// SetNillableMbti sets the "mbti" field if the given value is not nil.
func (au *AvatarUpdate) SetNillableMbti(s *string) *AvatarUpdate {
	if s != nil {
		au.SetMbti(*s)
	}
	return au
}

// ClearMbti clears the value of the "mbti" field.
func (au *AvatarUpdate) ClearMbti() *AvatarUpdate {
	au.mutation.ClearMbti()
	return au
}

// SetPicture sets the "picture" field.
func (au *AvatarUpdate) SetPicture(a avatar.Picture) *AvatarUpdate {
	au.mutation.SetPicture(a)
	return au
}

// SetNillablePicture sets the "picture" field if the given value is not nil.
func (au *AvatarUpdate) SetNillablePicture(a *avatar.Picture) *AvatarUpdate {
	if a != nil {
		au.SetPicture(*a)
	}
	return au
}

// SetNickname sets the "nickname" field.
func (au *AvatarUpdate) SetNickname(s string) *AvatarUpdate {
	au.mutation.SetNickname(s)
	return au
}

// SetNillableNickname sets the "nickname" field if the given value is not nil.
func (au *AvatarUpdate) SetNillableNickname(s *string) *AvatarUpdate {
	if s != nil {
		au.SetNickname(*s)
	}
	return au
}

// SetIntroduce sets the "introduce" field.
func (au *AvatarUpdate) SetIntroduce(s string) *AvatarUpdate {
	au.mutation.SetIntroduce(s)
	return au
}

// SetNillableIntroduce sets the "introduce" field if the given value is not nil.
func (au *AvatarUpdate) SetNillableIntroduce(s *string) *AvatarUpdate {
	if s != nil {
		au.SetIntroduce(*s)
	}
	return au
}

// ClearIntroduce clears the value of the "introduce" field.
func (au *AvatarUpdate) ClearIntroduce() *AvatarUpdate {
	au.mutation.ClearIntroduce()
	return au
}

// SetState sets the "state" field.
func (au *AvatarUpdate) SetState(a avatar.State) *AvatarUpdate {
	au.mutation.SetState(a)
	return au
}

// SetNillableState sets the "state" field if the given value is not nil.
func (au *AvatarUpdate) SetNillableState(a *avatar.State) *AvatarUpdate {
	if a != nil {
		au.SetState(*a)
	}
	return au
}

// SetUserID sets the "user" edge to the User entity by ID.
func (au *AvatarUpdate) SetUserID(id int) *AvatarUpdate {
	au.mutation.SetUserID(id)
	return au
}

// SetUser sets the "user" edge to the User entity.
func (au *AvatarUpdate) SetUser(u *User) *AvatarUpdate {
	return au.SetUserID(u.ID)
}

// AddAvatarRelationIDs adds the "avatar_relations" edge to the Relation entity by IDs.
func (au *AvatarUpdate) AddAvatarRelationIDs(ids ...int) *AvatarUpdate {
	au.mutation.AddAvatarRelationIDs(ids...)
	return au
}

// AddAvatarRelations adds the "avatar_relations" edges to the Relation entity.
func (au *AvatarUpdate) AddAvatarRelations(r ...*Relation) *AvatarUpdate {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return au.AddAvatarRelationIDs(ids...)
}

// AddFriendRelationIDs adds the "friend_relations" edge to the Relation entity by IDs.
func (au *AvatarUpdate) AddFriendRelationIDs(ids ...int) *AvatarUpdate {
	au.mutation.AddFriendRelationIDs(ids...)
	return au
}

// AddFriendRelations adds the "friend_relations" edges to the Relation entity.
func (au *AvatarUpdate) AddFriendRelations(r ...*Relation) *AvatarUpdate {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return au.AddFriendRelationIDs(ids...)
}

// AddMessageIDs adds the "messages" edge to the Message entity by IDs.
func (au *AvatarUpdate) AddMessageIDs(ids ...int) *AvatarUpdate {
	au.mutation.AddMessageIDs(ids...)
	return au
}

// AddMessages adds the "messages" edges to the Message entity.
func (au *AvatarUpdate) AddMessages(m ...*Message) *AvatarUpdate {
	ids := make([]int, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return au.AddMessageIDs(ids...)
}

// Mutation returns the AvatarMutation object of the builder.
func (au *AvatarUpdate) Mutation() *AvatarMutation {
	return au.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (au *AvatarUpdate) ClearUser() *AvatarUpdate {
	au.mutation.ClearUser()
	return au
}

// ClearAvatarRelations clears all "avatar_relations" edges to the Relation entity.
func (au *AvatarUpdate) ClearAvatarRelations() *AvatarUpdate {
	au.mutation.ClearAvatarRelations()
	return au
}

// RemoveAvatarRelationIDs removes the "avatar_relations" edge to Relation entities by IDs.
func (au *AvatarUpdate) RemoveAvatarRelationIDs(ids ...int) *AvatarUpdate {
	au.mutation.RemoveAvatarRelationIDs(ids...)
	return au
}

// RemoveAvatarRelations removes "avatar_relations" edges to Relation entities.
func (au *AvatarUpdate) RemoveAvatarRelations(r ...*Relation) *AvatarUpdate {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return au.RemoveAvatarRelationIDs(ids...)
}

// ClearFriendRelations clears all "friend_relations" edges to the Relation entity.
func (au *AvatarUpdate) ClearFriendRelations() *AvatarUpdate {
	au.mutation.ClearFriendRelations()
	return au
}

// RemoveFriendRelationIDs removes the "friend_relations" edge to Relation entities by IDs.
func (au *AvatarUpdate) RemoveFriendRelationIDs(ids ...int) *AvatarUpdate {
	au.mutation.RemoveFriendRelationIDs(ids...)
	return au
}

// RemoveFriendRelations removes "friend_relations" edges to Relation entities.
func (au *AvatarUpdate) RemoveFriendRelations(r ...*Relation) *AvatarUpdate {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return au.RemoveFriendRelationIDs(ids...)
}

// ClearMessages clears all "messages" edges to the Message entity.
func (au *AvatarUpdate) ClearMessages() *AvatarUpdate {
	au.mutation.ClearMessages()
	return au
}

// RemoveMessageIDs removes the "messages" edge to Message entities by IDs.
func (au *AvatarUpdate) RemoveMessageIDs(ids ...int) *AvatarUpdate {
	au.mutation.RemoveMessageIDs(ids...)
	return au
}

// RemoveMessages removes "messages" edges to Message entities.
func (au *AvatarUpdate) RemoveMessages(m ...*Message) *AvatarUpdate {
	ids := make([]int, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return au.RemoveMessageIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (au *AvatarUpdate) Save(ctx context.Context) (int, error) {
	if err := au.defaults(); err != nil {
		return 0, err
	}
	return withHooks(ctx, au.sqlSave, au.mutation, au.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (au *AvatarUpdate) SaveX(ctx context.Context) int {
	affected, err := au.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (au *AvatarUpdate) Exec(ctx context.Context) error {
	_, err := au.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (au *AvatarUpdate) ExecX(ctx context.Context) {
	if err := au.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (au *AvatarUpdate) defaults() error {
	if _, ok := au.mutation.UpdatedAt(); !ok {
		if avatar.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized avatar.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := avatar.UpdateDefaultUpdatedAt()
		au.mutation.SetUpdatedAt(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (au *AvatarUpdate) check() error {
	if v, ok := au.mutation.Sex(); ok {
		if err := avatar.SexValidator(v); err != nil {
			return &ValidationError{Name: "sex", err: fmt.Errorf(`ent: validator failed for field "Avatar.sex": %w`, err)}
		}
	}
	if v, ok := au.mutation.Birthday(); ok {
		if err := avatar.BirthdayValidator(v); err != nil {
			return &ValidationError{Name: "birthday", err: fmt.Errorf(`ent: validator failed for field "Avatar.birthday": %w`, err)}
		}
	}
	if v, ok := au.mutation.Mbti(); ok {
		if err := avatar.MbtiValidator(v); err != nil {
			return &ValidationError{Name: "mbti", err: fmt.Errorf(`ent: validator failed for field "Avatar.mbti": %w`, err)}
		}
	}
	if v, ok := au.mutation.Picture(); ok {
		if err := avatar.PictureValidator(v); err != nil {
			return &ValidationError{Name: "picture", err: fmt.Errorf(`ent: validator failed for field "Avatar.picture": %w`, err)}
		}
	}
	if v, ok := au.mutation.Nickname(); ok {
		if err := avatar.NicknameValidator(v); err != nil {
			return &ValidationError{Name: "nickname", err: fmt.Errorf(`ent: validator failed for field "Avatar.nickname": %w`, err)}
		}
	}
	if v, ok := au.mutation.State(); ok {
		if err := avatar.StateValidator(v); err != nil {
			return &ValidationError{Name: "state", err: fmt.Errorf(`ent: validator failed for field "Avatar.state": %w`, err)}
		}
	}
	if au.mutation.UserCleared() && len(au.mutation.UserIDs()) > 0 {
		return errors.New(`ent: clearing a required unique edge "Avatar.user"`)
	}
	return nil
}

func (au *AvatarUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := au.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(avatar.Table, avatar.Columns, sqlgraph.NewFieldSpec(avatar.FieldID, field.TypeInt))
	if ps := au.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := au.mutation.DeletedAt(); ok {
		_spec.SetField(avatar.FieldDeletedAt, field.TypeTime, value)
	}
	if au.mutation.DeletedAtCleared() {
		_spec.ClearField(avatar.FieldDeletedAt, field.TypeTime)
	}
	if value, ok := au.mutation.UpdatedAt(); ok {
		_spec.SetField(avatar.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := au.mutation.Sex(); ok {
		_spec.SetField(avatar.FieldSex, field.TypeEnum, value)
	}
	if value, ok := au.mutation.Birthday(); ok {
		_spec.SetField(avatar.FieldBirthday, field.TypeString, value)
	}
	if value, ok := au.mutation.Mbti(); ok {
		_spec.SetField(avatar.FieldMbti, field.TypeString, value)
	}
	if au.mutation.MbtiCleared() {
		_spec.ClearField(avatar.FieldMbti, field.TypeString)
	}
	if value, ok := au.mutation.Picture(); ok {
		_spec.SetField(avatar.FieldPicture, field.TypeEnum, value)
	}
	if value, ok := au.mutation.Nickname(); ok {
		_spec.SetField(avatar.FieldNickname, field.TypeString, value)
	}
	if value, ok := au.mutation.Introduce(); ok {
		_spec.SetField(avatar.FieldIntroduce, field.TypeString, value)
	}
	if au.mutation.IntroduceCleared() {
		_spec.ClearField(avatar.FieldIntroduce, field.TypeString)
	}
	if value, ok := au.mutation.State(); ok {
		_spec.SetField(avatar.FieldState, field.TypeEnum, value)
	}
	if au.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   avatar.UserTable,
			Columns: []string{avatar.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := au.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   avatar.UserTable,
			Columns: []string{avatar.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if au.mutation.AvatarRelationsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   avatar.AvatarRelationsTable,
			Columns: []string{avatar.AvatarRelationsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(relation.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := au.mutation.RemovedAvatarRelationsIDs(); len(nodes) > 0 && !au.mutation.AvatarRelationsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   avatar.AvatarRelationsTable,
			Columns: []string{avatar.AvatarRelationsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(relation.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := au.mutation.AvatarRelationsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   avatar.AvatarRelationsTable,
			Columns: []string{avatar.AvatarRelationsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(relation.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if au.mutation.FriendRelationsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   avatar.FriendRelationsTable,
			Columns: []string{avatar.FriendRelationsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(relation.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := au.mutation.RemovedFriendRelationsIDs(); len(nodes) > 0 && !au.mutation.FriendRelationsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   avatar.FriendRelationsTable,
			Columns: []string{avatar.FriendRelationsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(relation.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := au.mutation.FriendRelationsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   avatar.FriendRelationsTable,
			Columns: []string{avatar.FriendRelationsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(relation.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if au.mutation.MessagesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   avatar.MessagesTable,
			Columns: []string{avatar.MessagesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(message.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := au.mutation.RemovedMessagesIDs(); len(nodes) > 0 && !au.mutation.MessagesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   avatar.MessagesTable,
			Columns: []string{avatar.MessagesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(message.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := au.mutation.MessagesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   avatar.MessagesTable,
			Columns: []string{avatar.MessagesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(message.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, au.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{avatar.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	au.mutation.done = true
	return n, nil
}

// AvatarUpdateOne is the builder for updating a single Avatar entity.
type AvatarUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *AvatarMutation
}

// SetDeletedAt sets the "deleted_at" field.
func (auo *AvatarUpdateOne) SetDeletedAt(t time.Time) *AvatarUpdateOne {
	auo.mutation.SetDeletedAt(t)
	return auo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (auo *AvatarUpdateOne) SetNillableDeletedAt(t *time.Time) *AvatarUpdateOne {
	if t != nil {
		auo.SetDeletedAt(*t)
	}
	return auo
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (auo *AvatarUpdateOne) ClearDeletedAt() *AvatarUpdateOne {
	auo.mutation.ClearDeletedAt()
	return auo
}

// SetUpdatedAt sets the "updated_at" field.
func (auo *AvatarUpdateOne) SetUpdatedAt(t time.Time) *AvatarUpdateOne {
	auo.mutation.SetUpdatedAt(t)
	return auo
}

// SetSex sets the "sex" field.
func (auo *AvatarUpdateOne) SetSex(a avatar.Sex) *AvatarUpdateOne {
	auo.mutation.SetSex(a)
	return auo
}

// SetNillableSex sets the "sex" field if the given value is not nil.
func (auo *AvatarUpdateOne) SetNillableSex(a *avatar.Sex) *AvatarUpdateOne {
	if a != nil {
		auo.SetSex(*a)
	}
	return auo
}

// SetBirthday sets the "birthday" field.
func (auo *AvatarUpdateOne) SetBirthday(s string) *AvatarUpdateOne {
	auo.mutation.SetBirthday(s)
	return auo
}

// SetNillableBirthday sets the "birthday" field if the given value is not nil.
func (auo *AvatarUpdateOne) SetNillableBirthday(s *string) *AvatarUpdateOne {
	if s != nil {
		auo.SetBirthday(*s)
	}
	return auo
}

// SetMbti sets the "mbti" field.
func (auo *AvatarUpdateOne) SetMbti(s string) *AvatarUpdateOne {
	auo.mutation.SetMbti(s)
	return auo
}

// SetNillableMbti sets the "mbti" field if the given value is not nil.
func (auo *AvatarUpdateOne) SetNillableMbti(s *string) *AvatarUpdateOne {
	if s != nil {
		auo.SetMbti(*s)
	}
	return auo
}

// ClearMbti clears the value of the "mbti" field.
func (auo *AvatarUpdateOne) ClearMbti() *AvatarUpdateOne {
	auo.mutation.ClearMbti()
	return auo
}

// SetPicture sets the "picture" field.
func (auo *AvatarUpdateOne) SetPicture(a avatar.Picture) *AvatarUpdateOne {
	auo.mutation.SetPicture(a)
	return auo
}

// SetNillablePicture sets the "picture" field if the given value is not nil.
func (auo *AvatarUpdateOne) SetNillablePicture(a *avatar.Picture) *AvatarUpdateOne {
	if a != nil {
		auo.SetPicture(*a)
	}
	return auo
}

// SetNickname sets the "nickname" field.
func (auo *AvatarUpdateOne) SetNickname(s string) *AvatarUpdateOne {
	auo.mutation.SetNickname(s)
	return auo
}

// SetNillableNickname sets the "nickname" field if the given value is not nil.
func (auo *AvatarUpdateOne) SetNillableNickname(s *string) *AvatarUpdateOne {
	if s != nil {
		auo.SetNickname(*s)
	}
	return auo
}

// SetIntroduce sets the "introduce" field.
func (auo *AvatarUpdateOne) SetIntroduce(s string) *AvatarUpdateOne {
	auo.mutation.SetIntroduce(s)
	return auo
}

// SetNillableIntroduce sets the "introduce" field if the given value is not nil.
func (auo *AvatarUpdateOne) SetNillableIntroduce(s *string) *AvatarUpdateOne {
	if s != nil {
		auo.SetIntroduce(*s)
	}
	return auo
}

// ClearIntroduce clears the value of the "introduce" field.
func (auo *AvatarUpdateOne) ClearIntroduce() *AvatarUpdateOne {
	auo.mutation.ClearIntroduce()
	return auo
}

// SetState sets the "state" field.
func (auo *AvatarUpdateOne) SetState(a avatar.State) *AvatarUpdateOne {
	auo.mutation.SetState(a)
	return auo
}

// SetNillableState sets the "state" field if the given value is not nil.
func (auo *AvatarUpdateOne) SetNillableState(a *avatar.State) *AvatarUpdateOne {
	if a != nil {
		auo.SetState(*a)
	}
	return auo
}

// SetUserID sets the "user" edge to the User entity by ID.
func (auo *AvatarUpdateOne) SetUserID(id int) *AvatarUpdateOne {
	auo.mutation.SetUserID(id)
	return auo
}

// SetUser sets the "user" edge to the User entity.
func (auo *AvatarUpdateOne) SetUser(u *User) *AvatarUpdateOne {
	return auo.SetUserID(u.ID)
}

// AddAvatarRelationIDs adds the "avatar_relations" edge to the Relation entity by IDs.
func (auo *AvatarUpdateOne) AddAvatarRelationIDs(ids ...int) *AvatarUpdateOne {
	auo.mutation.AddAvatarRelationIDs(ids...)
	return auo
}

// AddAvatarRelations adds the "avatar_relations" edges to the Relation entity.
func (auo *AvatarUpdateOne) AddAvatarRelations(r ...*Relation) *AvatarUpdateOne {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return auo.AddAvatarRelationIDs(ids...)
}

// AddFriendRelationIDs adds the "friend_relations" edge to the Relation entity by IDs.
func (auo *AvatarUpdateOne) AddFriendRelationIDs(ids ...int) *AvatarUpdateOne {
	auo.mutation.AddFriendRelationIDs(ids...)
	return auo
}

// AddFriendRelations adds the "friend_relations" edges to the Relation entity.
func (auo *AvatarUpdateOne) AddFriendRelations(r ...*Relation) *AvatarUpdateOne {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return auo.AddFriendRelationIDs(ids...)
}

// AddMessageIDs adds the "messages" edge to the Message entity by IDs.
func (auo *AvatarUpdateOne) AddMessageIDs(ids ...int) *AvatarUpdateOne {
	auo.mutation.AddMessageIDs(ids...)
	return auo
}

// AddMessages adds the "messages" edges to the Message entity.
func (auo *AvatarUpdateOne) AddMessages(m ...*Message) *AvatarUpdateOne {
	ids := make([]int, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return auo.AddMessageIDs(ids...)
}

// Mutation returns the AvatarMutation object of the builder.
func (auo *AvatarUpdateOne) Mutation() *AvatarMutation {
	return auo.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (auo *AvatarUpdateOne) ClearUser() *AvatarUpdateOne {
	auo.mutation.ClearUser()
	return auo
}

// ClearAvatarRelations clears all "avatar_relations" edges to the Relation entity.
func (auo *AvatarUpdateOne) ClearAvatarRelations() *AvatarUpdateOne {
	auo.mutation.ClearAvatarRelations()
	return auo
}

// RemoveAvatarRelationIDs removes the "avatar_relations" edge to Relation entities by IDs.
func (auo *AvatarUpdateOne) RemoveAvatarRelationIDs(ids ...int) *AvatarUpdateOne {
	auo.mutation.RemoveAvatarRelationIDs(ids...)
	return auo
}

// RemoveAvatarRelations removes "avatar_relations" edges to Relation entities.
func (auo *AvatarUpdateOne) RemoveAvatarRelations(r ...*Relation) *AvatarUpdateOne {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return auo.RemoveAvatarRelationIDs(ids...)
}

// ClearFriendRelations clears all "friend_relations" edges to the Relation entity.
func (auo *AvatarUpdateOne) ClearFriendRelations() *AvatarUpdateOne {
	auo.mutation.ClearFriendRelations()
	return auo
}

// RemoveFriendRelationIDs removes the "friend_relations" edge to Relation entities by IDs.
func (auo *AvatarUpdateOne) RemoveFriendRelationIDs(ids ...int) *AvatarUpdateOne {
	auo.mutation.RemoveFriendRelationIDs(ids...)
	return auo
}

// RemoveFriendRelations removes "friend_relations" edges to Relation entities.
func (auo *AvatarUpdateOne) RemoveFriendRelations(r ...*Relation) *AvatarUpdateOne {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return auo.RemoveFriendRelationIDs(ids...)
}

// ClearMessages clears all "messages" edges to the Message entity.
func (auo *AvatarUpdateOne) ClearMessages() *AvatarUpdateOne {
	auo.mutation.ClearMessages()
	return auo
}

// RemoveMessageIDs removes the "messages" edge to Message entities by IDs.
func (auo *AvatarUpdateOne) RemoveMessageIDs(ids ...int) *AvatarUpdateOne {
	auo.mutation.RemoveMessageIDs(ids...)
	return auo
}

// RemoveMessages removes "messages" edges to Message entities.
func (auo *AvatarUpdateOne) RemoveMessages(m ...*Message) *AvatarUpdateOne {
	ids := make([]int, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return auo.RemoveMessageIDs(ids...)
}

// Where appends a list predicates to the AvatarUpdate builder.
func (auo *AvatarUpdateOne) Where(ps ...predicate.Avatar) *AvatarUpdateOne {
	auo.mutation.Where(ps...)
	return auo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (auo *AvatarUpdateOne) Select(field string, fields ...string) *AvatarUpdateOne {
	auo.fields = append([]string{field}, fields...)
	return auo
}

// Save executes the query and returns the updated Avatar entity.
func (auo *AvatarUpdateOne) Save(ctx context.Context) (*Avatar, error) {
	if err := auo.defaults(); err != nil {
		return nil, err
	}
	return withHooks(ctx, auo.sqlSave, auo.mutation, auo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (auo *AvatarUpdateOne) SaveX(ctx context.Context) *Avatar {
	node, err := auo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (auo *AvatarUpdateOne) Exec(ctx context.Context) error {
	_, err := auo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (auo *AvatarUpdateOne) ExecX(ctx context.Context) {
	if err := auo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (auo *AvatarUpdateOne) defaults() error {
	if _, ok := auo.mutation.UpdatedAt(); !ok {
		if avatar.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized avatar.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := avatar.UpdateDefaultUpdatedAt()
		auo.mutation.SetUpdatedAt(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (auo *AvatarUpdateOne) check() error {
	if v, ok := auo.mutation.Sex(); ok {
		if err := avatar.SexValidator(v); err != nil {
			return &ValidationError{Name: "sex", err: fmt.Errorf(`ent: validator failed for field "Avatar.sex": %w`, err)}
		}
	}
	if v, ok := auo.mutation.Birthday(); ok {
		if err := avatar.BirthdayValidator(v); err != nil {
			return &ValidationError{Name: "birthday", err: fmt.Errorf(`ent: validator failed for field "Avatar.birthday": %w`, err)}
		}
	}
	if v, ok := auo.mutation.Mbti(); ok {
		if err := avatar.MbtiValidator(v); err != nil {
			return &ValidationError{Name: "mbti", err: fmt.Errorf(`ent: validator failed for field "Avatar.mbti": %w`, err)}
		}
	}
	if v, ok := auo.mutation.Picture(); ok {
		if err := avatar.PictureValidator(v); err != nil {
			return &ValidationError{Name: "picture", err: fmt.Errorf(`ent: validator failed for field "Avatar.picture": %w`, err)}
		}
	}
	if v, ok := auo.mutation.Nickname(); ok {
		if err := avatar.NicknameValidator(v); err != nil {
			return &ValidationError{Name: "nickname", err: fmt.Errorf(`ent: validator failed for field "Avatar.nickname": %w`, err)}
		}
	}
	if v, ok := auo.mutation.State(); ok {
		if err := avatar.StateValidator(v); err != nil {
			return &ValidationError{Name: "state", err: fmt.Errorf(`ent: validator failed for field "Avatar.state": %w`, err)}
		}
	}
	if auo.mutation.UserCleared() && len(auo.mutation.UserIDs()) > 0 {
		return errors.New(`ent: clearing a required unique edge "Avatar.user"`)
	}
	return nil
}

func (auo *AvatarUpdateOne) sqlSave(ctx context.Context) (_node *Avatar, err error) {
	if err := auo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(avatar.Table, avatar.Columns, sqlgraph.NewFieldSpec(avatar.FieldID, field.TypeInt))
	id, ok := auo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Avatar.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := auo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, avatar.FieldID)
		for _, f := range fields {
			if !avatar.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != avatar.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := auo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := auo.mutation.DeletedAt(); ok {
		_spec.SetField(avatar.FieldDeletedAt, field.TypeTime, value)
	}
	if auo.mutation.DeletedAtCleared() {
		_spec.ClearField(avatar.FieldDeletedAt, field.TypeTime)
	}
	if value, ok := auo.mutation.UpdatedAt(); ok {
		_spec.SetField(avatar.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := auo.mutation.Sex(); ok {
		_spec.SetField(avatar.FieldSex, field.TypeEnum, value)
	}
	if value, ok := auo.mutation.Birthday(); ok {
		_spec.SetField(avatar.FieldBirthday, field.TypeString, value)
	}
	if value, ok := auo.mutation.Mbti(); ok {
		_spec.SetField(avatar.FieldMbti, field.TypeString, value)
	}
	if auo.mutation.MbtiCleared() {
		_spec.ClearField(avatar.FieldMbti, field.TypeString)
	}
	if value, ok := auo.mutation.Picture(); ok {
		_spec.SetField(avatar.FieldPicture, field.TypeEnum, value)
	}
	if value, ok := auo.mutation.Nickname(); ok {
		_spec.SetField(avatar.FieldNickname, field.TypeString, value)
	}
	if value, ok := auo.mutation.Introduce(); ok {
		_spec.SetField(avatar.FieldIntroduce, field.TypeString, value)
	}
	if auo.mutation.IntroduceCleared() {
		_spec.ClearField(avatar.FieldIntroduce, field.TypeString)
	}
	if value, ok := auo.mutation.State(); ok {
		_spec.SetField(avatar.FieldState, field.TypeEnum, value)
	}
	if auo.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   avatar.UserTable,
			Columns: []string{avatar.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := auo.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   avatar.UserTable,
			Columns: []string{avatar.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if auo.mutation.AvatarRelationsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   avatar.AvatarRelationsTable,
			Columns: []string{avatar.AvatarRelationsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(relation.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := auo.mutation.RemovedAvatarRelationsIDs(); len(nodes) > 0 && !auo.mutation.AvatarRelationsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   avatar.AvatarRelationsTable,
			Columns: []string{avatar.AvatarRelationsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(relation.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := auo.mutation.AvatarRelationsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   avatar.AvatarRelationsTable,
			Columns: []string{avatar.AvatarRelationsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(relation.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if auo.mutation.FriendRelationsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   avatar.FriendRelationsTable,
			Columns: []string{avatar.FriendRelationsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(relation.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := auo.mutation.RemovedFriendRelationsIDs(); len(nodes) > 0 && !auo.mutation.FriendRelationsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   avatar.FriendRelationsTable,
			Columns: []string{avatar.FriendRelationsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(relation.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := auo.mutation.FriendRelationsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   avatar.FriendRelationsTable,
			Columns: []string{avatar.FriendRelationsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(relation.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if auo.mutation.MessagesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   avatar.MessagesTable,
			Columns: []string{avatar.MessagesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(message.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := auo.mutation.RemovedMessagesIDs(); len(nodes) > 0 && !auo.mutation.MessagesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   avatar.MessagesTable,
			Columns: []string{avatar.MessagesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(message.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := auo.mutation.MessagesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   avatar.MessagesTable,
			Columns: []string{avatar.MessagesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(message.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Avatar{config: auo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, auo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{avatar.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	auo.mutation.done = true
	return _node, nil
}
