package memory_test

import (
	"context"
	"testing"

	"github.com/bign8/repository/impl/memory"
)

type Storable struct {
	ID    uint
	Value string
}

func (store Storable) GetOrCreateID() (*uint, error) {
	store.ID = 3
	return &store.ID, nil
}

func chk(tb testing.TB, err error, msg string) {
	if err != nil {
		tb.Errorf(msg+`: %v`, err)
	}
}

func TestCRUD(t *testing.T) {
	repo, err := memory.New[Storable, uint]()
	chk(t, err, `memory.New`)

	one := Storable{Value: `one`}
	two := Storable{Value: `two`}

	err = repo.Create(context.TODO(), &one, &two)
	chk(t, err, `memory.New`)

	v, err := repo.Get(context.TODO())
	chk(t, err, `repo.Get`)
	if v.Value != `one` {
		t.Errorf(`where is "one", got %q`, v.Value)
	}

	v.Value = `one more time!`
	err = repo.Update(context.TODO(), v)
	chk(t, err, `repo.Update`)

	err = repo.Delete(context.TODO(), &one, &two)
	chk(t, err, `repo.Delete`)
}
