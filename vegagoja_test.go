package vegagoja

import (
	"context"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"
)

func TestVersion(t *testing.T) {
	ver, err := Version()
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}
	if exp := strings.TrimSpace(string(vegaVersion)); ver != exp {
		t.Errorf("expected %s, got: %s", exp, ver)
	}
}

func TestRender(t *testing.T) {
	ctx := context.Background()
	var files []string
	err := filepath.Walk("testdata", func(name string, info fs.FileInfo, err error) error {
		switch {
		case err != nil:
			return err
		case info.IsDir() || !strings.HasSuffix(name, ".vl.json"):
			return nil
		}
		files = append(files, name)
		return nil
	})
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}
	for _, nn := range files {
		name := nn
		t.Run(strings.TrimSuffix(filepath.Base(name), ".vl.json"), func(t *testing.T) {
			testRender(t, ctx, name)
		})
	}
}

func testRender(t *testing.T, ctx context.Context, name string) {
	t.Helper()
	spec, err := os.ReadFile(name)
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}
	ctx, cancel := context.WithTimeout(ctx, 1*time.Minute)
	defer cancel()
	vm := New(
		WithLogger(t.Log),
		WithVegaDemoData(),
	)
	s, err := vm.Render(ctx, string(spec), "")
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}
	t.Logf("len: %d", len(s))
	t.Logf("---\n%s\n---", s)
	if err := os.WriteFile(name+".svg", []byte(s), 0o644); err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}
}
