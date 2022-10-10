package snippets

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTestNew(t *testing.T) {
	t.Parallel()

	f, out := jenHelper(t)

	TestNew(f, "Test")

	assert.Equal(t, `package test

import "testing"

func TestNew(t *testing.T) {
	t.Parallel()
}
`, out())
}

func TestTestAdapter_Get(t *testing.T) {
	t.Parallel()

	f, out := jenHelper(t)

	TestGet(f, "Test")

	assert.Equal(t, `package test

import (
	"context"
	assert "github.com/stretchr/testify/assert"
	"testing"
)

func TestTest_Get(t *testing.T) {
	t.Parallel()

	ctx := context.TODO()

	adapter := New()
	things, err := adapter.Get(ctx)

	assert.NoError(t, err)
	assert.ElementsMatch(t, things, []string{})
}
`, out())
}

func TestTestAdapter_Add(t *testing.T) {
	t.Parallel()

	f, out := jenHelper(t)

	TestAdd(f, "Test")

	assert.Equal(t, `package test

import (
	"context"
	assert "github.com/stretchr/testify/assert"
	"testing"
)

func TestTest_Add(t *testing.T) {
	t.Parallel()

	ctx := context.TODO()

	adapter := New()
	err := adapter.Add(ctx, []string{"foo"})

	assert.NoError(t, err)
}
`, out())
}

func TestTestAdapter_Remove(t *testing.T) {
	t.Parallel()

	f, out := jenHelper(t)

	TestRemove(f, "Test")

	assert.Equal(t, `package test

import (
	"context"
	assert "github.com/stretchr/testify/assert"
	"testing"
)

func TestTest_Remove(t *testing.T) {
	t.Parallel()

	ctx := context.TODO()

	adapter := New()
	err := adapter.Remove(ctx, []string{"bar"})

	assert.NoError(t, err)
}
`, out())
}
