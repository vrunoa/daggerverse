package main

import (
	"context"
	"path"

	"github.com/vrunoa/daggerverse/go/internal/dagger"
)

func (m *Go) CoverTreemap(ctx context.Context, src *dagger.Directory, coverfile string, out string) (*dagger.File, error) {
	ctr := dag.Container().
		From(coverageImage).
		WithWorkdir("/src").
		WithMountedDirectory("/src", src).
		WithEnvVariable("GOWORK", "off")

	output, err := ctr.WithExec([]string{"go-cover-treemap", "-coverprofile", coverfile}).
		Stdout(ctx)
	if err != nil {
		return nil, err
	}
	svgFile := path.Join("/src", out)
	ctr = ctr.WithNewFile(svgFile, output)

	return ctr.File(out), nil
}
