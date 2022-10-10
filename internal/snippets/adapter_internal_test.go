package snippets

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEnsureAdapterTypeSatisfiesInterface(t *testing.T) {
	t.Parallel()

	f, out := jenHelper(t)

	EnsureAdapterTypeSatisfiesInterface(f, "Test")

	assert.Equal(t, fmt.Sprintf(`package test

import ports "github.com/ovotech/go-sync/pkg/ports"

// %s
var _ ports.Adapter = &Test{}
`, ensureAdapterInterfaceComment), out())
}

func TestErrNotImplemented(t *testing.T) {
	t.Parallel()

	f, out := jenHelper(t)

	ErrNotImplemented(f, "")

	assert.Equal(t, fmt.Sprintf(`package test

import "errors"

// %s
var ErrNotImplemented = errors.New("not_implemented")
`, errNotImplementedComment), out())
}

func TestEmptyAdapterStruct(t *testing.T) {
	t.Parallel()

	f, out := jenHelper(t)

	EmptyAdapterStruct(f, "Test")

	assert.Equal(t, `package test

type Test struct{}
`, out())
}

func TestNewAdapter(t *testing.T) {
	t.Parallel()

	f, out := jenHelper(t)

	NewAdapter(f, "Test")

	assert.Equal(t, fmt.Sprintf(`package test

// %s
func New() *Test {
	return &Test{}
}
`, newComment), out())
}

func TestGetMethod(t *testing.T) {
	t.Parallel()

	f, out := jenHelper(t)

	GetMethod(f, "Test")

	assert.Equal(t, fmt.Sprintf(`package test

import (
	"context"
	"fmt"
)

// %s
func (t *Test) Get(_ context.Context) ([]string, error) {
	return nil, fmt.Errorf("test.get -> %%w", ErrNotImplemented)
}
`, getComment), out())
}

func TestAddMethod(t *testing.T) {
	t.Parallel()

	f, out := jenHelper(t)

	AddMethod(f, "Test")

	assert.Equal(t, fmt.Sprintf(`package test

import (
	"context"
	"fmt"
)

// %s
func (t *Test) Add(_ context.Context, _ []string) error {
	return fmt.Errorf("test.add -> %%w", ErrNotImplemented)
}
`, addComment), out())
}

func TestRemoveMethod(t *testing.T) {
	t.Parallel()

	f, out := jenHelper(t)

	RemoveMethod(f, "Test")

	assert.Equal(t, fmt.Sprintf(`package test

import (
	"context"
	"fmt"
)

// %s
func (t *Test) Remove(_ context.Context, _ []string) error {
	return fmt.Errorf("test.remove -> %%w", ErrNotImplemented)
}
`, removeComment), out())
}
