// A generated module for Go functions
//
// This module has been generated via dagger init and serves as a reference to
// basic module structure as you get started with Dagger.
//
// Two functions have been pre-created. You can modify, delete, or add to them,
// as needed. They demonstrate usage of arguments and return types using simple
// echo and grep commands. The functions can be called from the dagger CLI or
// from one of the SDKs.
//
// The first line in this comment block is a short description line and the
// rest is a long description with more detail on the module's purpose or usage,
// if appropriate. All modules should have a short description.

package main

import (
	"context"
	"dagger/go/internal/dagger"
)

const (
	golangciImage = "golangci/golangci-lint:v1.61-alpine"
)

type Go struct{}

// Returns a container that echoes whatever string argument is provided
func (m *Go) Lint(ctx context.Context, src *dagger.Directory) (string, error) {
	return dag.Container().
		From(golangciImage).
		WithWorkdir("/src").
		WithMountedDirectory("/src", src).
		WithEnvVariable("GOWORK", "off").
		WithExec([]string{"golangci-lint", "run", "/src/...", "-v", "--timeout=10m"}).
		Stdout(ctx)
}
