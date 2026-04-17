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

## Status

- source-first and usable today
- built on the stable `xifty-ffi` ABI
- CI validates the wrapper against the public XIFty core repo on every push
- prepared for future module-distribution hardening

## License

MIT
