package repository

import (
	"fmt"
	"strings"
)

var (
	_ Condition = (*joinCondition)(nil)
	_ Condition = (*notCondition)(nil)
)

type Condition interface {
	String() string
}

type joinCondition struct {
	Conditions []Condition
	Separator  string
}

func (join joinCondition) String() string {
	parts := make([]string, len(join.Conditions))
	for i, cond := range join.Conditions {
		parts[i] = cond.String()
	}
	whole := strings.Join(parts, ` `+join.Separator+` `)
	return `(` + whole + `)`
}

func And(conditions ...Condition) Condition {
	return &joinCondition{
		Conditions: conditions,
		Separator:  `AND`,
	}
}

type notCondition struct {
	Condition Condition
}

func (not notCondition) String() string {
	return `NOT(` + not.Condition.String() + `)`
}

func Not(condition Condition) Condition {
	return notCondition{
		Condition: condition,
	}
}

// binaryCondition defines binary operator as Specification
// It is used for =, >, <, >=, <= operators.
type binaryCondition[T any] struct {
	field    string
	operator string
	value    T
}

func (s binaryCondition[T]) String() string {
	v, ok := any(s.value).(Condition)
	if ok {
		return fmt.Sprintf(`%s %s %s`, s.field, s.operator, v.String())
	}
	return fmt.Sprintf("%s %s %v", s.field, s.operator, s.value)
}

// Not delivers = operator as Specification
func Equal[T any](field string, value T) Condition {
	return binaryCondition[T]{
		field:    field,
		operator: "=",
		value:    value,
	}
}
