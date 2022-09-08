package memory

import (
	"context"

	"github.com/bign8/repository"
)

// compile type type checking
var _ repository.Repository[any] = (*repo[any, int])(nil)

type Maker[T any, ID comparable] func(*T) ID

func New[T any, ID comparable](
	maker Maker[T, ID],
) (repository.Repository[T], error) {
	// TODO: perform type checking on T?
	return &repo[T, ID]{
		data: make(map[ID]*T, 128),
		make: maker,
	}, nil
}

type repo[T any, ID comparable] struct {
	data map[ID]*T
	make Maker[T, ID]
}

func (r *repo[T, ID]) Create(ctx context.Context, obj ...*T) error {
	for _, o := range obj {
		id := r.make(o)
		r.data[id] = o
	}
	return nil
}

func (r *repo[T, ID]) Get(ctx context.Context, conds ...repository.Condition) (*T, error) {
	for _, value := range r.data {
		return value, nil
	}
	return nil, repository.ErrNotFound
}

func (r *repo[T, ID]) List(ctx context.Context, conds ...repository.Condition) repository.Iterator[T] {
	return nil
}

func (r *repo[T, ID]) Update(ctx context.Context, obj ...*T) error {
	return nil
}

func (r *repo[T, ID]) Delete(ctx context.Context, obj ...*T) error {
	return nil
}
