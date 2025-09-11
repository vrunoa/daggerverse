package main

import (
	"context"

	"github.com/vrunoa/daggerverse/godot/internal/dagger"
)

const (
	godotImage = "ghcr.io/vrunoa/godot-ci:latest"
)

// Test running godot unit tests with Gut
func (m *Godot) Gut(ctx context.Context, src *dagger.Directory, testPath string) (*dagger.File, error) {
	cont, err := dag.Container().
		From(godotImage).
		WithWorkdir("/src").
		WithMountedDirectory("/src", src).
		WithEnvVariable("GOWORK", "off").
		WithExec([]string{
			"godot",
			"--headless",
			"--import",
		}).
		WithExec([]string{
			"godot",
			"-d",
			"-s",
			"--headless",
			"--path",
			"/src",
			"./addons/gut/gut_cmdln.gd",
			"-gdir",
			testPath,
			"-gexit",
			"-gjunit_xml_file",
			"coverage.xml",
		}).
		Sync(ctx)
	if err != nil {
		return nil, err
	}
	return cont.File("coverage.xml"), nil
}

func (m *Godot) XbfvGut(ctx context.Context, src *dagger.Directory, testPath string) (*dagger.File, error) {
	cont, err := dag.Container().
		From(godotImage).
		WithWorkdir("/src").
		WithMountedDirectory("/src", src).
		WithEnvVariable("GOWORK", "off").
		WithExec([]string{
			"godot",
			"--headless",
			"--import",
		}).
		WithExec([]string{
			"xvfb-run",
			"-a",
			"godot",
			"-d",
			"-s",
			"--headless",
			"--path",
			"/src",
			"./addons/gut/gut_cmdln.gd",
			"-gdir",
			testPath,
			"-gexit",
			"-gjunit_xml_file",
			"coverage.xml",
		}).
		Sync(ctx)
	if err != nil {
		return nil, err
	}
	return cont.File("coverage.xml"), nil
}

