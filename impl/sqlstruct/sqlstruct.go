package sqlstruct

import (
	"context"
	"database/sql"

	"github.com/kisielk/sqlstruct"

	"github.com/bign8/repository"
)

// compile type type checking
var _ repository.Repository[any] = (*repo[any])(nil)

func New[T any](db *sql.DB) (repository.Repository[T], error) {
	// TODO: perform type checking on T?
	return &repo[T]{db: db}, nil
}

type repo[T any] struct {
	// TODO: db connection
	db *sql.DB
}

func (r *repo[T]) Create(ctx context.Context, obj ...T) error {

	// dumy bit of code to get import to work
	println(sqlstruct.Columns(obj[0]))

	return repository.ErrNotImplemented
}

func (r *repo[T]) Get(ctx context.Context, conds ...repository.Condition) (T, error) {
	return *new(T), repository.ErrNotImplemented
}

func (r *repo[T]) List(ctx context.Context, conds ...repository.Condition) repository.Iterator[T] {
	return nil
}

func (r *repo[T]) Update(ctx context.Context, obj ...T) error {
	return repository.ErrNotImplemented
}

func (r *repo[T]) Delete(ctx context.Context, obj ...T) error {
	return repository.ErrNotImplemented
}
