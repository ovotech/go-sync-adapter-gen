# Contributing to Go Sync Adapter Gen

## Preparation 🍳
We recommend asdf, it's our recommended way of managing our runtime CLIs:

1. [asdf](https://asdf-vm.com/)
2. [asdf-golang](https://github.com/kennyp/asdf-golang) (install via `asdf plugin-add golang`)

Alternatively install Go from the [official documentation](https://go.dev/doc/install).
The version of Go you want can be [found here](https://github.com/ovotech/go-sync/blob/main/go.mod#L3).

We also run the following tooling to ensure code quality:

1. [golangci-lint](https://golangci-lint.run/) for code quality.
2. [gci](https://github.com/daixiang0/gci) for consistent, deterministic imports.
   ```shell
   go install github.com/daixiang0/gci@latest
   ```

We run linters to ensure that code being checked in matches our quality standards, and have included a Makefile in this
repo containing common commands to assist with this.

| Command              | Description                         |
|----------------------|-------------------------------------|
| `make` / `make help` | Display list of available commands. |
| `make lint`          | Lint Go Sync.                       |
| `make fix`           | Fix some common linter errors.      |
| `make generate`      | Regenerate snapshots for testing.   |
