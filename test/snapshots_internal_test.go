package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func readInFixtureFile(path string) (string, error) {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err //nolint:wrapcheck
	}

	return string(file), nil
}

func TestSnapshots(t *testing.T) {
	t.Parallel()

	for _, snip := range getIndividualSnippets() {
		snip := snip
		name := getFileName(snip)
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			buf := &bytes.Buffer{}

			err := generateText(snip).Render(buf)

			assert.NoError(t, err)

			filename := fmt.Sprintf("./snapshots/%s", name)
			fixture, err := readInFixtureFile(filename)

			assert.NoError(t, err)
			assert.Equal(t, fixture, buf.String(), "Snapshots are out of date, regenerate with `make generate`.")
		})
	}
}
