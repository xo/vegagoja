package vegagoja

import (
	"context"
	"io/fs"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"strings"
	"testing"
	"time"
)

func TestVersion(t *testing.T) {
	vegaVer, err := Version()
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}
	if exp := strings.TrimSpace(string(vegaVersionTxt)); vegaVer != exp {
		t.Errorf("expected %s, got: %s", exp, vegaVer)
	}
	t.Logf("vega version: %s", vegaVer)
}

func TestRender(t *testing.T) {
	ctx := context.Background()
	timeout := 1 * time.Minute
	if s := os.Getenv("TIMEOUT"); s != "" {
		var err error
		if timeout, err = time.ParseDuration(s); err != nil {
			t.Fatalf("could not parse timeout %q: %v", s, err)
		}
	}
	var files []string
	err := filepath.Walk("testdata", func(name string, info fs.FileInfo, err error) error {
		switch {
		case err != nil:
			return err
		case info.IsDir() || !suffixRE.MatchString(name):
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
		n := strings.Split(name, string(os.PathSeparator))
		n[len(n)-1] = suffixRE.ReplaceAllString(n[len(n)-1], "")
		testName := path.Join(n[1:]...)
		t.Run(testName, func(t *testing.T) {
			testRender(t, ctx, testName, name, timeout)
		})
	}
}

var suffixRE = regexp.MustCompile(`\.v[gl]\.json$`)

func testRender(t *testing.T, ctx context.Context, testName, name string, timeout time.Duration) {
	t.Helper()
	t.Parallel()
	spec, err := os.ReadFile(name)
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()
	vm := New(
		WithLogger(t.Log),
		WithVegaDemoData(),
	)
	start := time.Now()
	s, err := vm.Render(ctx, string(spec), "")
	total := time.Now().Sub(start)
	switch {
	case err != nil && contains(broken, testName):
		t.Logf("IGNORING: expected no error, got: %v", err)
		return
	case err != nil:
		t.Fatalf("expected no error, got: %v", err)
	}
	t.Logf("---\n%s\n---", s)
	t.Logf("duration: %s", total)
	if err := os.WriteFile(name+".svg", []byte(s), 0o644); err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}
}

func contains(v []string, s string) bool {
	for _, ss := range v {
		if ss == s {
			return true
		}
	}
	return false
}

var broken = []string{
	"lite/point_href",
	"lite/scatter_image",
	"vega/contour-plot",
	"vega/density-heatmaps",
	"vega/dorling-cartogram",
	"vega/warming-stripes",
}
