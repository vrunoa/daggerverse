package main

import (
	"context"
)

const (
	golangciImage = "golangci/golangci-lint:v1.61-alpine"
)

// Lint - Lint the Go source code with golangci-lint
func (m *Go) Lint(ctx context.Context, src *Directory) (*Container, error) {
	cli := dag.Pipeline("golangci-lint")
	return cli.Container().
		From(golangciImage).
		WithWorkdir("/src").
		WithMountedDirectory("/src", src).
		WithEnvVariable("GOWORK", "off").
		WithExec([]string{"golangci-lint", "run", "/src/...", "-v", "--timeout=10m"}).
		Sync(ctx)
}
