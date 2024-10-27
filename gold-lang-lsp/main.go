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

func (tsl *TreeShapeListener) EnterEveryRule(ctx antlr.ParserRuleContext) {
	fmt.Println(ctx.GetText())
}

func main() {
	input := antlr.NewInputStream("class aTestClass (aBaseClass)")
	lexer := parser.NewGoldLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewGoldParser(stream)
	p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	tree := p.GoldProg()
	antlr.ParseTreeWalkerDefault.Walk(NewTreeShapeListener(), tree)
}
