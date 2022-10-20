package internal

import (
	"strings"

	. "github.com/dave/jennifer/jen" //nolint:revive,stylecheck
)

func ExampleNew(jen *File, adapter string) {
	packageName := strings.ToLower(adapter)

	jen.Func().Id("ExampleNew").Params().Block(
		Id("adapter").Op(":=").Id(packageName).Dot("New").Call(),
		Line(),
		Qual("github.com/ovotech/go-sync", "New").Call(Id("adapter")),
	)
	jen.Line()
}

func ExampleInit(jen *File, adapter string) {
	packageName := strings.ToLower(adapter)
	jen.Func().Id("ExampleInit").Params().Block(
		Id("ctx").Op(":=").Qual("context", "Background").Call(),
		Line(),
		Id("adapter").Op(",").Id("err").Op(":=").Id(packageName).Dot("Init").Call(
			Id("ctx"), Map(Qual("github.com/ovotech/go-sync", "ConfigKey")).String().Values(
				Id(packageName).Dot(constExampleName).Op(":").Lit("example"),
			),
		),
		If(Id("err").Op("!=").Nil()).Block(
			Qual("log", "Fatal").Call(Id("err")),
		),
		Line(),
		Qual("github.com/ovotech/go-sync", "New").Call(Id("adapter")),
	)
	jen.Line()
}

var Examples = []Fn{ //nolint:gochecknoglobals
	ExampleNew,
	ExampleInit,
}
