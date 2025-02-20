package main

import (
	"context"

	"github.com/vrunoa/daggerverse/go/internal/dagger"
)

const (
	openapitoolsImage = "openapitools/openapi-generator-cli:latest"
)

// Validate - validate openapi spec
func (m *Go) OpenapiValidate(ctx context.Context, src *dagger.Directory, spec string) (string, error) {
	return dag.Container().
		From(openapitoolsImage).
		WithWorkdir("/src").
		WithMountedDirectory("/src", src).
		WithExec([]string{"/usr/local/bin/docker-entrypoint.sh", "validate", "-i", spec}).
		Stdout(ctx)
}
