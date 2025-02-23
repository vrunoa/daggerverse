package main

import (
	"context"

	"github.com/vrunoa/daggerverse/go/internal/dagger"
)

const (
	bufImage = "bufbuild/buf:1.50.0"
)


func (m *Go) Bufgen(
	ctx context.Context,
	src *dagger.Directory,
	target string,
) *dagger.Directory {
	return dag.Container().
		From(bufImage).
		WithWorkdir("/src").
		WithMountedDirectory("/src", src).
		WithEnvVariable("GOWORK", "off").
		WithExec([]string{"buf", "generate"}).
		Directory(target)
}
