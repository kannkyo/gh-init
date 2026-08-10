# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project overview

`gh-scaffolder` (module `github.com/kannkyo/gh-scaffolder`) is a [GitHub CLI extension](https://cli.github.com/manual/gh_extension). It is currently at template stage: `main.go` only prints a hello message and fetches the authenticated user via the GitHub REST API. No scaffolding logic, tests, or lint config exist yet — treat existing code as a starting point, not an established pattern to preserve.

Uses `github.com/cli/go-gh/v2` (the official gh extension SDK) as its only direct dependency. There is no CLI framework (no cobra) — flag/command handling is plain and will need to be introduced as functionality grows.

## Build & run

```
go build -o gh-scaffolder .
./gh-scaffolder
```

To run as an installed `gh` extension: `gh extension install .` from the repo root, then `gh scaffolder`.

## Testing & linting

No test files and no lint config exist yet. `golangci-lint` is not installed. Use `go vet ./...` and `gofmt` as the baseline until a proper test/lint setup is added.

## Release process

Releases are cut by pushing a tag matching `v*` (e.g. `git tag v0.1.0 && git push origin v0.1.0`). `.github/workflows/release.yml` uses `cli/gh-extension-precompile@v1` to cross-compile and publish the release automatically — there is no separate manual build/upload step. There is no CI workflow that runs on regular pushes/PRs (no build/test gate before merge).
