package main

import (
	"context"
	"dagger/my-module/internal/dagger"
	"time"
)

type MyModule struct{}

func (m *MyModule) ExpectFileContents(ctx context.Context, expect string, file *dagger.File) (string, error) {
	contents, err := dag.Container().
		From("alpine").
		WithEnvVariable("CACHE_BUSTER", time.Now().String()).
		WithFile("my-file", file).
		WithExec([]string{"touch", "out"}).
		WithExec([]string{"sh", "-c", "echo expected: " + expect + " > out"}).
		WithExec([]string{"sh", "-c", "echo actual: $(cat my-file) >> out"}).
		File("out").
		Contents(ctx)
	if err != nil {
		return "", err
	}

	return contents, nil
}
