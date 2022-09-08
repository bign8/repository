package memory

import (
	"context"

	"github.com/bign8/repository"
)

// compile type type checking
var _ repository.Repository[any] = (*repo[any])(nil)

func New[T any]() (repository.Repository[T], error) {
	// TODO: perform type checking on T?
	return &repo[T]{data: make([]T, 0, 128)}, nil
}

type repo[T any] struct {
	data []T
}

func (r *repo[T]) Create(ctx context.Context, obj ...*T) error {

	return repository.ErrNotImplemented
}

func (r *repo[T]) Get(ctx context.Context, conds ...repository.Condition) (*T, error) {
	return nil, repository.ErrNotImplemented
}

func (r *repo[T]) List(ctx context.Context, conds ...repository.Condition) repository.Iterator[T] {
	return nil
}

func (r *repo[T]) Update(ctx context.Context, obj ...*T) error {
	return repository.ErrNotImplemented
}

func (r *repo[T]) Delete(ctx context.Context, obj ...*T) error {
	return repository.ErrNotImplemented
}
