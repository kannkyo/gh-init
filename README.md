# gh-scaffolder

A [GitHub CLI extension](https://cli.github.com/manual/gh_extension) that generates common project files (LICENSE, `.gitignore`) from GitHub's official templates.

## Installation

```
gh extension install kannkyo/gh-scaffolder
```

## Usage

All commands use your existing `gh` authentication — no extra login step is required.

### `gh scaffolder license create <license-key>`

Fetches a license template from GitHub, fills in `[year]` (current year) and `[fullname]` (from `git config user.name`), and writes it to `./LICENSE`. Fails if `LICENSE` already exists.

```
gh scaffolder license create mit
gh scaffolder license create apache-2.0
```

`<license-key>` is a GitHub license identifier (e.g. `mit`, `apache-2.0`, `gpl-3.0`). See the [GitHub Licenses API](https://docs.github.com/en/rest/licenses) for the full list.

### `gh scaffolder ignore list`

Prints the names of all available `.gitignore` templates, one per line.

```
gh scaffolder ignore list
```

### `gh scaffolder ignore view <template>`

Prints the contents of a `.gitignore` template to stdout without writing any file.

```
gh scaffolder ignore view Go
```

### `gh scaffolder ignore create <template>`

Fetches a `.gitignore` template from GitHub and writes it to `./.gitignore`. Fails if `.gitignore` already exists.

```
gh scaffolder ignore create Go
```

`<template>` must match one of the names returned by `gh scaffolder ignore list`.

## Development

See [CLAUDE.md](./CLAUDE.md) for build, test, and release instructions.
