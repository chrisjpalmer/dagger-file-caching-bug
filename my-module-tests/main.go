// A generated module for MyModuleTests functions
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
	"dagger/my-module-tests/internal/dagger"

	"golang.org/x/sync/errgroup"
)

type MyModuleTests struct{}

// produces the caching bug
func (m *MyModuleTests) TestMyModule(
	ctx context.Context,
	// +defaultPath="/my-module-tests/fixtures"
	fixtures *dagger.Directory,
) error {
	errg, gctx := errgroup.WithContext(ctx)

	variants := []string{"a", "b", "c", "d", "e", "f"}

	for _, variant := range variants {
		errg.Go(func() error {

			// produces the bug
			// return dag.MyModule().ExpectFileContents(gctx, variant, fixtures.File(variant))

			// does not produce the bug
			return dag.MyModule().ExpectFileContents(gctx, variant, dag.File(variant, variant))
		})
	}

	return errg.Wait()
}

// doesn't produce the bug
func (m *MyModuleTests) TestMyModuleSeries(
	ctx context.Context,
	// +defaultPath="/my-module-tests/fixtures"
	fixtures *dagger.Directory,
) error {

	variants := []string{"a", "b", "c", "d", "e", "f"}

	for _, variant := range variants {
		err := dag.MyModule().ExpectFileContents(ctx, variant, fixtures.File(variant))
		if err != nil {
			return err
		}
	}

	return nil
}
