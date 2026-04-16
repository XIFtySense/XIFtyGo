# XIFtyGo

Go package for XIFty.

This package currently links against the XIFty core repository through the
stable `xifty-ffi` C ABI. Local development expects a sibling checkout of the
core repo at:

- `../XIFty`

Before running tests or examples, build the core FFI library:

```bash
cargo build -p xifty-ffi --manifest-path ../XIFty/Cargo.toml
go test ./...
go run ./examples/basic_usage
```

