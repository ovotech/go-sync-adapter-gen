package internal

import "github.com/dave/jennifer/jen"

func OutputFile(snippets []Fn, packageName string, adapter string, path string) error {
	file := jen.NewFile(packageName)

	for _, snippet := range snippets {
		snippet(file, adapter)
	}

	return file.Save(path) //nolint:wrapcheck
}
