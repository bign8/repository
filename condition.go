package repository

import (
	"fmt"
	"strings"
)

/*
	The implementations of Conditions will be different in each language.
	These should focus on the data representation of the conditions NOT on their implementation.
	YES, that means MATCH should NOT be a part of the Condition interface.
*/

var (
	_ Condition[any] = (*JoinCondition[any])(nil)
	_ Condition[any] = (*NotCondition[any])(nil)
	_ Condition[any] = (*BinaryCondition[any, int])(nil)
)

type Condition[T any] interface {
	String() string
	Match(T) bool
}

type Operator uint8

const (
	_ Operator = iota

	// join operators
	AND
	OR

	// binary operators
	EQUAL
	GREATER
	LESS
	GREATEROREQUAL
	LESSOREQUAL
)

func (join Operator) String() string {
	switch join {
	case AND:
		return `AND`
	case OR:
		return `OR`
	default:
		return `<unknown operator>`
	}
}

type JoinCondition[T any] struct {
	Conditions []Condition[T]
	Separator  Operator
}

func (join JoinCondition[T]) String() string {
	parts := make([]string, len(join.Conditions))
	for i, cond := range join.Conditions {
		parts[i] = cond.String()
	}
	whole := strings.Join(parts, ` `+join.Separator.String()+` `)
	return `(` + whole + `)`
}

func (join JoinCondition[T]) Match(t T) bool {
	switch join.Separator {
	case AND:
		for _, cond := range join.Conditions {
			if !cond.Match(t) {
				return false
			}
		}
		return true
	case OR:
		for _, cond := range join.Conditions {
			if cond.Match(t) {
				return true
			}
		}
		return false
	default:
		panic(`join condition: unknown separator`)
	}
}

func And[T any](conditions ...Condition[T]) Condition[T] {
	return &JoinCondition[T]{
		Conditions: conditions,
		Separator:  AND,
	}
}

type NotCondition[T any] struct {
	Condition Condition[T]
}

func (not NotCondition[T]) String() string {
	return `NOT(` + not.Condition.String() + `)`
}

func (not NotCondition[T]) Match(t T) bool {
	return !not.Condition.Match(t)
}

func Not[T any](condition Condition[T]) Condition[T] {
	return NotCondition[T]{
		Condition: condition,
	}
}

// binaryCondition defines binary operator as Specification
// It is used for =, >, <, >=, <= operators.
type BinaryCondition[T any, V comparable] struct {
	Field    string
	Operator Operator
	Value    V
	Accessor func(T) V
}

func (s BinaryCondition[T, V]) String() string {
	v, ok := any(s.Value).(Condition[T])
	if ok {
		return fmt.Sprintf(`%s %s %s`, s.Field, s.Operator, v.String())
	}
	return fmt.Sprintf("%s %s %v", s.Field, s.Operator, s.Value)
}

func (s BinaryCondition[T, V]) Match(t T) bool {
	v := s.Accessor(t)
	switch s.Operator {
	case EQUAL:
		return v == s.Value
	default:
		panic(`binary condition: unknown operator`)
	}
}

// Not delivers = operator as Specification
func Equal[T any, V comparable](field string, value V, accessor func(T) V) Condition[T] {
	// TODO: reflection? to fetch field from object (not needing accessor)
	return BinaryCondition[T, V]{
		Field:    field,
		Operator: EQUAL,
		Value:    value,
		Accessor: accessor,
	}
}

type FuncCondition[T any] func(T) bool

func (fn FuncCondition[T]) String() string {
	return `<user provided func>`
}

func (fn FuncCondition[T]) Match(t T) bool {
	return fn(t)
}
