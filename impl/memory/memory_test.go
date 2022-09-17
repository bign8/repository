package memory_test

import (
	"context"
	"testing"

	"github.com/bign8/repository"
	"github.com/bign8/repository/impl/memory"
)

type Storable struct {
	ID    uint64
	Value string
}

func chk(tb testing.TB, err error, msg string) {
	if err != nil {
		tb.Errorf(msg+`: %v`, err)
	}
}

func TestCRUD(t *testing.T) {
	repo, err := memory.New(func(s *Storable, id *uint64) {
		if *id == 0 {
			*id = s.ID
		} else {
			s.ID = *id
		}
	})
	chk(t, err, `memory.New`)

	one := Storable{Value: `one`}
	two := Storable{Value: `two`}

	err = repo.Create(context.TODO(), &one, &two)
	chk(t, err, `memory.New`)

	valueAccessor := func(t *Storable) string {
		return t.Value
	}

	cond := repository.Equal(`value`, `two`, valueAccessor)
	v, err := repo.Get(context.TODO(), cond)
	chk(t, err, `repo.Get`)
	if v.Value != `two` {
		t.Errorf(`wanted %s, got %q`, cond, v.Value)
	}

	cond = repository.Equal(`value`, `one`, valueAccessor)
	v, err = repo.Get(context.TODO(), cond)
	chk(t, err, `repo.Get`)
	if v.Value != `one` {
		t.Errorf(`wanted %s, got %q`, cond, v.Value)
	}

	v.Value = `one more time!`
	err = repo.Update(context.TODO(), v)
	chk(t, err, `repo.Update`)

	err = repo.Delete(context.TODO(), &one, &two)
	chk(t, err, `repo.Delete`)
}
