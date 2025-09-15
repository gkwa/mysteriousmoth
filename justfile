# Show available rules
default:
    @just --list

# Build the project
build:
    go build -o mysteriousmoth .

# Test by parsing the example markdown files
test: build
    ./mysteriousmoth example.md
    @echo ""
    ./mysteriousmoth example-toml.md

# Clean up build artifacts
teardown:
    rm -f mysteriousmoth
    go clean
