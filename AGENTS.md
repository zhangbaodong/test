## Cursor Cloud specific instructions

This is a Go library (`github.com/zhangbaodong/test`) providing greeting functionality. No external dependencies; uses only the Go standard library.

### Known build issue

`say.go` and `say_optimized.go` both declare `SayHi` in `package test` without build tags. Running `go test ./...` or `go build ./...` at the package level **will fail** with a redeclaration error. Use explicit file lists as `build.sh` does:

```bash
# Run benchmarks with the optimized implementation
go test -v -bench=. -benchmem ./say_test.go ./say_optimized.go

# Run benchmarks with the original implementation
go test -v -bench=. -benchmem ./say_test.go ./say.go
```

The `examples/` programs import the package and therefore also fail to compile via `go run`. Use the pre-compiled binaries (`test-app`, `test-app-optimized`) to run the web server on port 8080.

### Running the web server

```bash
./test-app          # starts on :8080
# or
./test-app-optimized  # optimized build, also :8080
```

Endpoints: `/` (HTML form), `/greet?name=X`, `/api/greet?name=X` (JSON), `/api/simple?name=X` (plain text).

### Building optimized binaries

Run `bash build.sh` — it builds into `bin/` and runs benchmarks.
