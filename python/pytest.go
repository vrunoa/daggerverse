package main

import (
	"context"

	"github.com/vrunoa/daggerverse/python/internal/dagger"
)

const (
	pyImage = "python:3.11"
)

func (m *Python) Pytest(
	ctx context.Context,
	src *dagger.Directory,
) (*dagger.Container, error) {
	return dag.Container().
		From(pyImage).
		WithWorkdir("/src").
		WithMountedDirectory("/src", src).
		WithEnvVariable("GOWORK", "off").
		WithExec([]string{"pip", "install", "-e", "."}).
		WithExec([]string{"pytest"}).
		Sync(ctx)
}
