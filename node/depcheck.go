package main

import (
	"context"

	"github.com/vrunoa/daggerverse/node/internal/dagger"
)

func (m *Node) Depcheck(
	ctx context.Context,
	src *dagger.Directory,
) (*dagger.Container, error) {
	return dag.Container().
		From("node:20.19.1-alpine3.21").
		WithWorkdir("/src").
		WithMountedDirectory("/src", src).
		WithEnvVariable("GOWORK", "off").
		WithExec([]string{"npx", "depcheck"}).
		Sync(ctx)
}
