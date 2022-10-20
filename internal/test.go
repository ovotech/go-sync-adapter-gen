package internal

import (
	"fmt"

	. "github.com/dave/jennifer/jen" //nolint:revive,stylecheck
)

func getTestParams() *Statement {
	return Id("t").Op("*").Qual("testing", "T")
}

func TestImports(file *File, _ string) {
	file.ImportName("github.com/stretchr/testify/assert", "assert")
}

func TestNew(file *File, _ string) {
	file.Func().Id("TestNew").Params(Id("t").Op("*").Qual("testing", "T")).Block(
		Id(string('t')).Dot("Parallel").Call(),
	)
	file.Line()
}

func TestGet(file *File, adapter string) {
	file.Func().Id(fmt.Sprintf("Test%s_Get", adapter)).Params(getTestParams()).Block(
		Id(string('t')).Dot("Parallel").Call(),
		Line(),
		Id("ctx").Op(":=").Qual("context", "TODO").Call(),
		Line(),
		Id("adapter").Op(":=").Id("New").Call(),
		Id("things").Op(",").Id("err").Op(":=").Id("adapter").Dot("Get").Call(Id("ctx")),
		Line(),
		Qual("github.com/stretchr/testify/assert", "NoError").Call(Id("t"), Id("err")),
		Qual("github.com/stretchr/testify/assert", "ElementsMatch").
			Call(Id("t"), Id("things"), Op("[]").String().Op("{}")),
	)
	file.Line()
}

func TestAdd(file *File, adapter string) {
	file.Func().Id(fmt.Sprintf("Test%s_Add", adapter)).Params(getTestParams()).Block(
		Id(string('t')).Dot("Parallel").Call(),
		Line(),
		Id("ctx").Op(":=").Qual("context", "TODO").Call(),
		Line(),
		Id("adapter").Op(":=").Id("New").Call(),
		Id("err").Op(":=").Id("adapter").Dot("Add").Call(Id("ctx"), Op("[]").String().Op("{\"foo\"}")),
		Line(),
		Qual("github.com/stretchr/testify/assert", "NoError").Call(Id("t"), Id("err")),
	)
	file.Line()
}

func TestRemove(file *File, adapter string) {
	file.Func().Id(fmt.Sprintf("Test%s_Remove", adapter)).Params(getTestParams()).Block(
		Id(string('t')).Dot("Parallel").Call(),
		Line(),
		Id("ctx").Op(":=").Qual("context", "TODO").Call(),
		Line(),
		Id("adapter").Op(":=").Id("New").Call(),
		Id("err").Op(":=").Id("adapter").Dot("Remove").Call(Id("ctx"), Op("[]").String().Op("{\"bar\"}")),
		Line(),
		Qual("github.com/stretchr/testify/assert", "NoError").Call(Id("t"), Id("err")),
	)
	file.Line()
}

func TestInit(file *File, adapter string) {
	file.Func().Id("TestInit").Params(getTestParams()).Block(
		Id(string('t')).Dot("Parallel").Call(),
		Line(),
		Id("ctx").Op(":=").Qual("context", "TODO").Call(),
		Line(),
		Id("t").Dot("Run").Call(Lit("success"), Func().Params(Id("t").Op("*").Qual("testing", "T")).Block(
			Id(string('t')).Dot("Parallel").Call(),
			Line(),
			Id("adapter").Op(",").Id("err").Op(":=").Id("Init").Call(
				Id("ctx"),
				Map(Qual("github.com/ovotech/go-sync", "ConfigKey")).String().Op("{}"),
			),
			Line(),
			Qual("github.com/stretchr/testify/assert", "NoError").Call(Id("t"), Id("err")),
			Qual("github.com/stretchr/testify/assert", "IsType").Call(Id("t"), Op("&").Id(adapter).Op("{}"), Id("adapter")),
		)),
	)
	file.Line()
}

var Tests = []Fn{ //nolint:gochecknoglobals
	TestImports,
	TestNew,
	TestGet,
	TestAdd,
	TestRemove,
	TestInit,
}
