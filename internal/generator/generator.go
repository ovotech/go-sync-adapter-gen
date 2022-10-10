package generator

import (
	"fmt"
	"strings"

	"github.com/dave/jennifer/jen"
)

type Fn = func(*jen.File, string)

type Generator struct {
	name       string
	generators []Fn
	imports    map[string]string

	importName func(f *jen.File, path, name string)
	save       func(f *jen.File, filename string) error
}

func WithGenerators(generators ...Fn) func(g *Generator) {
	return func(g *Generator) {
		g.generators = generators
	}
}

func WithImport(path string, name string) func(g *Generator) {
	return func(g *Generator) {
		g.imports[path] = name
	}
}

func importName(f *jen.File, path, name string) {
	f.ImportName(path, name)
}

func save(f *jen.File, filename string) error {
	return f.Save(filename) //nolint:wrapcheck
}

// New code generator.
func New(name string, optsFn ...func(g *Generator)) *Generator {
	adapter := &Generator{
		name:       name,
		generators: make([]Fn, 0),
		imports:    map[string]string{},

		importName: importName,
		save:       save,
	}

	for _, fn := range optsFn {
		fn(adapter)
	}

	return adapter
}

// Run the code generator.
func (g *Generator) Run(filename string) error {
	packageName := strings.ToLower(g.name)

	f := jen.NewFile(packageName) //nolint:varnamelen

	for path, name := range g.imports {
		g.importName(f, path, name)
	}

	for _, fn := range g.generators {
		fn(f, g.name)
	}

	err := g.save(f, fmt.Sprintf("%s/%s.go", packageName, filename))
	if err != nil {
		return fmt.Errorf("generator.run.save -> %w", err)
	}

	return nil
}
