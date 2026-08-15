# gh-init

[![release](https://github.com/kannkyo/gh-init/actions/workflows/release.yml/badge.svg)](https://github.com/kannkyo/gh-init/actions/workflows/release.yml)

A [GitHub CLI extension](https://cli.github.com/manual/gh_extension) that generates common project files (LICENSE, `.gitignore`) from GitHub's official templates.

## Installation

```bash
gh extension install kannkyo/gh-init
```

## Usage

All commands use your existing `gh` authentication — no extra login step is required.

### `gh init license create <license-key>`

Fetches a license template from GitHub, fills in `[year]` (current year) and `[fullname]` (from `git config user.name`), and writes it to `./LICENSE`. Fails if `LICENSE` already exists.

```bash
gh init license create mit
gh init license create apache-2.0
```

`<license-key>` is a GitHub license identifier (e.g. `mit`, `apache-2.0`, `gpl-3.0`). See the [GitHub Licenses API](https://docs.github.com/en/rest/licenses) for the full list.

### `gh init ignore list`

Prints the names of all available `.gitignore` templates, one per line.

```bash
gh init ignore list
```

### `gh init ignore view <template>`

Prints the contents of a `.gitignore` template to stdout without writing any file.

```bash
gh init ignore view Go
```

### `gh init ignore create <template>`

Fetches a `.gitignore` template from GitHub and writes it to `./.gitignore`. Fails if `.gitignore` already exists.

```bash
gh init ignore create Go
```

`<template>` must match one of the names returned by `gh init ignore list`.

### `gh init ignore append <template>`

Fetches a `.gitignore` template from GitHub and appends it to the existing `./.gitignore` file. Fails if `.gitignore` does not exist (use `ignore create` first).

```bash
gh init ignore append Node
```

`<template>` must match one of the names returned by `gh init ignore list`.

## Development

See [CLAUDE.md](./CLAUDE.md) for build, test, and release instructions.
