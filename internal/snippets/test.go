package snippets

import (
	"fmt"

	"github.com/dave/jennifer/jen"
)

func getTestParams() *jen.Statement {
	return jen.Id("t").Op("*").Qual("testing", "T")
}

func TestNew(f *jen.File, _ string) {
	f.Func().Id("TestNew").Params(jen.Id("t").Op("*").Qual("testing", "T")).Block(
		jen.Id(string('t')).Dot("Parallel").Call(),
	)
	f.Line()
}

func TestGet(f *jen.File, adapter string) { //nolint:varnamelen
	f.Func().Id(fmt.Sprintf("Test%s_Get", adapter)).Params(getTestParams()).Block(
		jen.Id(string('t')).Dot("Parallel").Call(),
		jen.Line(),
		jen.Id("ctx").Op(":=").Qual("context", "TODO").Call(),
		jen.Line(),
		jen.Id("adapter").Op(":=").Id("New").Call(),
		jen.Id("things").Op(",").Id("err").Op(":=").Id("adapter").Dot("Get").Call(jen.Id("ctx")),
		jen.Line(),
		jen.Qual("github.com/stretchr/testify/assert", "NoError").Call(jen.Id("t"), jen.Id("err")),
		jen.Qual("github.com/stretchr/testify/assert", "ElementsMatch").
			Call(jen.Id("t"), jen.Id("things"), jen.Op("[]").String().Op("{}")),
	)
	f.Line()
}

func TestAdd(f *jen.File, adapter string) { //nolint:varnamelen,dupl
	f.Func().Id(fmt.Sprintf("Test%s_Add", adapter)).Params(getTestParams()).Block(
		jen.Id(string('t')).Dot("Parallel").Call(),
		jen.Line(),
		jen.Id("ctx").Op(":=").Qual("context", "TODO").Call(),
		jen.Line(),
		jen.Id("adapter").Op(":=").Id("New").Call(),
		jen.Id("err").Op(":=").Id("adapter").Dot("Add").Call(jen.Id("ctx"), jen.Op("[]").String().Op("{\"foo\"}")),
		jen.Line(),
		jen.Qual("github.com/stretchr/testify/assert", "NoError").Call(jen.Id("t"), jen.Id("err")),
	)
	f.Line()
}

func TestRemove(f *jen.File, adapter string) { //nolint:varnamelen,dupl
	f.Func().Id(fmt.Sprintf("Test%s_Remove", adapter)).Params(getTestParams()).Block(
		jen.Id(string('t')).Dot("Parallel").Call(),
		jen.Line(),
		jen.Id("ctx").Op(":=").Qual("context", "TODO").Call(),
		jen.Line(),
		jen.Id("adapter").Op(":=").Id("New").Call(),
		jen.Id("err").Op(":=").Id("adapter").Dot("Remove").Call(jen.Id("ctx"), jen.Op("[]").String().Op("{\"bar\"}")),
		jen.Line(),
		jen.Qual("github.com/stretchr/testify/assert", "NoError").Call(jen.Id("t"), jen.Id("err")),
	)
	f.Line()
}

func TestInit(f *jen.File, adapter string) {
	f.Func().Id("TestInit").Params(getTestParams()).Block(
		jen.Id(string('t')).Dot("Parallel").Call(),
		jen.Line(),
		jen.Id("t").Dot("Run").Call(jen.Lit("success"), jen.Func().Params(jen.Id("t").Op("*").Qual("testing", "T")).Block(
			jen.Id(string('t')).Dot("Parallel").Call(),
			jen.Line(),
			jen.Id("adapter").Op(",").Id("err").Op(":=").Id("Init").Call(jen.Map(jen.Qual("github.com/ovotech/go-sync", "ConfigKey")).String().Op("{}")),
			jen.Line(),
			jen.Qual("github.com/stretchr/testify/assert", "NoError").Call(jen.Id("t"), jen.Id("err")),
			jen.Qual("github.com/stretchr/testify/assert", "IsType").Call(jen.Id("t"), jen.Op("&").Id(adapter).Op("{}"), jen.Id("adapter")),
		)),
	)
	f.Line()
}
