package snippets

import (
	"fmt"
	"strings"

	"github.com/dave/jennifer/jen"
)

const (
	ensureAdapterInterfaceComment = "Ensure the adapter type fully satisfies the ports.Adapter interface."
	errNotImplementedComment      = "ErrNotImplemented should be removed after implementation."
	newComment                    = "New instantiates a new adapter."
	getComment                    = "Get a list of things."
	addComment                    = "Add things to your service."
	removeComment                 = "Remove things from your service."
)

func EnsureAdapterTypeSatisfiesInterface(f *jen.File, adapter string) {
	f.Comment(ensureAdapterInterfaceComment)
	f.Var().Id("_").Qual("github.com/ovotech/go-sync/pkg/ports", "Adapter").
		Op("=").Op("&").Id(adapter).Op("{}").Line()
}

func ErrNotImplemented(f *jen.File, _ string) {
	f.Comment(errNotImplementedComment)
	f.Var().Id("ErrNotImplemented").Op("=").Qual("errors", "New").Call(jen.Lit("not_implemented"))
	f.Line()
}

func EmptyAdapterStruct(f *jen.File, adapter string) {
	f.Type().Id(adapter).Struct()
}

func NewAdapter(f *jen.File, adapter string) {
	f.Comment(newComment)
	f.Func().Id("New").Params().Op("*").Id(adapter).Block(jen.Return(jen.Op("&").Id(adapter).Op("{}")))
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
		Call(jen.Lit(fmt.Sprintf("%s.%s -> %%w", packageName, method)), jen.Id("ErrNotImplemented"))
}

func GetMethod(f *jen.File, adapter string) { //nolint:varnamelen
	f.Comment(getComment)
	f.Func().Params(methodReceiver(adapter)).
		Id("Get").Params(context()).
		Params(jen.Op("[]").String(), jen.Error()).
		Block(jen.Return(jen.Nil(), wrappedNotImplemented(adapter, "get")))
	f.Line()
}

func AddMethod(f *jen.File, adapter string) { //nolint:varnamelen
	f.Comment(addComment)
	f.Func().Params(methodReceiver(adapter)).
		Id("Add").Params(context(), jen.Id("_").Op("[]").String()).
		Params(jen.Error()).
		Block(jen.Return(wrappedNotImplemented(adapter, "add")))
	f.Line()
}

func RemoveMethod(f *jen.File, adapter string) { //nolint:varnamelen
	f.Comment(removeComment)
	f.Func().Params(methodReceiver(adapter)).
		Id("Remove").Params(context(), jen.Id("_").Op("[]").String()).
		Params(jen.Error()).
		Block(jen.Return(wrappedNotImplemented(adapter, "remove")))
	f.Line()
}
