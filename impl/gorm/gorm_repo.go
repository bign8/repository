package gorm

import "github.com/bign8/repository"

// compile type type checking
var _ repository.Repository[any] = (*repo[any])(nil)

type repo[T any] struct {
	// TODO: db connection
}

func (r *repo[T]) Get(query repository.Query) (*T, error) {
	return nil, repository.ErrTodo
}
