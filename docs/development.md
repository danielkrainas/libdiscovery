# libdiscovery Development

This is information related to developing libdiscovery.

## Tools

These are tools used for development.

### Vendoring

`gvt` is used to track and vendor dependencies, install with:

> $ go get github.com/FiloSottile/gvt

For details on usage, see [the project's github](https://github.com/FiloSottile/gvt).

## Building

Use `go` and build from the root of the project:

> $ go build

You can also use `make` to build all components:

> $ make compile

## Testing

Use `make` to run tests:

> $ make test

You can also use `go test` directly for any package without additional bootstrapping:

> $ go test ./locator/
