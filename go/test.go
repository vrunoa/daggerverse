package main

import (
	"context"

	"github.com/vrunoa/daggerverse/go/internal/dagger"
)

const (
	goImage = "golang/:1.22"
)

// Test running go test
func (m *Go) Test(ctx context.Context, src *dagger.Directory) (*dagger.File, error) {
	cont, err := dag.Container().
		From(golangciImage).
		WithWorkdir("/src").
		WithMountedDirectory("/src", src).
		WithEnvVariable("GOWORK", "off").
		WithExec([]string{"go", "test", "-coverprofile=coverage.out", "./...", "-shuffle=on", "-race", "-v"}).
		Sync(ctx)
	if err != nil {
		return nil, err
	}
	return cont.File("coverage.out"), nil
}
