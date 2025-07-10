#!/bin/bash

# Build script with optimization flags for better performance

set -e

echo "Building optimized versions..."

# Build flags for optimization
BUILD_FLAGS="-ldflags=-s -w"
RACE_FLAGS="-race"
PROFILE_FLAGS="-gcflags=-cpuprofile=cpu.prof -gcflags=-memprofile=mem.prof"

# Build optimized version (smaller binary, stripped symbols)
echo "Building optimized binary..."
go build -ldflags="-s -w" -o bin/greeting-server-optimized examples/web_server_optimized.go

# Build with race detection for development
echo "Building with race detection..."
go build -race -o bin/greeting-server-race examples/web_server_optimized.go

# Build CLI optimized version
echo "Building optimized CLI..."
go build -ldflags="-s -w" -o bin/greeting-cli-optimized examples/cli_app.go

# Show binary sizes
echo ""
echo "Binary sizes:"
ls -lh bin/

# Run benchmarks
echo ""
echo "Running benchmarks..."
go test -bench=. -benchmem ./say_test.go ./say_optimized.go

echo ""
echo "Build complete!"
echo "Optimized server: bin/greeting-server-optimized"
echo "CLI tool: bin/greeting-cli-optimized"