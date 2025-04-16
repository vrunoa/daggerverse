package main

import (
	"context"

	"github.com/vrunoa/daggerverse/python/internal/dagger"
)

const (
	uvImage = "ghcr.io/vrunoa/uv:latest"
)

func (m *Python) Ruff(
	ctx context.Context,
	src *dagger.Directory,
) (*dagger.Container, error) {
	return dag.Container().
		From(uvImage).
		WithWorkdir("/src").
		WithMountedDirectory("/src", src).
		WithEnvVariable("GOWORK", "off").
		WithExec([]string{"uv", "tool", "run", "ruff", "check", "/src/"}).
		Sync(ctx)
}
