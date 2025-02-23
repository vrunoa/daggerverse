package main

import (
	"context"

	"github.com/vrunoa/daggerverse/go/internal/dagger"
)

const (
	mockeryImage = "vektra/mockery:v2.52"
)

func (m *Go) Mockery(
	ctx context.Context,
	src *dagger.Directory,
	target string,
) *dagger.Directory {
	return dag.Container().
		From(mockeryImage).
		WithWorkdir("/src").
		WithMountedDirectory("/src", src).
		WithEnvVariable("GOWORK", "off").
		WithExec([]string{"mockery"}).
		Directory(target)
}
