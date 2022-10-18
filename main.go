package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/ovotech/go-sync-adapter-gen/internal/generator"
	"github.com/ovotech/go-sync-adapter-gen/internal/snippets"
)

func exec(name string) error {
	lowerCaseAdapterName := strings.ToLower(name)

	err := os.Mkdir(lowerCaseAdapterName, os.ModePerm)
	if err != nil {
		return fmt.Errorf("mkdir -> %w", err)
	}

	err = generator.New(
		name,
		generator.WithGenerators(
			snippets.EnsureTypesSatisfy,
			snippets.EmptyAdapterStruct,
			snippets.GetMethod,
			snippets.AddMethod,
			snippets.RemoveMethod,
			snippets.NewAdapter,
			snippets.InitFn,
		),
	).Run(lowerCaseAdapterName)
	if err != nil {
		return fmt.Errorf("adapter -> %w", err)
	}

	err = generator.New(
		name,
		generator.WithImport("github.com/stretchr/testify/assert", "assert"),
		generator.WithGenerators(
			snippets.TestNew,
			snippets.TestGet,
			snippets.TestAdd,
			snippets.TestRemove,
			snippets.TestInit,
		),
	).Run(fmt.Sprintf("%s_internal_test", lowerCaseAdapterName))
	if err != nil {
		return fmt.Errorf("tests -> %w", err)
	}

	return nil
}

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		fmt.Println("You must enter an adapter name e.g. go-sync-adapter-gen Foo.") //nolint:forbidigo
		os.Exit(1)
	}

	if strings.ToLower(args[0]) == args[0] {
		fmt.Println("Your adapter must start with an upper-case character e.g. Foo.") //nolint:forbidigo
		os.Exit(1)
	}

	err := exec(args[0])
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Successfully created: %s 🎉", args[0]) //nolint:forbidigo
}
