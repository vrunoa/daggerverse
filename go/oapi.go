package main

import (
	"context"

	"github.com/vrunoa/daggerverse/go/internal/dagger"
)

const (
	oapiImage = "ghcr.io/vrunoa/oapi-codegen:2.4.1"
)

// Codegen generates code from an OpenAPI specs using oapi-codegen.
func (m *Go) Codegen(
	ctx context.Context,
	src *dagger.Directory,
	config string,
	spec string,
	target string,
) *dagger.Directory {
	return dag.Container().
		From(oapiImage).
		WithWorkdir("/src").
		WithMountedDirectory("/src", src).
		WithEnvVariable("GOWORK", "off").
		WithExec([]string{"oapi-codegen", "--config", config, spec}).
		Directory(target)
}
