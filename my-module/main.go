package main

import (
	"context"
	"dagger/my-module/internal/dagger"
	"fmt"
	"math/rand"
	"time"
)

type MyModule struct{}

func (m *MyModule) ExpectFileContents(ctx context.Context, expect string, file *dagger.File) error {
	rd := rand.New(rand.NewSource(time.Now().Unix()))

	time.Sleep(time.Duration(rd.Int31n(1000)) * time.Millisecond)

	contents, err := dag.Container().
		From("alpine").
		WithEnvVariable("CACHE_BUSTER", time.Now().String()).
		WithFile("my-file", file).
		WithExec([]string{"cat", "my-file"}).
		Stdout(ctx)
	if err != nil {
		return err
	}

	if expect != contents {
		return fmt.Errorf("expected value %s doesn't match actual contents %s", expect, contents)
	}

	return nil
}
