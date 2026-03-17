# Cadenza Go Client SDK — Build & Test
#
# Usage:
#   just openapi     Regenerate SDK from OpenAPI spec
#   just build       Compile
#   just test        Run tests
#   just vet         Go vet
#   just doctor      Full CI check (vet + build + test)
#   just ci          Alias for doctor

set quiet
set shell := ["bash", "-eo", "pipefail", "-c"]

# OpenAPI spec path (override with CADENZA_DOCS_PATH env var)
docs_path := env("CADENZA_DOCS_PATH", justfile_directory() / "../cadenza-docs")
spec      := docs_path / "openapi/openapi.v3.yaml"

# Show available commands
[private]
default:
    @just --list

# Regenerate SDK from OpenAPI spec
openapi:
    #!/usr/bin/env bash
    set -eo pipefail
    echo "Generating Go client SDK from OpenAPI spec..."
    echo "Spec: {{spec}}"
    openapi-generator generate \
        -i "{{spec}}" \
        -g go \
        -o . \
        --package-name cadenza_client \
        --git-user-id cyberapper \
        --git-repo-id cadenza-client-go
    echo "Cleaning up unwanted generated files..."
    rm -f .travis.yml git_push.sh
    echo "Generation complete!"

alias o := openapi

# Compile the library
build:
    #!/usr/bin/env bash
    set -eo pipefail
    echo "Building..."
    go build -v ./...
    echo "Build successful!"

alias b := build

# Run tests
test:
    #!/usr/bin/env bash
    set -eo pipefail
    echo "Running tests..."
    go test -v ./...
    echo "All tests passed!"

alias t := test

# Run go vet
vet:
    #!/usr/bin/env bash
    set -eo pipefail
    echo "Running go vet..."
    go vet ./...
    echo "Vet passed!"

alias v := vet

# Full CI check (vet + build + test)
doctor:
    #!/usr/bin/env bash
    set -eo pipefail
    echo "=== Go vet ==="
    go vet ./...
    echo "=== Build ==="
    go build -v ./...
    echo "=== Tests ==="
    go test -v ./...
    echo "=== All checks passed! ==="

alias d := doctor
alias ci := doctor
