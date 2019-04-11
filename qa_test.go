package gocyclo

import (
	"path/filepath"
	"testing"
)

func TestComplexity(t *testing.T) {
	goFiles, _ := filepath.Glob("*.go")
	extra, _ := filepath.Glob("cmd/gocyclo/*.go")
	files := append(goFiles, extra...)
	max := 5
	result, ok := Assert(files, max)
	if !ok {
		t.Errorf("Maximum complexity %v, exceeded", max)
		for _, l := range result {
			t.Log(l)
		}
	}
}
