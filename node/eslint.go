package main

import (
	"context"

	"github.com/vrunoa/daggerverse/node/internal/dagger"
)

func (m *Node) Eslint(
	ctx context.Context,
	src *dagger.Directory,
) (*dagger.Container, error) {
	return dag.Container().
		From("node:22.12.0-alpine3.21").
		WithWorkdir("/src").
		WithMountedDirectory("/src", src).
		WithEnvVariable("GOWORK", "off").
		WithExec([]string{"npm", "ci"}).
		WithExec([]string{"npx", "eslint", "."}).
		Sync(ctx)
}
