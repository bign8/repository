package repository

import (
	"context"
	"errors"
)

var (
	ErrNotImplemented = errors.New(`repository: not implemented`)
	ErrNotFound       = errors.New(`repository: not found`)
	ErrDone           = errors.New(`repository: iterator has no more results`)
)

// TODO: transaction things with context?

type Repository[T any] interface {

	// Create constructs + stores obj T.
	// If any identifiers are created as part of the insert, they are mutated on the incomming argument.
	Create(ctx context.Context, obj ...*T) error

	// Get returns a single item matching a set of conditions
	Get(ctx context.Context, conds ...Condition) (*T, error)

	// List
	List(ctx context.Context, conds ...Condition) Iterator[T]

	// Update
	Update(ctx context.Context, obj ...*T) error

	// Delete
	Delete(ctx context.Context, obj ...*T) error
}

type Query interface {
	// TODO: lookup appengine query interface

	// GORM based logic
	// func (db *DB) Assign(attrs ...interface{}) (tx *DB)
	// func (db *DB) Association(column string) *Association
	// func (db *DB) Attrs(attrs ...interface{}) (tx *DB)
	// func (db *DB) Begin(opts ...*sql.TxOptions) *DB
	// func (db *DB) Clauses(conds ...clause.Expression) (tx *DB)
	// func (db *DB) Commit() *DB
	// func (db *DB) Count(count *int64) (tx *DB)
	// func (db *DB) Create(value interface{}) (tx *DB)
	// func (db *DB) CreateInBatches(value interface{}, batchSize int) (tx *DB)
	// func (db *DB) Debug() (tx *DB)
	// func (db *DB) Delete(value interface{}, conds ...interface{}) (tx *DB)
	// func (db *DB) Distinct(args ...interface{}) (tx *DB)
	// func (db *DB) Exec(sql string, values ...interface{}) (tx *DB)
	// func (db *DB) Find(dest interface{}, conds ...interface{}) (tx *DB)
	// func (db *DB) FindInBatches(dest interface{}, batchSize int, fc func(tx *DB, batch int) error) *DB
	// func (db *DB) First(dest interface{}, conds ...interface{}) (tx *DB)
	// func (db *DB) FirstOrCreate(dest interface{}, conds ...interface{}) (tx *DB)
	// func (db *DB) FirstOrInit(dest interface{}, conds ...interface{}) (tx *DB)
	// func (db *DB) Get(key string) (interface{}, bool)
	// func (db *DB) Group(name string) (tx *DB)
	// func (db *DB) Having(query interface{}, args ...interface{}) (tx *DB)
	// func (db *DB) InstanceGet(key string) (interface{}, bool)
	// func (db *DB) InstanceSet(key string, value interface{}) *DB
	// func (db *DB) Joins(query string, args ...interface{}) (tx *DB)
	// func (db *DB) Last(dest interface{}, conds ...interface{}) (tx *DB)
	// func (db *DB) Limit(limit int) (tx *DB)
	// func (db *DB) Model(value interface{}) (tx *DB)
	// func (db *DB) Not(query interface{}, args ...interface{}) (tx *DB)
	// func (db *DB) Offset(offset int) (tx *DB)
	// func (db *DB) Omit(columns ...string) (tx *DB)
	// func (db *DB) Or(query interface{}, args ...interface{}) (tx *DB)
	// func (db *DB) Order(value interface{}) (tx *DB)
	// func (db *DB) Pluck(column string, dest interface{}) (tx *DB)
	// func (db *DB) Preload(query string, args ...interface{}) (tx *DB)
	// func (db *DB) Raw(sql string, values ...interface{}) (tx *DB)
	// func (db *DB) Rollback() *DB
	// func (db *DB) RollbackTo(name string) *DB
	// func (db *DB) Row() *sql.Row
	// func (db *DB) Rows() (*sql.Rows, error)
	// func (db *DB) Save(value interface{}) (tx *DB)
	// func (db *DB) SavePoint(name string) *DB
	// func (db *DB) Scan(dest interface{}) (tx *DB)
	// func (db *DB) ScanRows(rows *sql.Rows, dest interface{}) error
	// func (db *DB) Scopes(funcs ...func(*DB) *DB) (tx *DB)
	// func (db *DB) Select(query interface{}, args ...interface{}) (tx *DB)
	// func (db *DB) Session(config *Session) *DB
	// func (db *DB) Set(key string, value interface{}) *DB
	// func (db *DB) SetupJoinTable(model interface{}, field string, joinTable interface{}) error
	// func (db *DB) Table(name string, args ...interface{}) (tx *DB)
	// func (db *DB) Take(dest interface{}, conds ...interface{}) (tx *DB)
	// func (db *DB) Transaction(fc func(tx *DB) error, opts ...*sql.TxOptions) (err error)
	// func (db *DB) Unscoped() (tx *DB)
	// func (db *DB) Update(column string, value interface{}) (tx *DB)
	// func (db *DB) UpdateColumn(column string, value interface{}) (tx *DB)
	// func (db *DB) UpdateColumns(values interface{}) (tx *DB)
	// func (db *DB) Updates(values interface{}) (tx *DB)
	// func (db *DB) Where(query interface{}, args ...interface{}) (tx *DB)

	// Datastore logic
	// NewQuery(kind string) *Query
	// Ancestor(ancestor *Key) *Query
	// BatchSize(size int) *Query
	// Count(c context.Context) (int, error)
	// Distinct() *Query
	// DistinctOn(fieldNames ...string) *Query
	// End(c Cursor) *Query
	// EventualConsistency() *Query
	// Filter(filterStr string, value interface{}) *Query
	// GetAll(c context.Context, dst interface{}) ([]*Key, error)
	// KeysOnly() *Query
	// Limit(limit int) *Query
	// Offset(offset int) *Query
	// Order(fieldName string) *Query
	// Project(fieldNames ...string) *Query
	// Run(c context.Context) *Iterator
	// Start(c Cursor) *Query

}

type Iterator[T any] interface {
	Next() (*T, error)
}
