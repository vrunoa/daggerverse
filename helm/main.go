package main

import (
	"context"
	"github.com/vrunoa/daggerverse/helm/internal/dagger"
)

const (
	helmImage = "ghcr.io/vrunoa/kube:latest"
)

type Helm struct{}

func (h *Helm) Lint(
	ctx context.Context,
	src *dagger.Directory,
	svc string,
	values string,
) (*dagger.Container, error) {
	return dag.Container().
		From(helmImage).
		WithWorkdir("/src").
		WithMountedDirectory("/src", src).
		WithExec([]string{"helm", "lint", svc, "-f" , values}).
		Sync(ctx)
}
