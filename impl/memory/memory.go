package memory

import (
	"context"
	"errors"

	"github.com/bign8/repository"
)

type Cardinality byte

const (
	_ Cardinality = iota
	One
	Multiple
	Unique
)

// type Attribute[T comparable] struct {
// 	Name  string
// 	Value T
// }
// func (attr Attribute[T]) onlyAttributeImplementsThis() {}
// type Attributes []interface { onlyAttributeImplementsThis() }

type Value[V comparable] struct {
	V V
}

func (value Value[T]) onlyValueImplementsThis() {}

type Values map[string]interface {
	onlyValueImplementsThis()
}

type AttributeDefinition struct {
	Ident       string      // unique name for your attribute
	ValueType   string      // type of data in the attribute
	Cardinality Cardinality // single or collection of values
	Doc         string      // human readable comment
}

type Entity[ID comparable] interface {
	GetOrCreateID() (*ID, error)
}

type Entity2 interface {
	Schema() []AttributeDefinition
	Values() Values
	Hydrate(Values)
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

	// entity: store a btree of identifier => attribute-value ID list
	enitty map[uint64]struct {
	}
	// attribute: each attribute from Schema will get a btree
	// values: unique set of values per type
	// time: TODO: figure out some sort of time reference
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
