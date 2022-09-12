package memory_test

import (
	"context"
	"testing"

	"github.com/bign8/repository"
	"github.com/bign8/repository/impl/memory"
)

var _ memory.Entity = (*Storable)(nil)

type Storable struct {
	ID    uint
	Value string
}

func (store Storable) Flatten() memory.Values {
	return memory.Values{
		`id`:    memory.Value[uint]{V: store.ID},
		`value`: memory.Value[string]{V: store.Value},
	}
}

func (store *Storable) Hydrate(values memory.Values) {
	for key, value := range values {
		switch key {
		case `id`:
			store.ID = value.(memory.Value[uint]).V
		case `value`:
			store.Value = value.(memory.Value[string]).V
		}
	}
}

func chk(tb testing.TB, err error, msg string) {
	if err != nil {
		tb.Errorf(msg+`: %v`, err)
	}
}

func TestCRUD(t *testing.T) {
	repo, err := memory.New[*Storable]()
	chk(t, err, `memory.New`)

	one := Storable{Value: `one`}
	two := Storable{Value: `two`}

	err = repo.Create(context.TODO(), &one, &two)
	chk(t, err, `memory.New`)

	cond := repository.Equal(`value`, `two`)
	v, err := repo.Get(context.TODO(), cond)
	chk(t, err, `repo.Get`)
	if v.Value != `two` {
		t.Errorf(`wanted %s, got %q`, cond, v.Value)
	}

	v.Value = `one more time!`
	err = repo.Update(context.TODO(), v)
	chk(t, err, `repo.Update`)

	err = repo.Delete(context.TODO(), &one, &two)
	chk(t, err, `repo.Delete`)
}
