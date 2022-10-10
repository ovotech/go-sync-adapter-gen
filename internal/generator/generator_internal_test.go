package generator

import (
	"bytes"
	"testing"

	"github.com/dave/jennifer/jen"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

//
type mockSnippet struct {
	mock.Mock
}

func (m *mockSnippet) snippet(f *jen.File, name string) {
	_ = m.Called(f, name)
}

// mockJen mocks out the ImportName and Save methods.
type mockJen struct {
	mock.Mock
}

func (m *mockJen) importName(f *jen.File, path, name string) {
	_ = m.Called(f, path, name)
}

func (m *mockJen) save(f *jen.File, filename string) error {
	args := m.Called(f, filename)

	return args.Error(0) //nolint:wrapcheck
}

func TestAdapter_Run(t *testing.T) {
	t.Parallel()

	newMockGenerator := new(mockSnippet)
	newMockGenerator.On("snippet", mock.Anything, "test").Run(func(args mock.Arguments) {
		args.Get(0).(*jen.File).Comment("Mock Function.")
	})

	newMockJen := new(mockJen)
	newMockJen.On("importName", mock.Anything, "foo", "bar")
	newMockJen.On("save", mock.Anything, "test/foo.go").Run(func(args mock.Arguments) {
		buf := &bytes.Buffer{}
		err := args.Get(0).(*jen.File).Render(buf)

		assert.NoError(t, err)
		assert.Equal(t, `package test

// Mock Function.
`, buf.String())
	}).Return(nil)

	adapter := New("test", WithGenerators(newMockGenerator.snippet), WithImport("foo", "bar"))
	adapter.importName = newMockJen.importName
	adapter.save = newMockJen.save

	err := adapter.Run("foo")

	assert.NoError(t, err)
}
