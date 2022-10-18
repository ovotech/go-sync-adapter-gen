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

<details>
<summary>foobar/foobar.go</summary>

```go
package example

import (
	"context"
	"fmt"
	gosync "github.com/ovotech/go-sync"
)

var (
	_ gosync.Adapter = &Example{} // Ensure Example fully satisfies the gosync.Adapter interface.
	_ gosync.InitFn  = Init       // Ensure the Init function fully satisfies the gosync.InitFn type.
)

type Example struct{}

// Get things in Example service.
func (e *Example) Get(_ context.Context) ([]string, error) {
	return nil, fmt.Errorf("example.get -> %w", gosync.ErrNotImplemented)
}

// Add things to Example service.
func (e *Example) Add(_ context.Context, _ []string) error {
	return fmt.Errorf("example.add -> %w", gosync.ErrNotImplemented)
}

// Remove things from Example service.
func (e *Example) Remove(_ context.Context, _ []string) error {
	return fmt.Errorf("example.remove -> %w", gosync.ErrNotImplemented)
}

// New Example gosync.adapter.
func New() *Example {
	return &Example{}
}

// Init a new Example gosync.Adapter.
func Init(config map[gosync.ConfigKey]string) (gosync.Adapter, error) {
	// Required ConfigKeys to initialise this adapter.
	for _, key := range []gosync.ConfigKey{} {
		if _, ok := config[key]; !ok {
			return nil, fmt.Errorf("example.init -> %w(%s)", gosync.ErrMissingConfig, key)
		}
	}

	return New(), nil
}
```
</details>

<details>
<summary>foobar/foobar_internal_test.go</summary>

```go
package example

import (
	"context"
	gosync "github.com/ovotech/go-sync"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNew(t *testing.T) {
	t.Parallel()
}

func TestExample_Get(t *testing.T) {
	t.Parallel()

	ctx := context.TODO()

	adapter := New()
	things, err := adapter.Get(ctx)

	assert.NoError(t, err)
	assert.ElementsMatch(t, things, []string{})
}

func TestExample_Add(t *testing.T) {
	t.Parallel()

	ctx := context.TODO()

	adapter := New()
	err := adapter.Add(ctx, []string{"foo"})

	assert.NoError(t, err)
}

func TestExample_Remove(t *testing.T) {
	t.Parallel()

	ctx := context.TODO()

	adapter := New()
	err := adapter.Remove(ctx, []string{"bar"})

	assert.NoError(t, err)
}

func TestInit(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		adapter, err := Init(map[gosync.ConfigKey]string{})

		assert.NoError(t, err)
		assert.IsType(t, &Example{}, adapter)
	})
}
```

</details>

### Made with 💚 by OVO Energy's DevEx team

<div>

![DevEx](./assets/devex.png)
![Platforms](./assets/platforms.png)
![Tools](./assets/tools.png)
![Golden Paths](./assets/golden_paths.png)
![Guard Rails](./assets/guard_rails.png)
![For You](./assets/for_you.png)

</div>
