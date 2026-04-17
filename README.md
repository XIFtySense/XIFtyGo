# XIFty for Go

`XIFtyGo` is the official Go binding repo for XIFty.

It gives Go applications a focused `cgo` bridge into the XIFty metadata engine
so you can probe files and extract stable metadata views without shelling out to
external tools.

## What It Does

XIFty exposes four complementary metadata views:

- `raw`
- `interpreted`
- `normalized`
- `report`

This Go binding keeps those views intact and adds only a thin Go API surface.

## Quick Example

```go
output, err := xifty.Extract("photo.jpg", xifty.ViewNormalized)
if err != nil {
    panic(err)
}

fields := map[string]any{}
for _, fieldAny := range output["normalized"].(map[string]any)["fields"].([]any) {
    field := fieldAny.(map[string]any)
    fields[field["field"].(string)] = field["value"].(map[string]any)["value"]
}

fmt.Println(fields["device.make"])
fmt.Println(fields["captured_at"])
```

## API

- `Version()`
- `Probe(path string)`
- `Extract(path string, view ViewMode)`
- `ExtractNamed(path string, view string)`

## Why Use It

Use this binding when you want:

- native Go access to XIFty
- normalized metadata fields for application logic
- raw and interpreted metadata when provenance matters
- explicit error surfaces rather than lossy wrapper logic

Good fits include ingestion services, media pipelines, and upload-time metadata
inspection.

## Local Setup

This repo no longer assumes a sibling `../XIFty` checkout.

Prepare the core dependency into a repo-local cache:

```bash
bash scripts/prepare-core.sh
```

Then run the wrapper:

```bash
go test ./...
go run ./examples/basic_usage
go run ./examples/gallery_ingest
```

You can still override the core location explicitly with `XIFTY_CORE_DIR`.

This binding is intentionally still source-first. It is not yet on the newer
canonical runtime-artifact path used by the Python and Rust package hardening
work.

## Status

- source-first and usable today
- not yet on the canonical runtime-artifact packaging path
- built on the stable `xifty-ffi` ABI
- CI validates the wrapper against the public XIFty core repo

## License

MIT
