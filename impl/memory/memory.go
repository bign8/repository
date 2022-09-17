package memory

import (
	"context"

	"github.com/bign8/repository"
)

func New[T any](accessor func(T, *uint64)) (repository.Repository[T], error) {
	// TODO: perform type checking on T?
	return &repo[T]{
		data:     make(map[uint64]T, 128),
		accessor: accessor,
	}, nil
}

type repo[T any] struct {
	ctr      uint64
	data     map[uint64]T
	accessor func(T, *uint64) // assign the ID of an object, zero is passed if we need ID
}

func (r *repo[T]) Create(ctx context.Context, obj ...T) error {
	for _, obj := range obj {
		r.ctr++
		r.accessor(obj, &r.ctr)
		r.data[r.ctr] = obj
	}
	return nil
}

func (r *repo[T]) Get(ctx context.Context, cond repository.Condition[T]) (T, error) {
	for _, value := range r.data {
		if cond.Match(value) {
			return value, nil
		}
	}
	return *new(T), repository.ErrNotFound
}

func (r *repo[T]) List(ctx context.Context, cond repository.Condition[T]) repository.Iterator[T] {
	return nil
}

func (r *repo[T]) Update(ctx context.Context, obj ...T) error {
	for _, obj := range obj {
		var id uint64
		r.accessor(obj, &id)
		r.data[id] = obj
	}
	return nil
}

func (r *repo[T]) Delete(ctx context.Context, obj ...T) error {
	for _, obj := range obj {
		var id uint64
		r.accessor(obj, &id)
		delete(r.data, id)
	}
	return nil
}
