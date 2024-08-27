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
	"toktok-backend/internal/adapter/persistence/mysql/ent/room"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// MessageUpdate is the builder for updating Message entities.
type MessageUpdate struct {
	config
	hooks    []Hook
	mutation *MessageMutation
}

// Where appends a list predicates to the MessageUpdate builder.
func (mu *MessageUpdate) Where(ps ...predicate.Message) *MessageUpdate {
	mu.mutation.Where(ps...)
	return mu
}

// SetDeletedAt sets the "deleted_at" field.
func (mu *MessageUpdate) SetDeletedAt(t time.Time) *MessageUpdate {
	mu.mutation.SetDeletedAt(t)
	return mu
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (mu *MessageUpdate) SetNillableDeletedAt(t *time.Time) *MessageUpdate {
	if t != nil {
		mu.SetDeletedAt(*t)
	}
	return mu
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (mu *MessageUpdate) ClearDeletedAt() *MessageUpdate {
	mu.mutation.ClearDeletedAt()
	return mu
}

// SetUpdatedAt sets the "updated_at" field.
func (mu *MessageUpdate) SetUpdatedAt(t time.Time) *MessageUpdate {
	mu.mutation.SetUpdatedAt(t)
	return mu
}

// SetState sets the "state" field.
func (mu *MessageUpdate) SetState(m message.State) *MessageUpdate {
	mu.mutation.SetState(m)
	return mu
}

// SetNillableState sets the "state" field if the given value is not nil.
func (mu *MessageUpdate) SetNillableState(m *message.State) *MessageUpdate {
	if m != nil {
		mu.SetState(*m)
	}
	return mu
}

// SetContent sets the "content" field.
func (mu *MessageUpdate) SetContent(s string) *MessageUpdate {
	mu.mutation.SetContent(s)
	return mu
}

// SetNillableContent sets the "content" field if the given value is not nil.
func (mu *MessageUpdate) SetNillableContent(s *string) *MessageUpdate {
	if s != nil {
		mu.SetContent(*s)
	}
	return mu
}

// SetEnteredAt sets the "entered_at" field.
func (mu *MessageUpdate) SetEnteredAt(t time.Time) *MessageUpdate {
	mu.mutation.SetEnteredAt(t)
	return mu
}

// SetNillableEnteredAt sets the "entered_at" field if the given value is not nil.
func (mu *MessageUpdate) SetNillableEnteredAt(t *time.Time) *MessageUpdate {
	if t != nil {
		mu.SetEnteredAt(*t)
	}
	return mu
}

// SetAvatarID sets the "avatar" edge to the Avatar entity by ID.
func (mu *MessageUpdate) SetAvatarID(id int) *MessageUpdate {
	mu.mutation.SetAvatarID(id)
	return mu
}

// SetNillableAvatarID sets the "avatar" edge to the Avatar entity by ID if the given value is not nil.
func (mu *MessageUpdate) SetNillableAvatarID(id *int) *MessageUpdate {
	if id != nil {
		mu = mu.SetAvatarID(*id)
	}
	return mu
}

// SetAvatar sets the "avatar" edge to the Avatar entity.
func (mu *MessageUpdate) SetAvatar(a *Avatar) *MessageUpdate {
	return mu.SetAvatarID(a.ID)
}

// SetRoomID sets the "room" edge to the Room entity by ID.
func (mu *MessageUpdate) SetRoomID(id int) *MessageUpdate {
	mu.mutation.SetRoomID(id)
	return mu
}

// SetNillableRoomID sets the "room" edge to the Room entity by ID if the given value is not nil.
func (mu *MessageUpdate) SetNillableRoomID(id *int) *MessageUpdate {
	if id != nil {
		mu = mu.SetRoomID(*id)
	}
	return mu
}

// SetRoom sets the "room" edge to the Room entity.
func (mu *MessageUpdate) SetRoom(r *Room) *MessageUpdate {
	return mu.SetRoomID(r.ID)
}

// Mutation returns the MessageMutation object of the builder.
func (mu *MessageUpdate) Mutation() *MessageMutation {
	return mu.mutation
}

// ClearAvatar clears the "avatar" edge to the Avatar entity.
func (mu *MessageUpdate) ClearAvatar() *MessageUpdate {
	mu.mutation.ClearAvatar()
	return mu
}

// ClearRoom clears the "room" edge to the Room entity.
func (mu *MessageUpdate) ClearRoom() *MessageUpdate {
	mu.mutation.ClearRoom()
	return mu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (mu *MessageUpdate) Save(ctx context.Context) (int, error) {
	if err := mu.defaults(); err != nil {
		return 0, err
	}
	return withHooks(ctx, mu.sqlSave, mu.mutation, mu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (mu *MessageUpdate) SaveX(ctx context.Context) int {
	affected, err := mu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (mu *MessageUpdate) Exec(ctx context.Context) error {
	_, err := mu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mu *MessageUpdate) ExecX(ctx context.Context) {
	if err := mu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (mu *MessageUpdate) defaults() error {
	if _, ok := mu.mutation.UpdatedAt(); !ok {
		if message.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized message.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := message.UpdateDefaultUpdatedAt()
		mu.mutation.SetUpdatedAt(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (mu *MessageUpdate) check() error {
	if v, ok := mu.mutation.State(); ok {
		if err := message.StateValidator(v); err != nil {
			return &ValidationError{Name: "state", err: fmt.Errorf(`ent: validator failed for field "Message.state": %w`, err)}
		}
	}
	if v, ok := mu.mutation.Content(); ok {
		if err := message.ContentValidator(v); err != nil {
			return &ValidationError{Name: "content", err: fmt.Errorf(`ent: validator failed for field "Message.content": %w`, err)}
		}
	}
	return nil
}

func (mu *MessageUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := mu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(message.Table, message.Columns, sqlgraph.NewFieldSpec(message.FieldID, field.TypeInt))
	if ps := mu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := mu.mutation.DeletedAt(); ok {
		_spec.SetField(message.FieldDeletedAt, field.TypeTime, value)
	}
	if mu.mutation.DeletedAtCleared() {
		_spec.ClearField(message.FieldDeletedAt, field.TypeTime)
	}
	if value, ok := mu.mutation.UpdatedAt(); ok {
		_spec.SetField(message.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := mu.mutation.State(); ok {
		_spec.SetField(message.FieldState, field.TypeEnum, value)
	}
	if value, ok := mu.mutation.Content(); ok {
		_spec.SetField(message.FieldContent, field.TypeString, value)
	}
	if value, ok := mu.mutation.EnteredAt(); ok {
		_spec.SetField(message.FieldEnteredAt, field.TypeTime, value)
	}
	if mu.mutation.AvatarCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   message.AvatarTable,
			Columns: []string{message.AvatarColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(avatar.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mu.mutation.AvatarIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   message.AvatarTable,
			Columns: []string{message.AvatarColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(avatar.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if mu.mutation.RoomCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   message.RoomTable,
			Columns: []string{message.RoomColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(room.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mu.mutation.RoomIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   message.RoomTable,
			Columns: []string{message.RoomColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(room.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, mu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{message.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	mu.mutation.done = true
	return n, nil
}

// MessageUpdateOne is the builder for updating a single Message entity.
type MessageUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *MessageMutation
}

// SetDeletedAt sets the "deleted_at" field.
func (muo *MessageUpdateOne) SetDeletedAt(t time.Time) *MessageUpdateOne {
	muo.mutation.SetDeletedAt(t)
	return muo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (muo *MessageUpdateOne) SetNillableDeletedAt(t *time.Time) *MessageUpdateOne {
	if t != nil {
		muo.SetDeletedAt(*t)
	}
	return muo
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (muo *MessageUpdateOne) ClearDeletedAt() *MessageUpdateOne {
	muo.mutation.ClearDeletedAt()
	return muo
}

// SetUpdatedAt sets the "updated_at" field.
func (muo *MessageUpdateOne) SetUpdatedAt(t time.Time) *MessageUpdateOne {
	muo.mutation.SetUpdatedAt(t)
	return muo
}

// SetState sets the "state" field.
func (muo *MessageUpdateOne) SetState(m message.State) *MessageUpdateOne {
	muo.mutation.SetState(m)
	return muo
}

// SetNillableState sets the "state" field if the given value is not nil.
func (muo *MessageUpdateOne) SetNillableState(m *message.State) *MessageUpdateOne {
	if m != nil {
		muo.SetState(*m)
	}
	return muo
}

// SetContent sets the "content" field.
func (muo *MessageUpdateOne) SetContent(s string) *MessageUpdateOne {
	muo.mutation.SetContent(s)
	return muo
}

// SetNillableContent sets the "content" field if the given value is not nil.
func (muo *MessageUpdateOne) SetNillableContent(s *string) *MessageUpdateOne {
	if s != nil {
		muo.SetContent(*s)
	}
	return muo
}

// SetEnteredAt sets the "entered_at" field.
func (muo *MessageUpdateOne) SetEnteredAt(t time.Time) *MessageUpdateOne {
	muo.mutation.SetEnteredAt(t)
	return muo
}

// SetNillableEnteredAt sets the "entered_at" field if the given value is not nil.
func (muo *MessageUpdateOne) SetNillableEnteredAt(t *time.Time) *MessageUpdateOne {
	if t != nil {
		muo.SetEnteredAt(*t)
	}
	return muo
}

// SetAvatarID sets the "avatar" edge to the Avatar entity by ID.
func (muo *MessageUpdateOne) SetAvatarID(id int) *MessageUpdateOne {
	muo.mutation.SetAvatarID(id)
	return muo
}

// SetNillableAvatarID sets the "avatar" edge to the Avatar entity by ID if the given value is not nil.
func (muo *MessageUpdateOne) SetNillableAvatarID(id *int) *MessageUpdateOne {
	if id != nil {
		muo = muo.SetAvatarID(*id)
	}
	return muo
}

// SetAvatar sets the "avatar" edge to the Avatar entity.
func (muo *MessageUpdateOne) SetAvatar(a *Avatar) *MessageUpdateOne {
	return muo.SetAvatarID(a.ID)
}

// SetRoomID sets the "room" edge to the Room entity by ID.
func (muo *MessageUpdateOne) SetRoomID(id int) *MessageUpdateOne {
	muo.mutation.SetRoomID(id)
	return muo
}

// SetNillableRoomID sets the "room" edge to the Room entity by ID if the given value is not nil.
func (muo *MessageUpdateOne) SetNillableRoomID(id *int) *MessageUpdateOne {
	if id != nil {
		muo = muo.SetRoomID(*id)
	}
	return muo
}

// SetRoom sets the "room" edge to the Room entity.
func (muo *MessageUpdateOne) SetRoom(r *Room) *MessageUpdateOne {
	return muo.SetRoomID(r.ID)
}

// Mutation returns the MessageMutation object of the builder.
func (muo *MessageUpdateOne) Mutation() *MessageMutation {
	return muo.mutation
}

// ClearAvatar clears the "avatar" edge to the Avatar entity.
func (muo *MessageUpdateOne) ClearAvatar() *MessageUpdateOne {
	muo.mutation.ClearAvatar()
	return muo
}

// ClearRoom clears the "room" edge to the Room entity.
func (muo *MessageUpdateOne) ClearRoom() *MessageUpdateOne {
	muo.mutation.ClearRoom()
	return muo
}

// Where appends a list predicates to the MessageUpdate builder.
func (muo *MessageUpdateOne) Where(ps ...predicate.Message) *MessageUpdateOne {
	muo.mutation.Where(ps...)
	return muo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (muo *MessageUpdateOne) Select(field string, fields ...string) *MessageUpdateOne {
	muo.fields = append([]string{field}, fields...)
	return muo
}

// Save executes the query and returns the updated Message entity.
func (muo *MessageUpdateOne) Save(ctx context.Context) (*Message, error) {
	if err := muo.defaults(); err != nil {
		return nil, err
	}
	return withHooks(ctx, muo.sqlSave, muo.mutation, muo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (muo *MessageUpdateOne) SaveX(ctx context.Context) *Message {
	node, err := muo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (muo *MessageUpdateOne) Exec(ctx context.Context) error {
	_, err := muo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (muo *MessageUpdateOne) ExecX(ctx context.Context) {
	if err := muo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (muo *MessageUpdateOne) defaults() error {
	if _, ok := muo.mutation.UpdatedAt(); !ok {
		if message.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized message.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := message.UpdateDefaultUpdatedAt()
		muo.mutation.SetUpdatedAt(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (muo *MessageUpdateOne) check() error {
	if v, ok := muo.mutation.State(); ok {
		if err := message.StateValidator(v); err != nil {
			return &ValidationError{Name: "state", err: fmt.Errorf(`ent: validator failed for field "Message.state": %w`, err)}
		}
	}
	if v, ok := muo.mutation.Content(); ok {
		if err := message.ContentValidator(v); err != nil {
			return &ValidationError{Name: "content", err: fmt.Errorf(`ent: validator failed for field "Message.content": %w`, err)}
		}
	}
	return nil
}

func (muo *MessageUpdateOne) sqlSave(ctx context.Context) (_node *Message, err error) {
	if err := muo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(message.Table, message.Columns, sqlgraph.NewFieldSpec(message.FieldID, field.TypeInt))
	id, ok := muo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Message.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := muo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, message.FieldID)
		for _, f := range fields {
			if !message.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != message.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := muo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := muo.mutation.DeletedAt(); ok {
		_spec.SetField(message.FieldDeletedAt, field.TypeTime, value)
	}
	if muo.mutation.DeletedAtCleared() {
		_spec.ClearField(message.FieldDeletedAt, field.TypeTime)
	}
	if value, ok := muo.mutation.UpdatedAt(); ok {
		_spec.SetField(message.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := muo.mutation.State(); ok {
		_spec.SetField(message.FieldState, field.TypeEnum, value)
	}
	if value, ok := muo.mutation.Content(); ok {
		_spec.SetField(message.FieldContent, field.TypeString, value)
	}
	if value, ok := muo.mutation.EnteredAt(); ok {
		_spec.SetField(message.FieldEnteredAt, field.TypeTime, value)
	}
	if muo.mutation.AvatarCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   message.AvatarTable,
			Columns: []string{message.AvatarColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(avatar.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := muo.mutation.AvatarIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   message.AvatarTable,
			Columns: []string{message.AvatarColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(avatar.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if muo.mutation.RoomCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   message.RoomTable,
			Columns: []string{message.RoomColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(room.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := muo.mutation.RoomIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   message.RoomTable,
			Columns: []string{message.RoomColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(room.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Message{config: muo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, muo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{message.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	muo.mutation.done = true
	return _node, nil
}
