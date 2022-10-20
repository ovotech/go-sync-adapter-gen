# Go Sync Adapter Generator

<div align="center">

[![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/ovotech/go-sync-adapter-gen?label=go&logo=go)](go.mod)
[![GitHub release (latest SemVer)](https://img.shields.io/github/v/release/ovotech/go-sync-adapter-gen)](https://github.com/ovotech/go-sync-adapter-gen/releases)
[![Go Report Card](https://goreportcard.com/badge/github.com/ovotech/go-sync-adapter-gen?style=flat)](https://goreportcard.com/report/github.com/ovotech/go-sync-adapter-gen)
[![Go Reference](https://pkg.go.dev/badge/github.com/ovotech/go-sync-adapter-gen.svg)](https://pkg.go.dev/github.com/ovotech/go-sync-adapter-gen)
[![Test Status](https://github.com/ovotech/go-sync-adapter-gen/actions/workflows/test.yml/badge.svg)](https://github.com/ovotech/go-sync-adapter-gen/actions/workflows/test.yml)
[![GitHub issues](https://img.shields.io/github/issues/ovotech/go-sync-adapter-gen?style=flat)](https://github.com/ovotech/go-sync-adapter-gen/issues)
[![GitHub pull requests](https://img.shields.io/github/issues-pr/ovotech/go-sync-adapter-gen?label=pull+requests&style=flat)](https://github.com/ovotech/go-sync-adapter-gen/pull-requests)
[![License](https://img.shields.io/github/license/ovotech/go-sync-adapter-gen?style=flat)](/LICENSE)

</div>

Automatically scaffold new [Go Sync](https://github.com/ovotech/go-sync) adapters.

## Installation
You can install pre-built binaries for supported platforms on our [releases page.](https://github.com/ovotech/go-sync-adapter-gen/releases) 

Alternatively, install via Go:
```shell
go install github.com/ovotech/go-sync-adapter-gen@latest
```

## Usage
Go Sync Adapter Generator takes an adapter name (note: capitalisation is important) and outputs a folder containing
an adapter and tests. 

```shell
$ go-sync-adapter-gen FooBar
Successfully created: FooBar 🎉

$ ls foobar
foobar.go   foobar_internal_test.go
```

### Made with 💚 by OVO Energy's DevEx team

<div>

![DevEx](./assets/devex.png)
![Platforms](./assets/platforms.png)
![Tools](./assets/tools.png)
![Golden Paths](./assets/golden_paths.png)
![Guard Rails](./assets/guard_rails.png)
![For You](./assets/for_you.png)

</div>
