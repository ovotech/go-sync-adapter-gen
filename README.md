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

Automatically scaffold new Go Sync adapters.

## Installation

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
package foobar

import (
	"context"
	"errors"
	"fmt"
	"github.com/ovotech/go-sync/pkg/ports"
)

// Ensure the adapter type fully satisfies the ports.Adapter interface.
var _ ports.Adapter = &FooBar{}

// ErrNotImplemented should be removed after implementation.
var ErrNotImplemented = errors.New("not_implemented")

type FooBar struct{}

// New instantiates a new adapter.
func New() *FooBar {
	return &FooBar{}
}

// Get a list of things.
func (f *FooBar) Get(_ context.Context) ([]string, error) {
	return nil, fmt.Errorf("foobar.get -> %w", ErrNotImplemented)
}

// Add things to your service.
func (f *FooBar) Add(_ context.Context, _ []string) error {
	return fmt.Errorf("foobar.add -> %w", ErrNotImplemented)
}

// Remove things from your service.
func (f *FooBar) Remove(_ context.Context, _ []string) error {
	return fmt.Errorf("foobar.remove -> %w", ErrNotImplemented)
}
```
</details>

<details>
<summary>foobar/foobar_internal_test.go</summary>

```go
package foobar

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNew(t *testing.T) {
	t.Parallel()
}

func TestFooBar_Get(t *testing.T) {
	t.Parallel()

	ctx := context.TODO()

	adapter := New()
	things, err := adapter.Get(ctx)

	assert.NoError(t, err)
	assert.ElementsMatch(t, things, []string{})
}

func TestFooBar_Add(t *testing.T) {
	t.Parallel()

	ctx := context.TODO()

	adapter := New()
	err := adapter.Add(ctx, []string{"foo"})

	assert.NoError(t, err)
}

func TestFooBar_Remove(t *testing.T) {
	t.Parallel()

	ctx := context.TODO()

	adapter := New()
	err := adapter.Remove(ctx, []string{"bar"})

	assert.NoError(t, err)
}
```

</details>

| *Made with 💚 by OVO's DevEx team* |
|------------------------------------|
