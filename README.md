# XIFtyGo

Go binding for [XIFty](https://github.com/XIFtySense/XIFty).

`XIFtyGo` is a focused `cgo` wrapper over the stable `xifty-ffi` C ABI. It is
ready for source-based use today and is intended to become the canonical Go
package for XIFty.

## What You Get

- `Version()` for the bound core version
- `Probe(path)` for format detection and structural inspection
- `Extract(path, view)` for the same JSON views exposed by the core engine
- a thin wrapper with no metadata logic duplicated in Go

## Quickstart

Clone the public core repo as a sibling checkout, then build the FFI library and
run the wrapper:

```bash
git clone git@github.com:XIFtySense/XIFty.git ../XIFty
cargo build -p xifty-ffi --manifest-path ../XIFty/Cargo.toml
go test ./...
go run ./examples/basic_usage
```

If your core checkout lives elsewhere, set both `XIFTY_CORE_DIR` and
`LD_LIBRARY_PATH` to point at it.

When you are consuming a tagged public release, the module path is:

```bash
go get github.com/XIFtySense/XIFtyGo@latest
```

## Status

- source-first and usable today
- built on the stable `xifty-ffi` ABI
- CI validates the wrapper against the public XIFty core repo on every push
- prepared for future module-distribution hardening

## Release Model

- Go modules are distributed through semver git tags
- v0 releases are the right default until the API settles
- consumers should install tagged versions rather than tracking random commits

## License

MIT
