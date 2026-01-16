package repository

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"
)

type QueryLoader struct {
	basePath string
	cache    map[string]string
	mu       sync.RWMutex
}

func newQueryLoader(basePath string) (*QueryLoader, error) {
	ql := QueryLoader{
		basePath: basePath,
		cache:    make(map[string]string),
	}

	if err := ql.loadAll(); err != nil {
		return nil, err
	}

	return &ql, nil
}

func (ql *QueryLoader) loadAll() error {
	return filepath.Walk(ql.basePath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() || filepath.Ext(path) != ".cypher" {
			return nil
		}

		contents, err := os.ReadFile(path)
		if err != nil {
			return err
		}

		ql.cache[filepath.Base(path)] = string(contents)
		return nil
	})
}

func (ql *QueryLoader) Get(name string) (string, error) {
	ql.mu.RLock()
	defer ql.mu.RUnlock()

	query, ok := ql.cache[name]
	if !ok {
		return "", fmt.Errorf("cypher query not found: %s", name)
	}

	return query, nil
}
