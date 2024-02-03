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
	vegaVer, liteVer, err := New().Version()
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}
	if v, exp := cleanString(vegaVer), cleanString(string(vegaVersionTxt)); v != exp {
		t.Errorf("expected %s, got: %s", exp, v)
	}
	if v, exp := cleanString(liteVer), cleanString(string(liteVersionTxt)); v != exp {
		t.Errorf("expected %s, got: %s", exp, v)
	}
	t.Logf("vega: %s vega-lite: %s", vegaVer, liteVer)
}

func TestCheckSpec(t *testing.T) {
	files := testFiles(t, func(name string) bool {
		return suffixRE.MatchString(name)
	})
	for _, nn := range files {
		name, testName := testName(nn)
		t.Run(testName, func(t *testing.T) {
			testCheckSpec(t, name)
		})
	}
}

func testCheckSpec(t *testing.T, name string) {
	t.Helper()
	t.Parallel()
	spec, err := os.ReadFile(name)
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}
	vm := New(WithLogger(t.Log))
	start := time.Now()
	res, err := vm.CheckSpec(string(spec))
	total := time.Since(start)
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}
	if os.Getenv("VERBOSE") != "" {
		t.Logf("---\n%s\n---", res)
	}
	t.Logf("duration: %s", total)
}

func TestCompile(t *testing.T) {
	files := testFiles(t, func(name string) bool {
		return strings.HasSuffix(name, ".vl.json")
	})
	for _, nn := range files {
		name, testName := testName(nn)
		t.Run(testName, func(t *testing.T) {
			testCompile(t, name)
		})
	}
}

func testCompile(t *testing.T, name string) {
	t.Helper()
	t.Parallel()
	spec, err := os.ReadFile(name)
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}
	vm := New(WithLogger(t.Log))
	start := time.Now()
	res, err := vm.Compile(string(spec))
	total := time.Since(start)
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}
	if os.Getenv("VERBOSE") != "" {
		t.Logf("---\n%s\n---", res)
	}
	t.Logf("duration: %s", total)
}

func TestParamExtract(t *testing.T) {
	files := testFiles(t, func(name string) bool {
		return strings.HasSuffix(name, ".vl.json")
	})
	for _, nn := range files {
		name, testName := testName(nn)
		t.Run(testName, func(t *testing.T) {
			testParamExtract(t, name)
		})
	}
}

func testParamExtract(t *testing.T, name string) {
	t.Helper()
	t.Parallel()
	spec, err := os.ReadFile(name)
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}
	vm := New(WithLogger(t.Log))
	switch params, err := vm.ParamExtract(string(spec)); {
	case err != nil:
		t.Fatalf("expected no error, got: %v", err)
	case len(params) != 0:
		t.Logf("params: %v", params)
	}
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
	files := testFiles(t, func(name string) bool {
		return suffixRE.MatchString(name)
	})
	for _, nn := range files {
		name, testName := testName(nn)
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
	opts := []Option{
		WithLogger(t.Log),
		WithPrefixedSourceDir("data/", "testdata/data/"),
	}
	if strings.HasPrefix(testName, "deneb/") {
		opts = append(opts, WithPrefixedSourceDir("deneb/", "testdata/deneb"))
	}
	vm := New(opts...)
	start := time.Now()
	res, err := vm.Render(ctx, string(spec))
	total := time.Since(start)
	switch {
	case err != nil && isBroken(testName):
		t.Logf("IGNORING: expected no error, got: %v", err)
		return
	case err != nil:
		t.Fatalf("expected no error, got: %v", err)
	}
	if os.Getenv("VERBOSE") != "" {
		t.Logf("---\n%s\n---", res)
	}
	t.Logf("duration: %s", total)
	if res = strings.TrimSpace(res); len(res) != 0 {
		if err := os.WriteFile(name+".svg", []byte(res), 0o644); err != nil {
			t.Fatalf("expected no error, got: %v", err)
		}
	}
}

func cleanString(s string) string {
	return strings.TrimPrefix(strings.TrimSpace(s), "v")
}

func isBroken(name string) bool {
	// these are all broken due to rendering image data as marks
	for _, ss := range []string{
		"compiled/scatter_image",
		"lite/scatter_image",
		"vega/contour-plot",
		"vega/density-heatmaps",
	} {
		if ss == name {
			return true
		}
	}
	return false
}

func testFiles(t *testing.T, f func(string) bool) []string {
	t.Helper()
	var files []string
	err := filepath.Walk("testdata", func(name string, fi fs.FileInfo, err error) error {
		switch {
		case err != nil:
			return err
		case fi.IsDir() || !f(name):
			return nil
		}
		files = append(files, name)
		return nil
	})
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}
	return files
}

func testName(name string) (string, string) {
	n := strings.Split(name, string(os.PathSeparator))
	n[len(n)-1] = suffixRE.ReplaceAllString(n[len(n)-1], "")
	return name, path.Join(n[1:]...)
}
