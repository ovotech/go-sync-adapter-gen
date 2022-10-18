package snippets

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEnsureAdapterTypeSatisfiesInterface(t *testing.T) {
	t.Parallel()

	f, out := jenHelper(t)

	EnsureTypesSatisfy(f, "Test")

	assert.Equal(t, fmt.Sprintf(`package test

import gosync "github.com/ovotech/go-sync"

var (
	_ gosync.Adapter = &Test{} // %s
	_ gosync.InitFn  = Init    // %s
)
`, fmt.Sprintf(ensureAdapterInterfaceComment, "Test"), ensureInitFnComment), out())
}

func TestEmptyAdapterStruct(t *testing.T) {
	t.Parallel()

	f, out := jenHelper(t)

	EmptyAdapterStruct(f, "Test")

	assert.Equal(t, `package test

type Test struct{}
`, out())
}

func TestGetMethod(t *testing.T) {
	t.Parallel()

	f, out := jenHelper(t)

	GetMethod(f, "Test")

	assert.Equal(t, fmt.Sprintf(`package test

import (
	"context"
	"fmt"
	gosync "github.com/ovotech/go-sync"
)

// %s
func (t *Test) Get(_ context.Context) ([]string, error) {
	return nil, fmt.Errorf("test.get -> %%w", gosync.ErrNotImplemented)
}
`, fmt.Sprintf(getComment, "Test")), out())
}

func TestAddMethod(t *testing.T) {
	t.Parallel()

	f, out := jenHelper(t)

	AddMethod(f, "Test")

	assert.Equal(t, fmt.Sprintf(`package test

import (
	"context"
	"fmt"
	gosync "github.com/ovotech/go-sync"
)

// %s
func (t *Test) Add(_ context.Context, _ []string) error {
	return fmt.Errorf("test.add -> %%w", gosync.ErrNotImplemented)
}
`, fmt.Sprintf(addComment, "Test")), out())
}

func TestRemoveMethod(t *testing.T) {
	t.Parallel()

	f, out := jenHelper(t)

	RemoveMethod(f, "Test")

	assert.Equal(t, fmt.Sprintf(`package test

import (
	"context"
	"fmt"
	gosync "github.com/ovotech/go-sync"
)

// %s
func (t *Test) Remove(_ context.Context, _ []string) error {
	return fmt.Errorf("test.remove -> %%w", gosync.ErrNotImplemented)
}
`, fmt.Sprintf(removeComment, "Test")), out())
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
`, fmt.Sprintf(newComment, "Test")), out())
}

func TestInitFn(t *testing.T) {
	t.Parallel()

	f, out := jenHelper(t)

	InitFn(f, "Test")

	assert.Equal(t, fmt.Sprintf(`package test

import (
	"fmt"
	gosync "github.com/ovotech/go-sync"
)

// %s
func Init(config map[gosync.ConfigKey]string) (gosync.Adapter, error) {
	// %s
	for _, key := range []gosync.ConfigKey{} {
		if _, ok := config[key]; !ok {
			return nil, fmt.Errorf("test.init -> %%w(%%s)", gosync.ErrMissingConfig, key)
		}
	}

	return New(), nil
}
`, fmt.Sprintf(initFnComment, "Test"), requiredKeysComment), out())
}
