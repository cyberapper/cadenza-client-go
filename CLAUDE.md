# cadenza-client-go

Auto-generated Go client SDK for the Cadenza API, generated from OpenAPI spec via `openapi-generator`.

## Commands

```bash
just doctor    # Full CI check: vet + build + test (run before every push)
just openapi   # Regenerate SDK from ../cadenza-docs/openapi/openapi.v3.yaml (override: CADENZA_DOCS_PATH)
just build     # Compile
just test      # Run tests
just vet       # Go vet
```

## Rules

- Run `just doctor` locally before pushing. All checks must pass.
- Commits follow Conventional Commits (`feat:`, `fix:`, `chore:`, etc.).
- Do not manually edit generated files — regenerate with `just openapi`.
- Hand-maintained files are listed in `.openapi-generator-ignore` (e.g., `ws/**`, `justfile`, `.github/**`).
- Versioning is managed by release-please — do not manually bump `go.mod` version.

## Structure

| Path | Description |
|------|-------------|
| `api/` | Generated API client modules |
| `ws/` | Custom WebSocket wrapper (hand-maintained) |
| `justfile` | Build/CI recipes |
| `.github/workflows/` | CI (go.yml) and release (release-please.yml) |
| `.openapi-generator-ignore` | Files protected from regeneration |
