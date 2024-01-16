package vegagoja

import (
	"io/fs"
	"os"
	"strings"
)

// Source is a cascading [fs.FS] implementation.
type Source struct {
	sources []fs.FS
}

// NewSource creates a source data source for the supplied file systems.
func NewSource(sources ...fs.FS) fs.FS {
	if len(sources) == 1 {
		return sources[0]
	}
	return &Source{
		sources: sources,
	}
}

// Open satisfies the [fs.FS] interface.
func (s *Source) Open(name string) (fs.File, error) {
	for _, source := range s.sources {
		if file, err := source.Open(name); err == nil {
			return file, nil
		}
	}
	return nil, os.ErrNotExist
}

// PrefixedSource is a prefixed source.
type PrefixedSource struct {
	prefix string
	source fs.FS
}

// NewPrefixedSource creates a prefixed source.
func NewPrefixedSource(prefix string, source fs.FS) *PrefixedSource {
	return &PrefixedSource{
		prefix: prefix,
		source: source,
	}
}

// NewPrefixedSourceDir creates a prefixed source for the specified directory.
func NewPrefixedSourceDir(prefix, dir string) *PrefixedSource {
	return &PrefixedSource{
		prefix: prefix,
		source: os.DirFS(dir),
	}
}

// Open satisfies the [fs.FS] interface.
func (s *PrefixedSource) Open(name string) (fs.File, error) {
	if !strings.HasPrefix(name, s.prefix) {
		return nil, os.ErrNotExist
	}
	return s.source.Open(strings.TrimPrefix(name, s.prefix))
}

// ResultSet is the shared interface for a result set.
type ResultSet interface {
	Next() bool
	Scan(...interface{}) error
	Columns() ([]string, error)
	Close() error
	Err() error
	NextResultSet() bool
}
