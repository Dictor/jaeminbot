// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/dictor/jaeminbot/ent/command"
	"github.com/dictor/jaeminbot/ent/predicate"
	"github.com/dictor/jaeminbot/ent/resultlog"

	"entgo.io/ent"
)

const (
	// Operation types.
	OpCreate    = ent.OpCreate
	OpDelete    = ent.OpDelete
	OpDeleteOne = ent.OpDeleteOne
	OpUpdate    = ent.OpUpdate
	OpUpdateOne = ent.OpUpdateOne

	// Node types.
	TypeCommand   = "Command"
	TypeResultLog = "ResultLog"
)

// CommandMutation represents an operation that mutates the Command nodes in the graph.
type CommandMutation struct {
	config
	op            Op
	typ           string
	id            *string
	keyword       *string
	detail        *string
	created_at    *time.Time
	updated_at    *time.Time
	creator       *string
	server        *string
	code          *string
	clearedFields map[string]struct{}
	logs          map[int]struct{}
	removedlogs   map[int]struct{}
	clearedlogs   bool
	done          bool
	oldValue      func(context.Context) (*Command, error)
	predicates    []predicate.Command
}

var _ ent.Mutation = (*CommandMutation)(nil)

// commandOption allows management of the mutation configuration using functional options.
type commandOption func(*CommandMutation)

// newCommandMutation creates new mutation for the Command entity.
func newCommandMutation(c config, op Op, opts ...commandOption) *CommandMutation {
	m := &CommandMutation{
		config:        c,
		op:            op,
		typ:           TypeCommand,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withCommandID sets the ID field of the mutation.
func withCommandID(id string) commandOption {
	return func(m *CommandMutation) {
		var (
			err   error
			once  sync.Once
			value *Command
		)
		m.oldValue = func(ctx context.Context) (*Command, error) {
			once.Do(func() {
				if m.done {
					err = fmt.Errorf("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().Command.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withCommand sets the old Command of the mutation.
func withCommand(node *Command) commandOption {
	return func(m *CommandMutation) {
		m.oldValue = func(context.Context) (*Command, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m CommandMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m CommandMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, fmt.Errorf("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// SetID sets the value of the id field. Note that this
// operation is only accepted on creation of Command entities.
func (m *CommandMutation) SetID(id string) {
	m.id = &id
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *CommandMutation) ID() (id string, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// SetKeyword sets the "keyword" field.
func (m *CommandMutation) SetKeyword(s string) {
	m.keyword = &s
}

// Keyword returns the value of the "keyword" field in the mutation.
func (m *CommandMutation) Keyword() (r string, exists bool) {
	v := m.keyword
	if v == nil {
		return
	}
	return *v, true
}

// OldKeyword returns the old "keyword" field's value of the Command entity.
// If the Command object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *CommandMutation) OldKeyword(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldKeyword is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldKeyword requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldKeyword: %w", err)
	}
	return oldValue.Keyword, nil
}

// ResetKeyword resets all changes to the "keyword" field.
func (m *CommandMutation) ResetKeyword() {
	m.keyword = nil
}

// SetDetail sets the "detail" field.
func (m *CommandMutation) SetDetail(s string) {
	m.detail = &s
}

// Detail returns the value of the "detail" field in the mutation.
func (m *CommandMutation) Detail() (r string, exists bool) {
	v := m.detail
	if v == nil {
		return
	}
	return *v, true
}

// OldDetail returns the old "detail" field's value of the Command entity.
// If the Command object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *CommandMutation) OldDetail(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldDetail is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldDetail requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldDetail: %w", err)
	}
	return oldValue.Detail, nil
}

// ClearDetail clears the value of the "detail" field.
func (m *CommandMutation) ClearDetail() {
	m.detail = nil
	m.clearedFields[command.FieldDetail] = struct{}{}
}

// DetailCleared returns if the "detail" field was cleared in this mutation.
func (m *CommandMutation) DetailCleared() bool {
	_, ok := m.clearedFields[command.FieldDetail]
	return ok
}

// ResetDetail resets all changes to the "detail" field.
func (m *CommandMutation) ResetDetail() {
	m.detail = nil
	delete(m.clearedFields, command.FieldDetail)
}

// SetCreatedAt sets the "created_at" field.
func (m *CommandMutation) SetCreatedAt(t time.Time) {
	m.created_at = &t
}

// CreatedAt returns the value of the "created_at" field in the mutation.
func (m *CommandMutation) CreatedAt() (r time.Time, exists bool) {
	v := m.created_at
	if v == nil {
		return
	}
	return *v, true
}

// OldCreatedAt returns the old "created_at" field's value of the Command entity.
// If the Command object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *CommandMutation) OldCreatedAt(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldCreatedAt is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldCreatedAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldCreatedAt: %w", err)
	}
	return oldValue.CreatedAt, nil
}

// ResetCreatedAt resets all changes to the "created_at" field.
func (m *CommandMutation) ResetCreatedAt() {
	m.created_at = nil
}

// SetUpdatedAt sets the "updated_at" field.
func (m *CommandMutation) SetUpdatedAt(t time.Time) {
	m.updated_at = &t
}

// UpdatedAt returns the value of the "updated_at" field in the mutation.
func (m *CommandMutation) UpdatedAt() (r time.Time, exists bool) {
	v := m.updated_at
	if v == nil {
		return
	}
	return *v, true
}

// OldUpdatedAt returns the old "updated_at" field's value of the Command entity.
// If the Command object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *CommandMutation) OldUpdatedAt(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldUpdatedAt is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldUpdatedAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldUpdatedAt: %w", err)
	}
	return oldValue.UpdatedAt, nil
}

// ResetUpdatedAt resets all changes to the "updated_at" field.
func (m *CommandMutation) ResetUpdatedAt() {
	m.updated_at = nil
}

// SetCreator sets the "creator" field.
func (m *CommandMutation) SetCreator(s string) {
	m.creator = &s
}

// Creator returns the value of the "creator" field in the mutation.
func (m *CommandMutation) Creator() (r string, exists bool) {
	v := m.creator
	if v == nil {
		return
	}
	return *v, true
}

// OldCreator returns the old "creator" field's value of the Command entity.
// If the Command object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *CommandMutation) OldCreator(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldCreator is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldCreator requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldCreator: %w", err)
	}
	return oldValue.Creator, nil
}

// ResetCreator resets all changes to the "creator" field.
func (m *CommandMutation) ResetCreator() {
	m.creator = nil
}

// SetServer sets the "server" field.
func (m *CommandMutation) SetServer(s string) {
	m.server = &s
}

// Server returns the value of the "server" field in the mutation.
func (m *CommandMutation) Server() (r string, exists bool) {
	v := m.server
	if v == nil {
		return
	}
	return *v, true
}

// OldServer returns the old "server" field's value of the Command entity.
// If the Command object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *CommandMutation) OldServer(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldServer is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldServer requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldServer: %w", err)
	}
	return oldValue.Server, nil
}

// ResetServer resets all changes to the "server" field.
func (m *CommandMutation) ResetServer() {
	m.server = nil
}

// SetCode sets the "code" field.
func (m *CommandMutation) SetCode(s string) {
	m.code = &s
}

// Code returns the value of the "code" field in the mutation.
func (m *CommandMutation) Code() (r string, exists bool) {
	v := m.code
	if v == nil {
		return
	}
	return *v, true
}

// OldCode returns the old "code" field's value of the Command entity.
// If the Command object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *CommandMutation) OldCode(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldCode is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldCode requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldCode: %w", err)
	}
	return oldValue.Code, nil
}

// ResetCode resets all changes to the "code" field.
func (m *CommandMutation) ResetCode() {
	m.code = nil
}

// AddLogIDs adds the "logs" edge to the ResultLog entity by ids.
func (m *CommandMutation) AddLogIDs(ids ...int) {
	if m.logs == nil {
		m.logs = make(map[int]struct{})
	}
	for i := range ids {
		m.logs[ids[i]] = struct{}{}
	}
}

// ClearLogs clears the "logs" edge to the ResultLog entity.
func (m *CommandMutation) ClearLogs() {
	m.clearedlogs = true
}

// LogsCleared reports if the "logs" edge to the ResultLog entity was cleared.
func (m *CommandMutation) LogsCleared() bool {
	return m.clearedlogs
}

// RemoveLogIDs removes the "logs" edge to the ResultLog entity by IDs.
func (m *CommandMutation) RemoveLogIDs(ids ...int) {
	if m.removedlogs == nil {
		m.removedlogs = make(map[int]struct{})
	}
	for i := range ids {
		delete(m.logs, ids[i])
		m.removedlogs[ids[i]] = struct{}{}
	}
}

// RemovedLogs returns the removed IDs of the "logs" edge to the ResultLog entity.
func (m *CommandMutation) RemovedLogsIDs() (ids []int) {
	for id := range m.removedlogs {
		ids = append(ids, id)
	}
	return
}

// LogsIDs returns the "logs" edge IDs in the mutation.
func (m *CommandMutation) LogsIDs() (ids []int) {
	for id := range m.logs {
		ids = append(ids, id)
	}
	return
}

// ResetLogs resets all changes to the "logs" edge.
func (m *CommandMutation) ResetLogs() {
	m.logs = nil
	m.clearedlogs = false
	m.removedlogs = nil
}

// Where appends a list predicates to the CommandMutation builder.
func (m *CommandMutation) Where(ps ...predicate.Command) {
	m.predicates = append(m.predicates, ps...)
}

// Op returns the operation name.
func (m *CommandMutation) Op() Op {
	return m.op
}

// Type returns the node type of this mutation (Command).
func (m *CommandMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *CommandMutation) Fields() []string {
	fields := make([]string, 0, 7)
	if m.keyword != nil {
		fields = append(fields, command.FieldKeyword)
	}
	if m.detail != nil {
		fields = append(fields, command.FieldDetail)
	}
	if m.created_at != nil {
		fields = append(fields, command.FieldCreatedAt)
	}
	if m.updated_at != nil {
		fields = append(fields, command.FieldUpdatedAt)
	}
	if m.creator != nil {
		fields = append(fields, command.FieldCreator)
	}
	if m.server != nil {
		fields = append(fields, command.FieldServer)
	}
	if m.code != nil {
		fields = append(fields, command.FieldCode)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *CommandMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case command.FieldKeyword:
		return m.Keyword()
	case command.FieldDetail:
		return m.Detail()
	case command.FieldCreatedAt:
		return m.CreatedAt()
	case command.FieldUpdatedAt:
		return m.UpdatedAt()
	case command.FieldCreator:
		return m.Creator()
	case command.FieldServer:
		return m.Server()
	case command.FieldCode:
		return m.Code()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *CommandMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case command.FieldKeyword:
		return m.OldKeyword(ctx)
	case command.FieldDetail:
		return m.OldDetail(ctx)
	case command.FieldCreatedAt:
		return m.OldCreatedAt(ctx)
	case command.FieldUpdatedAt:
		return m.OldUpdatedAt(ctx)
	case command.FieldCreator:
		return m.OldCreator(ctx)
	case command.FieldServer:
		return m.OldServer(ctx)
	case command.FieldCode:
		return m.OldCode(ctx)
	}
	return nil, fmt.Errorf("unknown Command field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *CommandMutation) SetField(name string, value ent.Value) error {
	switch name {
	case command.FieldKeyword:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetKeyword(v)
		return nil
	case command.FieldDetail:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetDetail(v)
		return nil
	case command.FieldCreatedAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetCreatedAt(v)
		return nil
	case command.FieldUpdatedAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetUpdatedAt(v)
		return nil
	case command.FieldCreator:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetCreator(v)
		return nil
	case command.FieldServer:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetServer(v)
		return nil
	case command.FieldCode:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetCode(v)
		return nil
	}
	return fmt.Errorf("unknown Command field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *CommandMutation) AddedFields() []string {
	return nil
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *CommandMutation) AddedField(name string) (ent.Value, bool) {
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *CommandMutation) AddField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown Command numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *CommandMutation) ClearedFields() []string {
	var fields []string
	if m.FieldCleared(command.FieldDetail) {
		fields = append(fields, command.FieldDetail)
	}
	return fields
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *CommandMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *CommandMutation) ClearField(name string) error {
	switch name {
	case command.FieldDetail:
		m.ClearDetail()
		return nil
	}
	return fmt.Errorf("unknown Command nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *CommandMutation) ResetField(name string) error {
	switch name {
	case command.FieldKeyword:
		m.ResetKeyword()
		return nil
	case command.FieldDetail:
		m.ResetDetail()
		return nil
	case command.FieldCreatedAt:
		m.ResetCreatedAt()
		return nil
	case command.FieldUpdatedAt:
		m.ResetUpdatedAt()
		return nil
	case command.FieldCreator:
		m.ResetCreator()
		return nil
	case command.FieldServer:
		m.ResetServer()
		return nil
	case command.FieldCode:
		m.ResetCode()
		return nil
	}
	return fmt.Errorf("unknown Command field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *CommandMutation) AddedEdges() []string {
	edges := make([]string, 0, 1)
	if m.logs != nil {
		edges = append(edges, command.EdgeLogs)
	}
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *CommandMutation) AddedIDs(name string) []ent.Value {
	switch name {
	case command.EdgeLogs:
		ids := make([]ent.Value, 0, len(m.logs))
		for id := range m.logs {
			ids = append(ids, id)
		}
		return ids
	}
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *CommandMutation) RemovedEdges() []string {
	edges := make([]string, 0, 1)
	if m.removedlogs != nil {
		edges = append(edges, command.EdgeLogs)
	}
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *CommandMutation) RemovedIDs(name string) []ent.Value {
	switch name {
	case command.EdgeLogs:
		ids := make([]ent.Value, 0, len(m.removedlogs))
		for id := range m.removedlogs {
			ids = append(ids, id)
		}
		return ids
	}
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *CommandMutation) ClearedEdges() []string {
	edges := make([]string, 0, 1)
	if m.clearedlogs {
		edges = append(edges, command.EdgeLogs)
	}
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *CommandMutation) EdgeCleared(name string) bool {
	switch name {
	case command.EdgeLogs:
		return m.clearedlogs
	}
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *CommandMutation) ClearEdge(name string) error {
	switch name {
	}
	return fmt.Errorf("unknown Command unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *CommandMutation) ResetEdge(name string) error {
	switch name {
	case command.EdgeLogs:
		m.ResetLogs()
		return nil
	}
	return fmt.Errorf("unknown Command edge %s", name)
}

// ResultLogMutation represents an operation that mutates the ResultLog nodes in the graph.
type ResultLogMutation struct {
	config
	op             Op
	typ            string
	id             *int
	level          *string
	log            *string
	clearedFields  map[string]struct{}
	command        map[string]struct{}
	removedcommand map[string]struct{}
	clearedcommand bool
	done           bool
	oldValue       func(context.Context) (*ResultLog, error)
	predicates     []predicate.ResultLog
}

var _ ent.Mutation = (*ResultLogMutation)(nil)

// resultlogOption allows management of the mutation configuration using functional options.
type resultlogOption func(*ResultLogMutation)

// newResultLogMutation creates new mutation for the ResultLog entity.
func newResultLogMutation(c config, op Op, opts ...resultlogOption) *ResultLogMutation {
	m := &ResultLogMutation{
		config:        c,
		op:            op,
		typ:           TypeResultLog,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withResultLogID sets the ID field of the mutation.
func withResultLogID(id int) resultlogOption {
	return func(m *ResultLogMutation) {
		var (
			err   error
			once  sync.Once
			value *ResultLog
		)
		m.oldValue = func(ctx context.Context) (*ResultLog, error) {
			once.Do(func() {
				if m.done {
					err = fmt.Errorf("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().ResultLog.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withResultLog sets the old ResultLog of the mutation.
func withResultLog(node *ResultLog) resultlogOption {
	return func(m *ResultLogMutation) {
		m.oldValue = func(context.Context) (*ResultLog, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m ResultLogMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m ResultLogMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, fmt.Errorf("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *ResultLogMutation) ID() (id int, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// SetLevel sets the "level" field.
func (m *ResultLogMutation) SetLevel(s string) {
	m.level = &s
}

// Level returns the value of the "level" field in the mutation.
func (m *ResultLogMutation) Level() (r string, exists bool) {
	v := m.level
	if v == nil {
		return
	}
	return *v, true
}

// OldLevel returns the old "level" field's value of the ResultLog entity.
// If the ResultLog object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *ResultLogMutation) OldLevel(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldLevel is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldLevel requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldLevel: %w", err)
	}
	return oldValue.Level, nil
}

// ResetLevel resets all changes to the "level" field.
func (m *ResultLogMutation) ResetLevel() {
	m.level = nil
}

// SetLog sets the "log" field.
func (m *ResultLogMutation) SetLog(s string) {
	m.log = &s
}

// Log returns the value of the "log" field in the mutation.
func (m *ResultLogMutation) Log() (r string, exists bool) {
	v := m.log
	if v == nil {
		return
	}
	return *v, true
}

// OldLog returns the old "log" field's value of the ResultLog entity.
// If the ResultLog object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *ResultLogMutation) OldLog(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldLog is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldLog requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldLog: %w", err)
	}
	return oldValue.Log, nil
}

// ResetLog resets all changes to the "log" field.
func (m *ResultLogMutation) ResetLog() {
	m.log = nil
}

// AddCommandIDs adds the "command" edge to the Command entity by ids.
func (m *ResultLogMutation) AddCommandIDs(ids ...string) {
	if m.command == nil {
		m.command = make(map[string]struct{})
	}
	for i := range ids {
		m.command[ids[i]] = struct{}{}
	}
}

// ClearCommand clears the "command" edge to the Command entity.
func (m *ResultLogMutation) ClearCommand() {
	m.clearedcommand = true
}

// CommandCleared reports if the "command" edge to the Command entity was cleared.
func (m *ResultLogMutation) CommandCleared() bool {
	return m.clearedcommand
}

// RemoveCommandIDs removes the "command" edge to the Command entity by IDs.
func (m *ResultLogMutation) RemoveCommandIDs(ids ...string) {
	if m.removedcommand == nil {
		m.removedcommand = make(map[string]struct{})
	}
	for i := range ids {
		delete(m.command, ids[i])
		m.removedcommand[ids[i]] = struct{}{}
	}
}

// RemovedCommand returns the removed IDs of the "command" edge to the Command entity.
func (m *ResultLogMutation) RemovedCommandIDs() (ids []string) {
	for id := range m.removedcommand {
		ids = append(ids, id)
	}
	return
}

// CommandIDs returns the "command" edge IDs in the mutation.
func (m *ResultLogMutation) CommandIDs() (ids []string) {
	for id := range m.command {
		ids = append(ids, id)
	}
	return
}

// ResetCommand resets all changes to the "command" edge.
func (m *ResultLogMutation) ResetCommand() {
	m.command = nil
	m.clearedcommand = false
	m.removedcommand = nil
}

// Where appends a list predicates to the ResultLogMutation builder.
func (m *ResultLogMutation) Where(ps ...predicate.ResultLog) {
	m.predicates = append(m.predicates, ps...)
}

// Op returns the operation name.
func (m *ResultLogMutation) Op() Op {
	return m.op
}

// Type returns the node type of this mutation (ResultLog).
func (m *ResultLogMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *ResultLogMutation) Fields() []string {
	fields := make([]string, 0, 2)
	if m.level != nil {
		fields = append(fields, resultlog.FieldLevel)
	}
	if m.log != nil {
		fields = append(fields, resultlog.FieldLog)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *ResultLogMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case resultlog.FieldLevel:
		return m.Level()
	case resultlog.FieldLog:
		return m.Log()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *ResultLogMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case resultlog.FieldLevel:
		return m.OldLevel(ctx)
	case resultlog.FieldLog:
		return m.OldLog(ctx)
	}
	return nil, fmt.Errorf("unknown ResultLog field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *ResultLogMutation) SetField(name string, value ent.Value) error {
	switch name {
	case resultlog.FieldLevel:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetLevel(v)
		return nil
	case resultlog.FieldLog:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetLog(v)
		return nil
	}
	return fmt.Errorf("unknown ResultLog field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *ResultLogMutation) AddedFields() []string {
	return nil
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *ResultLogMutation) AddedField(name string) (ent.Value, bool) {
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *ResultLogMutation) AddField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown ResultLog numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *ResultLogMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *ResultLogMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *ResultLogMutation) ClearField(name string) error {
	return fmt.Errorf("unknown ResultLog nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *ResultLogMutation) ResetField(name string) error {
	switch name {
	case resultlog.FieldLevel:
		m.ResetLevel()
		return nil
	case resultlog.FieldLog:
		m.ResetLog()
		return nil
	}
	return fmt.Errorf("unknown ResultLog field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *ResultLogMutation) AddedEdges() []string {
	edges := make([]string, 0, 1)
	if m.command != nil {
		edges = append(edges, resultlog.EdgeCommand)
	}
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *ResultLogMutation) AddedIDs(name string) []ent.Value {
	switch name {
	case resultlog.EdgeCommand:
		ids := make([]ent.Value, 0, len(m.command))
		for id := range m.command {
			ids = append(ids, id)
		}
		return ids
	}
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *ResultLogMutation) RemovedEdges() []string {
	edges := make([]string, 0, 1)
	if m.removedcommand != nil {
		edges = append(edges, resultlog.EdgeCommand)
	}
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *ResultLogMutation) RemovedIDs(name string) []ent.Value {
	switch name {
	case resultlog.EdgeCommand:
		ids := make([]ent.Value, 0, len(m.removedcommand))
		for id := range m.removedcommand {
			ids = append(ids, id)
		}
		return ids
	}
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *ResultLogMutation) ClearedEdges() []string {
	edges := make([]string, 0, 1)
	if m.clearedcommand {
		edges = append(edges, resultlog.EdgeCommand)
	}
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *ResultLogMutation) EdgeCleared(name string) bool {
	switch name {
	case resultlog.EdgeCommand:
		return m.clearedcommand
	}
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *ResultLogMutation) ClearEdge(name string) error {
	switch name {
	}
	return fmt.Errorf("unknown ResultLog unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *ResultLogMutation) ResetEdge(name string) error {
	switch name {
	case resultlog.EdgeCommand:
		m.ResetCommand()
		return nil
	}
	return fmt.Errorf("unknown ResultLog edge %s", name)
}
