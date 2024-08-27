// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"log"
	"reflect"

	"toktok-backend/internal/adapter/persistence/mysql/ent/migrate"

	"toktok-backend/internal/adapter/persistence/mysql/ent/avatar"
	"toktok-backend/internal/adapter/persistence/mysql/ent/message"
	"toktok-backend/internal/adapter/persistence/mysql/ent/relation"
	"toktok-backend/internal/adapter/persistence/mysql/ent/room"
	"toktok-backend/internal/adapter/persistence/mysql/ent/user"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Avatar is the client for interacting with the Avatar builders.
	Avatar *AvatarClient
	// Message is the client for interacting with the Message builders.
	Message *MessageClient
	// Relation is the client for interacting with the Relation builders.
	Relation *RelationClient
	// Room is the client for interacting with the Room builders.
	Room *RoomClient
	// User is the client for interacting with the User builders.
	User *UserClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	client := &Client{config: newConfig(opts...)}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.Avatar = NewAvatarClient(c.config)
	c.Message = NewMessageClient(c.config)
	c.Relation = NewRelationClient(c.config)
	c.Room = NewRoomClient(c.config)
	c.User = NewUserClient(c.config)
}

type (
	// config is the configuration for the client and its builder.
	config struct {
		// driver used for executing database requests.
		driver dialect.Driver
		// debug enable a debug logging.
		debug bool
		// log used for logging on debug mode.
		log func(...any)
		// hooks to execute on mutations.
		hooks *hooks
		// interceptors to execute on queries.
		inters *inters
	}
	// Option function to configure the client.
	Option func(*config)
)

// newConfig creates a new config for the client.
func newConfig(opts ...Option) config {
	cfg := config{log: log.Println, hooks: &hooks{}, inters: &inters{}}
	cfg.options(opts...)
	return cfg
}

// options applies the options on the config object.
func (c *config) options(opts ...Option) {
	for _, opt := range opts {
		opt(c)
	}
	if c.debug {
		c.driver = dialect.Debug(c.driver, c.log)
	}
}

// Debug enables debug logging on the ent.Driver.
func Debug() Option {
	return func(c *config) {
		c.debug = true
	}
}

// Log sets the logging function for debug mode.
func Log(fn func(...any)) Option {
	return func(c *config) {
		c.log = fn
	}
}

// Driver configures the client driver.
func Driver(driver dialect.Driver) Option {
	return func(c *config) {
		c.driver = driver
	}
}

// Open opens a database/sql.DB specified by the driver name and
// the data source name, and returns a new client attached to it.
// Optional parameters can be added for configuring the client.
func Open(driverName, dataSourceName string, options ...Option) (*Client, error) {
	switch driverName {
	case dialect.MySQL, dialect.Postgres, dialect.SQLite:
		drv, err := sql.Open(driverName, dataSourceName)
		if err != nil {
			return nil, err
		}
		return NewClient(append(options, Driver(drv))...), nil
	default:
		return nil, fmt.Errorf("unsupported driver: %q", driverName)
	}
}

// ErrTxStarted is returned when trying to start a new transaction from a transactional client.
var ErrTxStarted = errors.New("ent: cannot start a transaction within a transaction")

// Tx returns a new transactional client. The provided context
// is used until the transaction is committed or rolled back.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, ErrTxStarted
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = tx
	return &Tx{
		ctx:      ctx,
		config:   cfg,
		Avatar:   NewAvatarClient(cfg),
		Message:  NewMessageClient(cfg),
		Relation: NewRelationClient(cfg),
		Room:     NewRoomClient(cfg),
		User:     NewUserClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with specified options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, errors.New("ent: cannot start a transaction within a transaction")
	}
	tx, err := c.driver.(interface {
		BeginTx(context.Context, *sql.TxOptions) (dialect.Tx, error)
	}).BeginTx(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = &txDriver{tx: tx, drv: c.driver}
	return &Tx{
		ctx:      ctx,
		config:   cfg,
		Avatar:   NewAvatarClient(cfg),
		Message:  NewMessageClient(cfg),
		Relation: NewRelationClient(cfg),
		Room:     NewRoomClient(cfg),
		User:     NewUserClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Avatar.
//		Query().
//		Count(ctx)
func (c *Client) Debug() *Client {
	if c.debug {
		return c
	}
	cfg := c.config
	cfg.driver = dialect.Debug(c.driver, c.log)
	client := &Client{config: cfg}
	client.init()
	return client
}

// Close closes the database connection and prevents new queries from starting.
func (c *Client) Close() error {
	return c.driver.Close()
}

// Use adds the mutation hooks to all the entity clients.
// In order to add hooks to a specific client, call: `client.Node.Use(...)`.
func (c *Client) Use(hooks ...Hook) {
	c.Avatar.Use(hooks...)
	c.Message.Use(hooks...)
	c.Relation.Use(hooks...)
	c.Room.Use(hooks...)
	c.User.Use(hooks...)
}

// Intercept adds the query interceptors to all the entity clients.
// In order to add interceptors to a specific client, call: `client.Node.Intercept(...)`.
func (c *Client) Intercept(interceptors ...Interceptor) {
	c.Avatar.Intercept(interceptors...)
	c.Message.Intercept(interceptors...)
	c.Relation.Intercept(interceptors...)
	c.Room.Intercept(interceptors...)
	c.User.Intercept(interceptors...)
}

// Mutate implements the ent.Mutator interface.
func (c *Client) Mutate(ctx context.Context, m Mutation) (Value, error) {
	switch m := m.(type) {
	case *AvatarMutation:
		return c.Avatar.mutate(ctx, m)
	case *MessageMutation:
		return c.Message.mutate(ctx, m)
	case *RelationMutation:
		return c.Relation.mutate(ctx, m)
	case *RoomMutation:
		return c.Room.mutate(ctx, m)
	case *UserMutation:
		return c.User.mutate(ctx, m)
	default:
		return nil, fmt.Errorf("ent: unknown mutation type %T", m)
	}
}

// AvatarClient is a client for the Avatar schema.
type AvatarClient struct {
	config
}

// NewAvatarClient returns a client for the Avatar from the given config.
func NewAvatarClient(c config) *AvatarClient {
	return &AvatarClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `avatar.Hooks(f(g(h())))`.
func (c *AvatarClient) Use(hooks ...Hook) {
	c.hooks.Avatar = append(c.hooks.Avatar, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `avatar.Intercept(f(g(h())))`.
func (c *AvatarClient) Intercept(interceptors ...Interceptor) {
	c.inters.Avatar = append(c.inters.Avatar, interceptors...)
}

// Create returns a builder for creating a Avatar entity.
func (c *AvatarClient) Create() *AvatarCreate {
	mutation := newAvatarMutation(c.config, OpCreate)
	return &AvatarCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Avatar entities.
func (c *AvatarClient) CreateBulk(builders ...*AvatarCreate) *AvatarCreateBulk {
	return &AvatarCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *AvatarClient) MapCreateBulk(slice any, setFunc func(*AvatarCreate, int)) *AvatarCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &AvatarCreateBulk{err: fmt.Errorf("calling to AvatarClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*AvatarCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &AvatarCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Avatar.
func (c *AvatarClient) Update() *AvatarUpdate {
	mutation := newAvatarMutation(c.config, OpUpdate)
	return &AvatarUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *AvatarClient) UpdateOne(a *Avatar) *AvatarUpdateOne {
	mutation := newAvatarMutation(c.config, OpUpdateOne, withAvatar(a))
	return &AvatarUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *AvatarClient) UpdateOneID(id int) *AvatarUpdateOne {
	mutation := newAvatarMutation(c.config, OpUpdateOne, withAvatarID(id))
	return &AvatarUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Avatar.
func (c *AvatarClient) Delete() *AvatarDelete {
	mutation := newAvatarMutation(c.config, OpDelete)
	return &AvatarDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *AvatarClient) DeleteOne(a *Avatar) *AvatarDeleteOne {
	return c.DeleteOneID(a.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *AvatarClient) DeleteOneID(id int) *AvatarDeleteOne {
	builder := c.Delete().Where(avatar.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &AvatarDeleteOne{builder}
}

// Query returns a query builder for Avatar.
func (c *AvatarClient) Query() *AvatarQuery {
	return &AvatarQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeAvatar},
		inters: c.Interceptors(),
	}
}

// Get returns a Avatar entity by its id.
func (c *AvatarClient) Get(ctx context.Context, id int) (*Avatar, error) {
	return c.Query().Where(avatar.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *AvatarClient) GetX(ctx context.Context, id int) *Avatar {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryUser queries the user edge of a Avatar.
func (c *AvatarClient) QueryUser(a *Avatar) *UserQuery {
	query := (&UserClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := a.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(avatar.Table, avatar.FieldID, id),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, avatar.UserTable, avatar.UserColumn),
		)
		fromV = sqlgraph.Neighbors(a.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryAvatarRelations queries the avatar_relations edge of a Avatar.
func (c *AvatarClient) QueryAvatarRelations(a *Avatar) *RelationQuery {
	query := (&RelationClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := a.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(avatar.Table, avatar.FieldID, id),
			sqlgraph.To(relation.Table, relation.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, avatar.AvatarRelationsTable, avatar.AvatarRelationsColumn),
		)
		fromV = sqlgraph.Neighbors(a.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryFriendRelations queries the friend_relations edge of a Avatar.
func (c *AvatarClient) QueryFriendRelations(a *Avatar) *RelationQuery {
	query := (&RelationClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := a.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(avatar.Table, avatar.FieldID, id),
			sqlgraph.To(relation.Table, relation.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, avatar.FriendRelationsTable, avatar.FriendRelationsColumn),
		)
		fromV = sqlgraph.Neighbors(a.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryMessages queries the messages edge of a Avatar.
func (c *AvatarClient) QueryMessages(a *Avatar) *MessageQuery {
	query := (&MessageClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := a.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(avatar.Table, avatar.FieldID, id),
			sqlgraph.To(message.Table, message.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, avatar.MessagesTable, avatar.MessagesColumn),
		)
		fromV = sqlgraph.Neighbors(a.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *AvatarClient) Hooks() []Hook {
	hooks := c.hooks.Avatar
	return append(hooks[:len(hooks):len(hooks)], avatar.Hooks[:]...)
}

// Interceptors returns the client interceptors.
func (c *AvatarClient) Interceptors() []Interceptor {
	inters := c.inters.Avatar
	return append(inters[:len(inters):len(inters)], avatar.Interceptors[:]...)
}

func (c *AvatarClient) mutate(ctx context.Context, m *AvatarMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&AvatarCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&AvatarUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&AvatarUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&AvatarDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown Avatar mutation op: %q", m.Op())
	}
}

// MessageClient is a client for the Message schema.
type MessageClient struct {
	config
}

// NewMessageClient returns a client for the Message from the given config.
func NewMessageClient(c config) *MessageClient {
	return &MessageClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `message.Hooks(f(g(h())))`.
func (c *MessageClient) Use(hooks ...Hook) {
	c.hooks.Message = append(c.hooks.Message, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `message.Intercept(f(g(h())))`.
func (c *MessageClient) Intercept(interceptors ...Interceptor) {
	c.inters.Message = append(c.inters.Message, interceptors...)
}

// Create returns a builder for creating a Message entity.
func (c *MessageClient) Create() *MessageCreate {
	mutation := newMessageMutation(c.config, OpCreate)
	return &MessageCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Message entities.
func (c *MessageClient) CreateBulk(builders ...*MessageCreate) *MessageCreateBulk {
	return &MessageCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *MessageClient) MapCreateBulk(slice any, setFunc func(*MessageCreate, int)) *MessageCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &MessageCreateBulk{err: fmt.Errorf("calling to MessageClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*MessageCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &MessageCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Message.
func (c *MessageClient) Update() *MessageUpdate {
	mutation := newMessageMutation(c.config, OpUpdate)
	return &MessageUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *MessageClient) UpdateOne(m *Message) *MessageUpdateOne {
	mutation := newMessageMutation(c.config, OpUpdateOne, withMessage(m))
	return &MessageUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *MessageClient) UpdateOneID(id int) *MessageUpdateOne {
	mutation := newMessageMutation(c.config, OpUpdateOne, withMessageID(id))
	return &MessageUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Message.
func (c *MessageClient) Delete() *MessageDelete {
	mutation := newMessageMutation(c.config, OpDelete)
	return &MessageDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *MessageClient) DeleteOne(m *Message) *MessageDeleteOne {
	return c.DeleteOneID(m.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *MessageClient) DeleteOneID(id int) *MessageDeleteOne {
	builder := c.Delete().Where(message.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &MessageDeleteOne{builder}
}

// Query returns a query builder for Message.
func (c *MessageClient) Query() *MessageQuery {
	return &MessageQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeMessage},
		inters: c.Interceptors(),
	}
}

// Get returns a Message entity by its id.
func (c *MessageClient) Get(ctx context.Context, id int) (*Message, error) {
	return c.Query().Where(message.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *MessageClient) GetX(ctx context.Context, id int) *Message {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryAvatar queries the avatar edge of a Message.
func (c *MessageClient) QueryAvatar(m *Message) *AvatarQuery {
	query := (&AvatarClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := m.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(message.Table, message.FieldID, id),
			sqlgraph.To(avatar.Table, avatar.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, message.AvatarTable, message.AvatarColumn),
		)
		fromV = sqlgraph.Neighbors(m.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryRoom queries the room edge of a Message.
func (c *MessageClient) QueryRoom(m *Message) *RoomQuery {
	query := (&RoomClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := m.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(message.Table, message.FieldID, id),
			sqlgraph.To(room.Table, room.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, message.RoomTable, message.RoomColumn),
		)
		fromV = sqlgraph.Neighbors(m.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *MessageClient) Hooks() []Hook {
	hooks := c.hooks.Message
	return append(hooks[:len(hooks):len(hooks)], message.Hooks[:]...)
}

// Interceptors returns the client interceptors.
func (c *MessageClient) Interceptors() []Interceptor {
	inters := c.inters.Message
	return append(inters[:len(inters):len(inters)], message.Interceptors[:]...)
}

func (c *MessageClient) mutate(ctx context.Context, m *MessageMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&MessageCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&MessageUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&MessageUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&MessageDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown Message mutation op: %q", m.Op())
	}
}

// RelationClient is a client for the Relation schema.
type RelationClient struct {
	config
}

// NewRelationClient returns a client for the Relation from the given config.
func NewRelationClient(c config) *RelationClient {
	return &RelationClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `relation.Hooks(f(g(h())))`.
func (c *RelationClient) Use(hooks ...Hook) {
	c.hooks.Relation = append(c.hooks.Relation, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `relation.Intercept(f(g(h())))`.
func (c *RelationClient) Intercept(interceptors ...Interceptor) {
	c.inters.Relation = append(c.inters.Relation, interceptors...)
}

// Create returns a builder for creating a Relation entity.
func (c *RelationClient) Create() *RelationCreate {
	mutation := newRelationMutation(c.config, OpCreate)
	return &RelationCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Relation entities.
func (c *RelationClient) CreateBulk(builders ...*RelationCreate) *RelationCreateBulk {
	return &RelationCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *RelationClient) MapCreateBulk(slice any, setFunc func(*RelationCreate, int)) *RelationCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &RelationCreateBulk{err: fmt.Errorf("calling to RelationClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*RelationCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &RelationCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Relation.
func (c *RelationClient) Update() *RelationUpdate {
	mutation := newRelationMutation(c.config, OpUpdate)
	return &RelationUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *RelationClient) UpdateOne(r *Relation) *RelationUpdateOne {
	mutation := newRelationMutation(c.config, OpUpdateOne, withRelation(r))
	return &RelationUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *RelationClient) UpdateOneID(id int) *RelationUpdateOne {
	mutation := newRelationMutation(c.config, OpUpdateOne, withRelationID(id))
	return &RelationUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Relation.
func (c *RelationClient) Delete() *RelationDelete {
	mutation := newRelationMutation(c.config, OpDelete)
	return &RelationDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *RelationClient) DeleteOne(r *Relation) *RelationDeleteOne {
	return c.DeleteOneID(r.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *RelationClient) DeleteOneID(id int) *RelationDeleteOne {
	builder := c.Delete().Where(relation.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &RelationDeleteOne{builder}
}

// Query returns a query builder for Relation.
func (c *RelationClient) Query() *RelationQuery {
	return &RelationQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeRelation},
		inters: c.Interceptors(),
	}
}

// Get returns a Relation entity by its id.
func (c *RelationClient) Get(ctx context.Context, id int) (*Relation, error) {
	return c.Query().Where(relation.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *RelationClient) GetX(ctx context.Context, id int) *Relation {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryAvatar queries the avatar edge of a Relation.
func (c *RelationClient) QueryAvatar(r *Relation) *AvatarQuery {
	query := (&AvatarClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := r.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(relation.Table, relation.FieldID, id),
			sqlgraph.To(avatar.Table, avatar.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, relation.AvatarTable, relation.AvatarColumn),
		)
		fromV = sqlgraph.Neighbors(r.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryFriend queries the friend edge of a Relation.
func (c *RelationClient) QueryFriend(r *Relation) *AvatarQuery {
	query := (&AvatarClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := r.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(relation.Table, relation.FieldID, id),
			sqlgraph.To(avatar.Table, avatar.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, relation.FriendTable, relation.FriendColumn),
		)
		fromV = sqlgraph.Neighbors(r.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryAvatarRooms queries the avatar_rooms edge of a Relation.
func (c *RelationClient) QueryAvatarRooms(r *Relation) *RoomQuery {
	query := (&RoomClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := r.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(relation.Table, relation.FieldID, id),
			sqlgraph.To(room.Table, room.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, relation.AvatarRoomsTable, relation.AvatarRoomsColumn),
		)
		fromV = sqlgraph.Neighbors(r.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryFriendRooms queries the friend_rooms edge of a Relation.
func (c *RelationClient) QueryFriendRooms(r *Relation) *RoomQuery {
	query := (&RoomClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := r.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(relation.Table, relation.FieldID, id),
			sqlgraph.To(room.Table, room.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, relation.FriendRoomsTable, relation.FriendRoomsColumn),
		)
		fromV = sqlgraph.Neighbors(r.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *RelationClient) Hooks() []Hook {
	hooks := c.hooks.Relation
	return append(hooks[:len(hooks):len(hooks)], relation.Hooks[:]...)
}

// Interceptors returns the client interceptors.
func (c *RelationClient) Interceptors() []Interceptor {
	inters := c.inters.Relation
	return append(inters[:len(inters):len(inters)], relation.Interceptors[:]...)
}

func (c *RelationClient) mutate(ctx context.Context, m *RelationMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&RelationCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&RelationUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&RelationUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&RelationDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown Relation mutation op: %q", m.Op())
	}
}

// RoomClient is a client for the Room schema.
type RoomClient struct {
	config
}

// NewRoomClient returns a client for the Room from the given config.
func NewRoomClient(c config) *RoomClient {
	return &RoomClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `room.Hooks(f(g(h())))`.
func (c *RoomClient) Use(hooks ...Hook) {
	c.hooks.Room = append(c.hooks.Room, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `room.Intercept(f(g(h())))`.
func (c *RoomClient) Intercept(interceptors ...Interceptor) {
	c.inters.Room = append(c.inters.Room, interceptors...)
}

// Create returns a builder for creating a Room entity.
func (c *RoomClient) Create() *RoomCreate {
	mutation := newRoomMutation(c.config, OpCreate)
	return &RoomCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Room entities.
func (c *RoomClient) CreateBulk(builders ...*RoomCreate) *RoomCreateBulk {
	return &RoomCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *RoomClient) MapCreateBulk(slice any, setFunc func(*RoomCreate, int)) *RoomCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &RoomCreateBulk{err: fmt.Errorf("calling to RoomClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*RoomCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &RoomCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Room.
func (c *RoomClient) Update() *RoomUpdate {
	mutation := newRoomMutation(c.config, OpUpdate)
	return &RoomUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *RoomClient) UpdateOne(r *Room) *RoomUpdateOne {
	mutation := newRoomMutation(c.config, OpUpdateOne, withRoom(r))
	return &RoomUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *RoomClient) UpdateOneID(id int) *RoomUpdateOne {
	mutation := newRoomMutation(c.config, OpUpdateOne, withRoomID(id))
	return &RoomUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Room.
func (c *RoomClient) Delete() *RoomDelete {
	mutation := newRoomMutation(c.config, OpDelete)
	return &RoomDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *RoomClient) DeleteOne(r *Room) *RoomDeleteOne {
	return c.DeleteOneID(r.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *RoomClient) DeleteOneID(id int) *RoomDeleteOne {
	builder := c.Delete().Where(room.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &RoomDeleteOne{builder}
}

// Query returns a query builder for Room.
func (c *RoomClient) Query() *RoomQuery {
	return &RoomQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeRoom},
		inters: c.Interceptors(),
	}
}

// Get returns a Room entity by its id.
func (c *RoomClient) Get(ctx context.Context, id int) (*Room, error) {
	return c.Query().Where(room.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *RoomClient) GetX(ctx context.Context, id int) *Room {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryAvatar queries the avatar edge of a Room.
func (c *RoomClient) QueryAvatar(r *Room) *RelationQuery {
	query := (&RelationClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := r.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(room.Table, room.FieldID, id),
			sqlgraph.To(relation.Table, relation.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, room.AvatarTable, room.AvatarColumn),
		)
		fromV = sqlgraph.Neighbors(r.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryFriend queries the friend edge of a Room.
func (c *RoomClient) QueryFriend(r *Room) *RelationQuery {
	query := (&RelationClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := r.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(room.Table, room.FieldID, id),
			sqlgraph.To(relation.Table, relation.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, room.FriendTable, room.FriendColumn),
		)
		fromV = sqlgraph.Neighbors(r.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryMessages queries the messages edge of a Room.
func (c *RoomClient) QueryMessages(r *Room) *MessageQuery {
	query := (&MessageClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := r.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(room.Table, room.FieldID, id),
			sqlgraph.To(message.Table, message.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, room.MessagesTable, room.MessagesColumn),
		)
		fromV = sqlgraph.Neighbors(r.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *RoomClient) Hooks() []Hook {
	hooks := c.hooks.Room
	return append(hooks[:len(hooks):len(hooks)], room.Hooks[:]...)
}

// Interceptors returns the client interceptors.
func (c *RoomClient) Interceptors() []Interceptor {
	inters := c.inters.Room
	return append(inters[:len(inters):len(inters)], room.Interceptors[:]...)
}

func (c *RoomClient) mutate(ctx context.Context, m *RoomMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&RoomCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&RoomUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&RoomUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&RoomDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown Room mutation op: %q", m.Op())
	}
}

// UserClient is a client for the User schema.
type UserClient struct {
	config
}

// NewUserClient returns a client for the User from the given config.
func NewUserClient(c config) *UserClient {
	return &UserClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `user.Hooks(f(g(h())))`.
func (c *UserClient) Use(hooks ...Hook) {
	c.hooks.User = append(c.hooks.User, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `user.Intercept(f(g(h())))`.
func (c *UserClient) Intercept(interceptors ...Interceptor) {
	c.inters.User = append(c.inters.User, interceptors...)
}

// Create returns a builder for creating a User entity.
func (c *UserClient) Create() *UserCreate {
	mutation := newUserMutation(c.config, OpCreate)
	return &UserCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of User entities.
func (c *UserClient) CreateBulk(builders ...*UserCreate) *UserCreateBulk {
	return &UserCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *UserClient) MapCreateBulk(slice any, setFunc func(*UserCreate, int)) *UserCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &UserCreateBulk{err: fmt.Errorf("calling to UserClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*UserCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &UserCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for User.
func (c *UserClient) Update() *UserUpdate {
	mutation := newUserMutation(c.config, OpUpdate)
	return &UserUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *UserClient) UpdateOne(u *User) *UserUpdateOne {
	mutation := newUserMutation(c.config, OpUpdateOne, withUser(u))
	return &UserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *UserClient) UpdateOneID(id int) *UserUpdateOne {
	mutation := newUserMutation(c.config, OpUpdateOne, withUserID(id))
	return &UserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for User.
func (c *UserClient) Delete() *UserDelete {
	mutation := newUserMutation(c.config, OpDelete)
	return &UserDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *UserClient) DeleteOne(u *User) *UserDeleteOne {
	return c.DeleteOneID(u.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *UserClient) DeleteOneID(id int) *UserDeleteOne {
	builder := c.Delete().Where(user.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &UserDeleteOne{builder}
}

// Query returns a query builder for User.
func (c *UserClient) Query() *UserQuery {
	return &UserQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeUser},
		inters: c.Interceptors(),
	}
}

// Get returns a User entity by its id.
func (c *UserClient) Get(ctx context.Context, id int) (*User, error) {
	return c.Query().Where(user.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *UserClient) GetX(ctx context.Context, id int) *User {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryAvatar queries the avatar edge of a User.
func (c *UserClient) QueryAvatar(u *User) *AvatarQuery {
	query := (&AvatarClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := u.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(user.Table, user.FieldID, id),
			sqlgraph.To(avatar.Table, avatar.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, user.AvatarTable, user.AvatarColumn),
		)
		fromV = sqlgraph.Neighbors(u.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *UserClient) Hooks() []Hook {
	hooks := c.hooks.User
	return append(hooks[:len(hooks):len(hooks)], user.Hooks[:]...)
}

// Interceptors returns the client interceptors.
func (c *UserClient) Interceptors() []Interceptor {
	inters := c.inters.User
	return append(inters[:len(inters):len(inters)], user.Interceptors[:]...)
}

func (c *UserClient) mutate(ctx context.Context, m *UserMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&UserCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&UserUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&UserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&UserDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown User mutation op: %q", m.Op())
	}
}

// hooks and interceptors per client, for fast access.
type (
	hooks struct {
		Avatar, Message, Relation, Room, User []ent.Hook
	}
	inters struct {
		Avatar, Message, Relation, Room, User []ent.Interceptor
	}
)
