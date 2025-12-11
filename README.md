# Race condition / file caching bug

This demo repo shows a race condition when multiple files are passed to a module in parallel.
Though the files are different, occassionally the contents of one file is mixed with another.

## Running

```sh
❯ dagger call -m my-module-tests test-my-module

✔ connect 0.3s
✔ load module: my-module-tests 0.3s

✔ myModuleTests: MyModuleTests! 0.0s
✘ .testMyModule(
  ┆ fixtures: context /Users/christopher.palmer/workspace/dagger-caching-bug/my-module-tests/my-module-tests/fixtures (exclude: [])
  ): Void 1.3s ERROR
✘ MyModule.expectFileContents(
  ┆ expect: "d"
  ┆ file: Directory.file(path: "d"): File!
  ): Void 1.1s ERROR
! expected value d doesn't match actual contents c
```

Trace URL: https://dagger.cloud/nine/traces/986467472d6fa2f5a1eedbcf7e99a978

## Setup

The `my-module-tests` module has 6 files (`a`-`f`) each containing one character
which is the same as the file name.

`my-module-tests` calls `my-module`'s `ExpectContents` function, passing in each file and
expecting the contents to match the file's letter. Sometimes this function passes, other times
it does not.

## Conditions

The issue appears to occur on engine versions:

- 0.19.7
- 0.19.8

It doesn't occur on these engine versions:

- 0.19.3
- 0.19.4
- 0.19.5
- 0.19.6

Changing the code from this...

```go
// fixtures is a directory fetched from the host.
return dag.MyModule().ExpectFileContents(gctx, variant, fixtures.File(variant))
```

to this,

```go
return dag.MyModule().ExpectFileContents(gctx, variant, dag.File(variant, variant))
```

seems to make the issue go away.