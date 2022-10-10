package snippets

import (
	"bytes"
	"testing"

	"github.com/dave/jennifer/jen"
	"github.com/stretchr/testify/assert"
)

func jenHelper(t *testing.T) (*jen.File, func() string) {
	t.Helper()

	file := jen.NewFile("test")
	buf := &bytes.Buffer{}

	return file, func() string {
		err := file.Render(buf)
		assert.NoError(t, err)

		return buf.String()
	}
}
