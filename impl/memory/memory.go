package memory

import (
	"context"
	"errors"

	"github.com/bign8/repository"
)

type Entity[ID comparable] interface {
	GetOrCreateID() (*ID, error)
}

func New[T Entity[ID], ID comparable]() (repository.Repository[T], error) {
	// TODO: perform type checking on T?
	return &repo[T, ID]{
		data: make(map[ID]*T, 128),
	}, nil
}

type repo[T Entity[ID], ID comparable] struct {
	data    map[ID]*T
	indexes map[string]btree // attribute => index
}

func (r *repo[T, ID]) Create(ctx context.Context, obj ...*T) error {
	for _, o := range obj {
		if o == nil {
			return errors.New(`nil object`)
		}
		id, err := (*o).GetOrCreateID()
		if err != nil {
			return err
		}
		if id == nil {
			return errors.New(`nil id returned`)
		}
		r.data[*id] = o
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
