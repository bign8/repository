package gorm

import (
	"context"

	"gorm.io/gorm"

	"github.com/bign8/repository"
)

// compile type type checking
var _ repository.Repository[any] = (*repo[any])(nil)

func New[T any](db *gorm.DB) (repository.Repository[T], error) {
	if err := db.AutoMigrate(new(T)); err != nil {
		return nil, err
	}
	return &repo[T]{db: db}, nil
}

type repo[T any] struct {
	db *gorm.DB
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
