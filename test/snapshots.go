package main

import (
	"fmt"
	"log"
	"path/filepath"
	"reflect"
	"runtime"
	"strings"

	"github.com/dave/jennifer/jen"
	"github.com/ovotech/go-sync-adapter-gen/internal"
)

const adapter = "Test"

func getIndividualSnippets() []internal.Fn {
	var snips = make([]internal.Fn, 0, len(internal.Adapters)+len(internal.Tests)+len(internal.Examples))
	snips = append(snips, internal.Adapters...)
	snips = append(snips, internal.Tests...)
	snips = append(snips, internal.Examples...)

	return snips
}

func getFileName(fn internal.Fn) string {
	name := runtime.FuncForPC(reflect.ValueOf(fn).Pointer()).Name()
	name = strings.ReplaceAll(filepath.Base(name), ".", "_")

	return fmt.Sprintf("%s.txt", name)
}

func generateText(fn internal.Fn) *jen.File {
	packageName := strings.ToLower(adapter)

	file := jen.NewFile(packageName)
	fn(file, adapter)

	return file
}

func main() {
	for _, snip := range getIndividualSnippets() {
		name := getFileName(snip)

		err := internal.OutputFile(
			[]internal.Fn{snip},
			"test",
			"Test",
			fmt.Sprintf("test/snapshots/%s", name),
		)
		if err != nil {
			log.Fatal(err)
		}
	}
}
