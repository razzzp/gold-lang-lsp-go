package main

import (
	"fmt"
	"gold-lang-lsp/parser"

	"github.com/antlr4-go/antlr/v4"
)

type TreeShapeListener struct {
	*parser.BaseGoldListener
}

func NewTreeShapeListener() *TreeShapeListener {
	return new(TreeShapeListener)
}

func (tsl *TreeShapeListener) EnterGoldClassDefinition(ctx *parser.GoldClassDefinitionContext) {
	fmt.Printf("found class def %s\n", ctx.GetName().GetText())
}

func (tsl *TreeShapeListener) EnterGoldProcedureDefinition(ctx *parser.GoldProcedureDefinitionContext) {
	fmt.Printf("found proc def %s\n", ctx.GetName().GetText())
	fmt.Printf("statements: %t\n", ctx.GetStatements())
}

func main() {
	input, _ := antlr.NewFileStream("../test-inputs/aTestClass.god")
	lexer := parser.NewGoldLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewGoldParser(stream)
	p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	tree := p.GoldProg()
	antlr.ParseTreeWalkerDefault.Walk(NewTreeShapeListener(), tree)
}
