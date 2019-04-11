package gocyclo

import (
	"path/filepath"
	"testing"
)

func TestComplexity(t *testing.T) {
	files, err := filepath.Glob("*.go")
	if err != nil {
		t.Fatal(err)
	}
	max := 5
	result, ok := Assert(files, max)
	if !ok {
		for _, l := range result {
			t.Log(l)
		}
		t.Errorf("Exceeded maximum complexity %v", max)
	}
}
