# filtgen-example

This is an example of how to use the `filtgen`.

## Tasks

We recommend that this section be run with [`xc`](https://github.com/joerdav/xc).

### install:local

```sh
cd ../
go install -installsuffix filtgen
```

### intall:latest-release

```sh
go install github.com/miyamo2/filtgen@latest
```

### setup:goenv

Set `GOEXPERIMENT` to `rangefunc` if Go version is 1.22.

```sh
GOVER=$(go mod graph)
if [[ $GOVER == *"go@1.22"* ]]; then
  go env -w GOEXPERIMENT=rangefunc
fi
```

### run:go-generate

```sh
go generate ./...
```
