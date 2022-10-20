package internal

import (
	"fmt"
	"strings"

	. "github.com/dave/jennifer/jen" //nolint:revive,stylecheck
)

const (
	packageComment = `Package %s synchronises things with %s.

TODO: Write documentation.
`
	constExampleName              = "AnExampleConfig"
	constExampleAboveComment      = "%s is an example config for use with [%s.Init]."
	constExampleInlineComment     = "TODO: Remove and replace with config."
	ensureAdapterInterfaceComment = "Ensure [%s.%s] fully satisfies the [gosync.Adapter] interface."
	ensureInitFnComment           = "Ensure the [%s.Init] function fully satisfies the [gosync.InitFn] type."
	requiredKeysComment           = "Required [gosync.ConfigKeys] to initialise this adapter."
	getComment                    = "Get things in %s service."
	addComment                    = "Add things to %s service."
	removeComment                 = "Remove things from %s service."
	newComment                    = "New %s [gosync.adapter]."
	initFnComment                 = `Init a new %s [gosync.Adapter].

Required config:
  - [%s.%s]
`
)

func PackageComment(file *File, adapter string) {
	packageName := strings.ToLower(adapter)

	file.PackageComment(fmt.Sprintf(packageComment, packageName, adapter))
}

func ExampleConfig(file *File, adapter string) {
	packageName := strings.ToLower(adapter)

	file.Comment(fmt.Sprintf(constExampleAboveComment, constExampleName, packageName))
	file.Const().Id(constExampleName).Qual("github.com/ovotech/go-sync", "ConfigKey").Op("=").
		Lit(fmt.Sprintf("%s_example_config", packageName)).Comment(constExampleInlineComment)
	file.Line()
}

func EnsureTypesSatisfy(file *File, adapter string) {
	packageName := strings.ToLower(adapter)

	file.Var().Defs(
		Id("_").Qual("github.com/ovotech/go-sync", "Adapter").Op("=").Op("&").Id(adapter).Op("{}").
			Comment(fmt.Sprintf(ensureAdapterInterfaceComment, packageName, adapter)),
		Id("_").Qual("github.com/ovotech/go-sync", "InitFn").Op("=").Id("Init").
			Comment(fmt.Sprintf(ensureInitFnComment, packageName)),
	)
	file.Line()
}

func EmptyAdapterStruct(jen *File, adapter string) {
	jen.Type().Id(adapter).Struct()
}

func methodReceiver(adapter string) *Statement {
	receiver := strings.ToLower(string(adapter[0]))

	return Id(receiver).Op("*").Id(adapter)
}

func context() *Statement {
	return Id("_").Qual("context", "Context")
}

func wrappedNotImplemented(adapter string, method string) *Statement {
	packageName := strings.ToLower(adapter)

	return Qual("fmt", "Errorf").
		Call(
			Lit(fmt.Sprintf("%s.%s -> %%w", packageName, method)),
			Qual("github.com/ovotech/go-sync", "ErrNotImplemented"),
		)
}

func GetMethod(file *File, adapter string) {
	file.Comment(fmt.Sprintf(getComment, adapter))
	file.Func().Params(methodReceiver(adapter)).
		Id("Get").Params(context()).
		Params(Op("[]").String(), Error()).
		Block(Return(Nil(), wrappedNotImplemented(adapter, "get")))
	file.Line()
}

func AddMethod(file *File, adapter string) {
	file.Comment(fmt.Sprintf(addComment, adapter))
	file.Func().Params(methodReceiver(adapter)).
		Id("Add").Params(context(), Id("_").Op("[]").String()).
		Params(Error()).
		Block(Return(wrappedNotImplemented(adapter, "add")))
	file.Line()
}

func RemoveMethod(file *File, adapter string) {
	file.Comment(fmt.Sprintf(removeComment, adapter))
	file.Func().Params(methodReceiver(adapter)).
		Id("Remove").Params(context(), Id("_").Op("[]").String()).
		Params(Error()).
		Block(Return(wrappedNotImplemented(adapter, "remove")))
	file.Line()
}

func NewAdapter(jen *File, adapter string) {
	jen.Comment(fmt.Sprintf(newComment, adapter))
	jen.Func().Id("New").Params().Op("*").Id(adapter).Block(Return(Op("&").Id(adapter).Op("{}")))
	jen.Line()
}

func InitFn(file *File, adapter string) {
	packageName := strings.ToLower(adapter)

	file.Comment(fmt.Sprintf(initFnComment, adapter, packageName, constExampleName))
	file.Func().Id("Init").Params(
		context(),
		Id("config").Map(Qual("github.com/ovotech/go-sync", "ConfigKey")).String(),
	).Params(
		Qual("github.com/ovotech/go-sync", "Adapter"),
		Error(),
	).Block(
		Comment(requiredKeysComment),
		For(
			Id("_").Op(",").Id("key").Op(":=").
				Range().Index().Qual("github.com/ovotech/go-sync", "ConfigKey").Values(Id(constExampleName)),
		).Block(
			If(
				Id("_").Op(",").Id("ok").Op(":=").
					Id("config").Index(Id("key")), Op("!").Id("ok")).Block(
				Return(Nil(), Qual("fmt", "Errorf").
					Call(
						Lit(fmt.Sprintf("%s.init -> %%w(%%s)", packageName)),
						Qual("github.com/ovotech/go-sync", "ErrMissingConfig"),
						Id("key"),
					)),
			),
		),
		Line(),
		Return(Id("New").Call(), Nil()),
	)
	file.Line()
}

var Adapters = []Fn{ //nolint:gochecknoglobals
	PackageComment,
	ExampleConfig,
	EnsureTypesSatisfy,
	EmptyAdapterStruct,
	GetMethod,
	AddMethod,
	RemoveMethod,
	NewAdapter,
	InitFn,
}
