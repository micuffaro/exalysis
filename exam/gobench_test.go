package exam

import (
	"fmt"
	"os"
	"path"
	"testing"

	"github.com/exercism/exalysis/extypes"
	"github.com/exercism/exalysis/gtpl"
	"github.com/exercism/exalysis/testhelper"
	"github.com/stretchr/testify/assert"
)

var benchTests = []struct {
	path       string
	expected   bool
	expectSugg bool
	suggestion gtpl.Template
	pkgName    string
}{
	{path: "./solutions/0", expected: false},
	{path: "./solutions/1", expected: true, pkgName: "twofer"},
	{path: "./solutions/2", expected: true, pkgName: "hamming"},
}

func TestGoBench(t *testing.T) {
	Benchmarks = true
	dir, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		if r := os.Chdir(dir); r != nil {
			err = r
		}
	}()

	for _, test := range benchTests {
		folder, _, err := testhelper.LoadFolder(path.Join(dir, test.path))
		if err != nil {
			t.Fatal(err)
		}
		if err := os.Chdir(folder.GetPath()); err != nil {
			t.Fatal(err)
		}

		r := extypes.NewResponse()
		ok := GoBench(folder, r, test.pkgName, false)

		failMsg := fmt.Sprintf("test failed: %+v", test)
		assert.Equal(t, test.expected, ok, failMsg)
		if test.suggestion != nil {
			assert.Equal(t, test.expectSugg, r.HasSuggestion(test.suggestion), failMsg)
		}
	}
}
