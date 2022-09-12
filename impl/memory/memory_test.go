package memory_test

import (
	"context"
	"testing"

	"github.com/bign8/repository"
	"github.com/bign8/repository/impl/memory"
)

var _ memory.Entity2 = (*Movie)(nil)

type Movie struct {
	Title       string
	Genre       string
	ReleaseYear uint16
}

func (movie Movie) Schema() []memory.AttributeDefinition {
	return []memory.AttributeDefinition{
		{
			Ident:       `:movie/title`,
			ValueType:   `:db.type/string`,
			Cardinality: memory.One,
			Doc:         `The title of the movie`,
		},
		{
			Ident:       `:movie/genre`,
			ValueType:   `:db.type/string`,
			Cardinality: memory.One,
			Doc:         `The genre of the movie`,
		},
		{
			Ident:       `:movie/release-year`,
			ValueType:   `:db.type/long`,
			Cardinality: memory.One,
			Doc:         `The year the movie was released in theaters`,
		},
	}
}

func (movie Movie) Values() memory.Values {
	return memory.Values{
		`:movie/title`:        memory.Value[string]{movie.Title},
		`:movie/genre`:        memory.Value[string]{movie.Genre},
		`:movie/release-year`: memory.Value[uint16]{movie.ReleaseYear},
	}
}

func (movie *Movie) Hydrate(values memory.Values) {
	for key, value := range values {
		switch key {
		case `:movie/title`:
			movie.Title = value.(memory.Value[string]).V
		case `:movie/genre`:
			movie.Genre = value.(memory.Value[string]).V
		case `:movie:release-year`:
			movie.ReleaseYear = value.(memory.Value[uint16]).V
		}
	}
}

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
	movies := []Movie{
		{
			Title:       "The Goonies",
			Genre:       "action/adventure",
			ReleaseYear: 1985,
		},
		{
			Title:       "Commando",
			Genre:       "action/adventure",
			ReleaseYear: 1985,
		},
		{
			Title:       "Repo Man",
			Genre:       "punk dystopia",
			ReleaseYear: 1984,
		},
	}
	_ = movies

	repo, err := memory.New[Storable, uint]()
	chk(t, err, `memory.New`)

	one := Storable{Value: `one`}
	two := Storable{Value: `two`}

	err = repo.Create(context.TODO(), &one, &two)
	chk(t, err, `memory.New`)

	v, err := repo.Get(context.TODO(), repository.Equal(`Value`, `one`))
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
