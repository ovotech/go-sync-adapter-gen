package snippets

import (
	"fmt"
	"strings"

	"github.com/dave/jennifer/jen"
)

const (
	ensureAdapterInterfaceComment = "Ensure %s fully satisfies the gosync.Adapter interface."
	ensureInitFnComment           = "Ensure the Init function fully satisfies the gosync.InitFn type."
	requiredKeysComment           = "Required ConfigKeys to initialise this adapter."
	initFnComment                 = "Init a new %s gosync.Adapter."
	newComment                    = "New %s gosync.adapter."
	getComment                    = "Get things in %s service."
	addComment                    = "Add things to %s service."
	removeComment                 = "Remove things from %s service."
)

func EnsureTypesSatisfy(f *jen.File, adapter string) {
	f.Var().Defs(
		jen.Id("_").Qual("github.com/ovotech/go-sync", "Adapter").Op("=").Op("&").Id(adapter).Op("{}").
			Comment(fmt.Sprintf(ensureAdapterInterfaceComment, adapter)),
		jen.Id("_").Qual("github.com/ovotech/go-sync", "InitFn").Op("=").Id("Init").Comment(ensureInitFnComment),
	)
	f.Line()
}

func EmptyAdapterStruct(f *jen.File, adapter string) {
	f.Type().Id(adapter).Struct()
}

func NewAdapter(f *jen.File, adapter string) {
	f.Comment(fmt.Sprintf(newComment, adapter))
	f.Func().Id("New").Params().Op("*").Id(adapter).Block(jen.Return(jen.Op("&").Id(adapter).Op("{}")))
	f.Line()
}

func InitFn(f *jen.File, adapter string) { //nolint:varnamelen
	packageName := strings.ToLower(adapter)

	f.Comment(fmt.Sprintf(initFnComment, adapter))
	f.Func().Id("Init").Params(
		jen.Id("config").Map(jen.Qual("github.com/ovotech/go-sync", "ConfigKey")).String(),
	).Params(
		jen.Qual("github.com/ovotech/go-sync", "Adapter"),
		jen.Error(),
	).Block(
		jen.Comment(requiredKeysComment),
		jen.For(
			jen.Id("_").Op(",").Id("key").Op(":=").
				Range().Index().Qual("github.com/ovotech/go-sync", "ConfigKey").Values(),
		).Block(
			jen.If(
				jen.Id("_").Op(",").Id("ok").Op(":=").
					Id("config").Index(jen.Id("key")), jen.Op("!").Id("ok")).Block(
				jen.Return(jen.Nil(), jen.Qual("fmt", "Errorf").
					Call(
						jen.Lit(fmt.Sprintf("%s.init -> %%w(%%s)", packageName)),
						jen.Qual("github.com/ovotech/go-sync", "ErrMissingConfig"),
						jen.Id("key"),
					)),
			),
		),
		jen.Line(),
		jen.Return(jen.Id("New").Call(), jen.Nil()),
	)
	f.Line()
}

func methodReceiver(adapter string) *jen.Statement {
	receiver := strings.ToLower(string(adapter[0]))

	return jen.Id(receiver).Op("*").Id(adapter)
}

func context() *jen.Statement {
	return jen.Id("_").Qual("context", "Context")
}

func wrappedNotImplemented(adapter string, method string) *jen.Statement {
	packageName := strings.ToLower(adapter)

	return jen.Qual("fmt", "Errorf").
		Call(
			jen.Lit(fmt.Sprintf("%s.%s -> %%w", packageName, method)),
			jen.Qual("github.com/ovotech/go-sync", "ErrNotImplemented"),
		)
}

func GetMethod(f *jen.File, adapter string) { //nolint:varnamelen
	f.Comment(fmt.Sprintf(getComment, adapter))
	f.Func().Params(methodReceiver(adapter)).
		Id("Get").Params(context()).
		Params(jen.Op("[]").String(), jen.Error()).
		Block(jen.Return(jen.Nil(), wrappedNotImplemented(adapter, "get")))
	f.Line()
}

func AddMethod(f *jen.File, adapter string) { //nolint:varnamelen
	f.Comment(fmt.Sprintf(addComment, adapter))
	f.Func().Params(methodReceiver(adapter)).
		Id("Add").Params(context(), jen.Id("_").Op("[]").String()).
		Params(jen.Error()).
		Block(jen.Return(wrappedNotImplemented(adapter, "add")))
	f.Line()
}

func RemoveMethod(f *jen.File, adapter string) { //nolint:varnamelen
	f.Comment(fmt.Sprintf(removeComment, adapter))
	f.Func().Params(methodReceiver(adapter)).
		Id("Remove").Params(context(), jen.Id("_").Op("[]").String()).
		Params(jen.Error()).
		Block(jen.Return(wrappedNotImplemented(adapter, "remove")))
	f.Line()
}
