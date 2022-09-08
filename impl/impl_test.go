package impl

import (
	"errors"
	"os"
	"path/filepath"
	"testing"
)

func TestSubModules(t *testing.T) {
	modules, err := os.ReadDir(`.`)
	if err != nil {
		t.Fatalf(`unable to listdir: %v`, err)
	}
	for _, entry := range modules {
		if !entry.IsDir() {
			continue
		}
		name := entry.Name()
		_, err = os.Stat(filepath.Join(`.`, name, `go.mod`))
		if errors.Is(err, os.ErrNotExist) {
			t.Errorf(`impl/%s/go.mod does not exist`, name)
			t.Log("run the following commands to resolve this")
			t.Logf("\tpushd impl/%s", name)
			t.Logf("\tgo mod init github.com/bign8/repository/impl/%s", name)
			t.Logf("\tpopd")
			t.Logf("\tgo work use ./impl/%s", name)
		} else {
			t.Logf(`impl/%s/go.mod exists!`, name)
		}
	}
}
