package main

import (
	"context"

	"github.com/vrunoa/daggerverse/go/internal/dagger"
)

const (
	goImage = "golang/:1.22"
)

// Test running go test
func (m *Go) Test(ctx context.Context, src *dagger.Directory) (*dagger.Container, error) {
	return dag.Container().
		From(golangciImage).
		WithWorkdir("/src").
		WithMountedDirectory("/src", src).
		WithEnvVariable("GOWORK", "off").
		WithExec([]string{"go", "test", "-shuffle=on", "-race", "-v"}).
		Sync(ctx)
}
