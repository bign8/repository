package repository

import "errors"

var ErrTodo = errors.New(`todo: not yet implemented`)

type Repository[T any] interface {
	// TODO: gorm based accessors, but strongly typed
}

type Query interface {
	// TODO: lookup appengine query interface
}
