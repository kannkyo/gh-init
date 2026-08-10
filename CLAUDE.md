# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project overview

`gh-scaffolder` (module `github.com/kannkyo/gh-scaffolder`) is a [GitHub CLI extension](https://cli.github.com/manual/gh_extension) that generates common project files (LICENSE, `.gitignore`) from GitHub's template APIs. See [README.md](./README.md) for user-facing command usage.

`main.go` only builds and executes the cobra root command from `cmd/`. Command implementations live in `cmd/`:
- `cmd/root.go` — root command wiring
- `cmd/client.go` — `restGetter` interface wrapping `api.RESTClient.Get`, `newRESTClient` (swappable in tests), `formatAPIError` (maps 404s to `"<kind> \"<name>\" not found"`)
- `cmd/license.go` — `license create <key>`
- `cmd/ignore.go` — `ignore list` / `ignore view <template>` / `ignore create <template>`

Uses `github.com/cli/go-gh/v2` (the official gh extension SDK) for GitHub API access and `github.com/spf13/cobra` for CLI structure — these are the only two direct dependencies. GitHub API calls always go through the `restGetter` interface (not `api.RESTClient` directly) so commands can be unit tested with a fake client instead of hitting the network.

Output files (`LICENSE`, `.gitignore`) are never overwritten or appended to — commands error out if the target file already exists. There is no `--force` flag by design.

## Build & run

```
go build -o gh-scaffolder .
./gh-scaffolder
```

To run as an installed `gh` extension: `gh extension install .` from the repo root, then `gh scaffolder`.

## Testing & linting

`go test ./...` runs the unit test suite (`cmd/*_test.go`). Network-dependent code paths (the actual GitHub API calls) are excluded from tests via the `restGetter` fake in `cmd/client_test.go` — don't add tests that hit the real API.

`golangci-lint` is not installed. Use `go vet ./...` and `gofmt` as the baseline until a proper lint setup is added.

**Sandbox note:** in a sandboxed shell with network access disabled, `go get`/`go mod tidy` and any command needing to write to `GOMODCACHE`/`GOCACHE`/`~/.gnupg` (including `git commit` with GPG signing) will fail with `read-only file system` or network errors — rerun with the sandbox disabled in that case.

## Release process

Releases are cut by pushing a tag matching `v*` (e.g. `git tag v0.1.0 && git push origin v0.1.0`). `.github/workflows/release.yml` uses `cli/gh-extension-precompile@v1` to cross-compile and publish the release automatically — there is no separate manual build/upload step. There is no CI workflow that runs on regular pushes/PRs (no build/test gate before merge).
