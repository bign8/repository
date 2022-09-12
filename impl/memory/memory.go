package memory

import (
	"context"

	"github.com/bign8/repository"
)

type Value[V comparable] struct {
	V V
}

func (v Value[V]) onlyValueImplementsMe() {}

type Values map[string]interface{ onlyValueImplementsMe() }

type Entity interface {
	Flatten() Values
	Hydrate(Values)
}

func New[T Entity]() (repository.Repository[T], error) {
	// TODO: perform type checking on T?
	return &repo[T]{
		data: map[uint64]T{},
	}, nil
}

type repo[T Entity] struct {
	ctr     uint64
	data    map[uint64]T
	indexes map[string]btree // attribute => index
}

func (r *repo[T]) Create(ctx context.Context, obj ...T) error {
	for _, obj := range obj {
		r.ctr++
		r.data[r.ctr] = obj
		r.index(r.ctr, obj.Flatten())
	}
	return nil
}

func (r *repo[T]) index(id uint64, values Values) {
	for key, value := range values {
		tree := r.indexes[key]
		tree.insert(id, value)
	}
}

func (r *repo[T]) Get(ctx context.Context, conds ...repository.Condition) (T, error) {
	for _, value := range r.data {
		return value, nil
	}
	return *new(T), repository.ErrNotFound
}

func (r *repo[T]) List(ctx context.Context, conds ...repository.Condition) repository.Iterator[T] {
	return nil
}

func (r *repo[T]) Update(ctx context.Context, obj ...T) error {
	return nil
}

func (r *repo[T]) Delete(ctx context.Context, obj ...T) error {
	return nil
}
