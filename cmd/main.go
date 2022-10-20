package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/ovotech/go-sync-adapter-gen/internal"
)

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

	adapter := args[0]
	packageName := strings.ToLower(adapter)

	err := os.Mkdir(packageName, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}

	err = internal.OutputFile(
		internal.Adapters,
		packageName,
		adapter,
		fmt.Sprintf("%s/%s.go", packageName, packageName),
	)
	if err != nil {
		log.Fatal(err)
	}

	err = internal.OutputFile(
		internal.Tests,
		packageName,
		adapter,
		fmt.Sprintf("%s/%s_internal_test.go", packageName, packageName),
	)
	if err != nil {
		log.Fatal(err)
	}

	err = internal.OutputFile(
		internal.Examples,
		fmt.Sprintf("%s_test", packageName),
		adapter,
		fmt.Sprintf("%s/%s_test.go", packageName, packageName),
	)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Successfully created: %s 🎉", args[0]) //nolint:forbidigo
}
