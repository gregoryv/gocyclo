package gocyclo

import (
	"container/list"
	"testing"

	"github.com/gregoryv/find"
)

func TestComplexity(t *testing.T) {
	goFiles, err := find.ByName("*.go", ".")
	if err != nil {
		t.Fatal(err)
	}

	files := fileNames(goFiles)
	max := 2
	result, ok := Assert(files, max)
	if !ok {
		t.Errorf("Maximum complexity %v, exceeded", max)
		for _, l := range result {
			t.Log(l)
		}
	}
}

func fileNames(files *list.List) (names []string) {
	var i int
	names = make([]string, files.Len())
	for e := files.Front(); e != nil; e = e.Next() {
		if s, ok := e.Value.(string); ok {
			names[i] = s
		}
		i++
	}
	return
}
