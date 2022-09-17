package memory

import (
	"context"

	"github.com/bign8/repository"
)

func New[T any, ID comparable](accessor func(T) ID) (repository.Repository[T], error) {
	// TODO: perform type checking on T?
	return &repo[T, ID]{
		data:     make(map[ID]T, 128),
		accessor: accessor,
	}, nil
}

type repo[T any, ID comparable] struct {
	data     map[ID]T
	accessor func(T) ID
}

func (r *repo[T, ID]) Create(ctx context.Context, obj ...T) error {
	for _, obj := range obj {
		r.data[r.accessor(obj)] = obj
	}
	return nil
}

func (r *repo[T, ID]) Get(ctx context.Context, cond repository.Condition[T]) (T, error) {
	for _, value := range r.data {
		if cond.Match(value) {
			return value, nil
		}
	}
	return *new(T), repository.ErrNotFound
}

func (r *repo[T, ID]) List(ctx context.Context, cond repository.Condition[T]) repository.Iterator[T] {
	return nil
}

func (r *repo[T, ID]) Update(ctx context.Context, obj ...T) error {
	for _, obj := range obj {
		r.data[r.accessor(obj)] = obj
	}
	return nil
}

func (r *repo[T, ID]) Delete(ctx context.Context, obj ...T) error {
	for _, obj := range obj {
		delete(r.data, r.accessor(obj))
	}
	return nil
}
