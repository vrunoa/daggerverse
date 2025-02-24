package main

import (
	"context"
	"strconv"

	"github.com/vrunoa/daggerverse/go/internal/dagger"
)

const (
	coverageImage = "ghcr.io/vrunoa/gocoverage:1.24"
)

// Coevrage - a simple threshold check for go coverage
func (m *Go) Coverage(ctx context.Context, src *dagger.Directory, coverfile string, threshold int) (*dagger.Container, error) {
	return dag.Container().
		From(coverageImage).
		WithWorkdir("/src").
		WithMountedDirectory("/src", src).
		WithEnvVariable("GOWORK", "off").
		WithExec([]string{"/usr/local/bin/coverage.sh", coverfile, strconv.Itoa(threshold)}).
		Sync(ctx)
}
