// Code generated from antlr4/Gold.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // Gold
import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr4-go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type GoldParser struct {
	*antlr.BaseParser
}

var GoldParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func goldParserInit() {
	staticData := &GoldParserStaticData
	staticData.LiteralNames = []string{
		"", "'class'", "'('", "')'", "'module'", "'procedure'", "'proc'", "','",
		"'end'", "'endproc'", "'function'", "'func'", "':'", "'var'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "PARAMETER_MODIFIER",
		"WHITESPACE", "INT_LITERAL", "DEICMAL_LITERAL", "STRING_LITERAL", "ID",
	}
	staticData.RuleNames = []string{
		"goldProg", "goldDefinition", "goldClassDefinition", "goldModuleDefinition",
		"goldProcedureDefinition", "goldFunctionDefinition", "parameterDefinition",
		"statement", "variableDefinition", "varType",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 19, 83, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 1, 0, 4,
		0, 22, 8, 0, 11, 0, 12, 0, 23, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 3, 1,
		32, 8, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 3, 2, 39, 8, 2, 1, 3, 1, 3, 1,
		3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 3, 4, 50, 8, 4, 1, 4, 1, 4, 3, 4,
		54, 8, 4, 1, 4, 5, 4, 57, 8, 4, 10, 4, 12, 4, 60, 9, 4, 1, 4, 1, 4, 1,
		5, 1, 5, 1, 6, 3, 6, 67, 8, 6, 1, 6, 1, 6, 1, 6, 3, 6, 72, 8, 6, 1, 7,
		1, 7, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 9, 1, 9, 1, 9, 0, 0, 10, 0, 2, 4,
		6, 8, 10, 12, 14, 16, 18, 0, 3, 1, 0, 5, 6, 1, 0, 8, 9, 1, 0, 10, 11, 82,
		0, 21, 1, 0, 0, 0, 2, 31, 1, 0, 0, 0, 4, 33, 1, 0, 0, 0, 6, 40, 1, 0, 0,
		0, 8, 43, 1, 0, 0, 0, 10, 63, 1, 0, 0, 0, 12, 66, 1, 0, 0, 0, 14, 73, 1,
		0, 0, 0, 16, 75, 1, 0, 0, 0, 18, 80, 1, 0, 0, 0, 20, 22, 3, 2, 1, 0, 21,
		20, 1, 0, 0, 0, 22, 23, 1, 0, 0, 0, 23, 21, 1, 0, 0, 0, 23, 24, 1, 0, 0,
		0, 24, 25, 1, 0, 0, 0, 25, 26, 5, 0, 0, 1, 26, 1, 1, 0, 0, 0, 27, 32, 3,
		4, 2, 0, 28, 32, 3, 6, 3, 0, 29, 32, 3, 10, 5, 0, 30, 32, 3, 8, 4, 0, 31,
		27, 1, 0, 0, 0, 31, 28, 1, 0, 0, 0, 31, 29, 1, 0, 0, 0, 31, 30, 1, 0, 0,
		0, 32, 3, 1, 0, 0, 0, 33, 34, 5, 1, 0, 0, 34, 38, 5, 19, 0, 0, 35, 36,
		5, 2, 0, 0, 36, 37, 5, 19, 0, 0, 37, 39, 5, 3, 0, 0, 38, 35, 1, 0, 0, 0,
		38, 39, 1, 0, 0, 0, 39, 5, 1, 0, 0, 0, 40, 41, 5, 4, 0, 0, 41, 42, 5, 19,
		0, 0, 42, 7, 1, 0, 0, 0, 43, 44, 7, 0, 0, 0, 44, 53, 5, 19, 0, 0, 45, 46,
		5, 2, 0, 0, 46, 49, 3, 12, 6, 0, 47, 48, 5, 7, 0, 0, 48, 50, 3, 12, 6,
		0, 49, 47, 1, 0, 0, 0, 49, 50, 1, 0, 0, 0, 50, 51, 1, 0, 0, 0, 51, 52,
		5, 3, 0, 0, 52, 54, 1, 0, 0, 0, 53, 45, 1, 0, 0, 0, 53, 54, 1, 0, 0, 0,
		54, 58, 1, 0, 0, 0, 55, 57, 3, 14, 7, 0, 56, 55, 1, 0, 0, 0, 57, 60, 1,
		0, 0, 0, 58, 56, 1, 0, 0, 0, 58, 59, 1, 0, 0, 0, 59, 61, 1, 0, 0, 0, 60,
		58, 1, 0, 0, 0, 61, 62, 7, 1, 0, 0, 62, 9, 1, 0, 0, 0, 63, 64, 7, 2, 0,
		0, 64, 11, 1, 0, 0, 0, 65, 67, 5, 14, 0, 0, 66, 65, 1, 0, 0, 0, 66, 67,
		1, 0, 0, 0, 67, 68, 1, 0, 0, 0, 68, 71, 5, 19, 0, 0, 69, 70, 5, 12, 0,
		0, 70, 72, 3, 18, 9, 0, 71, 69, 1, 0, 0, 0, 71, 72, 1, 0, 0, 0, 72, 13,
		1, 0, 0, 0, 73, 74, 3, 16, 8, 0, 74, 15, 1, 0, 0, 0, 75, 76, 5, 13, 0,
		0, 76, 77, 5, 19, 0, 0, 77, 78, 5, 12, 0, 0, 78, 79, 3, 18, 9, 0, 79, 17,
		1, 0, 0, 0, 80, 81, 5, 19, 0, 0, 81, 19, 1, 0, 0, 0, 8, 23, 31, 38, 49,
		53, 58, 66, 71,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// GoldParserInit initializes any static state used to implement GoldParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewGoldParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func GoldParserInit() {
	staticData := &GoldParserStaticData
	staticData.once.Do(goldParserInit)
}

// NewGoldParser produces a new parser instance for the optional input antlr.TokenStream.
func NewGoldParser(input antlr.TokenStream) *GoldParser {
	GoldParserInit()
	this := new(GoldParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &GoldParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "Gold.g4"

	return this
}

// GoldParser tokens.
const (
	GoldParserEOF                = antlr.TokenEOF
	GoldParserT__0               = 1
	GoldParserT__1               = 2
	GoldParserT__2               = 3
	GoldParserT__3               = 4
	GoldParserT__4               = 5
	GoldParserT__5               = 6
	GoldParserT__6               = 7
	GoldParserT__7               = 8
	GoldParserT__8               = 9
	GoldParserT__9               = 10
	GoldParserT__10              = 11
	GoldParserT__11              = 12
	GoldParserT__12              = 13
	GoldParserPARAMETER_MODIFIER = 14
	GoldParserWHITESPACE         = 15
	GoldParserINT_LITERAL        = 16
	GoldParserDEICMAL_LITERAL    = 17
	GoldParserSTRING_LITERAL     = 18
	GoldParserID                 = 19
)

// GoldParser rules.
const (
	GoldParserRULE_goldProg                = 0
	GoldParserRULE_goldDefinition          = 1
	GoldParserRULE_goldClassDefinition     = 2
	GoldParserRULE_goldModuleDefinition    = 3
	GoldParserRULE_goldProcedureDefinition = 4
	GoldParserRULE_goldFunctionDefinition  = 5
	GoldParserRULE_parameterDefinition     = 6
	GoldParserRULE_statement               = 7
	GoldParserRULE_variableDefinition      = 8
	GoldParserRULE_varType                 = 9
)

// IGoldProgContext is an interface to support dynamic dispatch.
type IGoldProgContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	EOF() antlr.TerminalNode
	AllGoldDefinition() []IGoldDefinitionContext
	GoldDefinition(i int) IGoldDefinitionContext

	// IsGoldProgContext differentiates from other interfaces.
	IsGoldProgContext()
}

type GoldProgContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyGoldProgContext() *GoldProgContext {
	var p = new(GoldProgContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoldParserRULE_goldProg
	return p
}

func InitEmptyGoldProgContext(p *GoldProgContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoldParserRULE_goldProg
}

func (*GoldProgContext) IsGoldProgContext() {}

func NewGoldProgContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *GoldProgContext {
	var p = new(GoldProgContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoldParserRULE_goldProg

	return p
}

func (s *GoldProgContext) GetParser() antlr.Parser { return s.parser }

func (s *GoldProgContext) EOF() antlr.TerminalNode {
	return s.GetToken(GoldParserEOF, 0)
}

func (s *GoldProgContext) AllGoldDefinition() []IGoldDefinitionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IGoldDefinitionContext); ok {
			len++
		}
	}

	tst := make([]IGoldDefinitionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IGoldDefinitionContext); ok {
			tst[i] = t.(IGoldDefinitionContext)
			i++
		}
	}

	return tst
}

func (s *GoldProgContext) GoldDefinition(i int) IGoldDefinitionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IGoldDefinitionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IGoldDefinitionContext)
}

func (s *GoldProgContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GoldProgContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *GoldProgContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoldListener); ok {
		listenerT.EnterGoldProg(s)
	}
}

func (s *GoldProgContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoldListener); ok {
		listenerT.ExitGoldProg(s)
	}
}

func (p *GoldParser) GoldProg() (localctx IGoldProgContext) {
	localctx = NewGoldProgContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, GoldParserRULE_goldProg)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(21)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&3186) != 0) {
		{
			p.SetState(20)
			p.GoldDefinition()
		}

		p.SetState(23)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(25)
		p.Match(GoldParserEOF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IGoldDefinitionContext is an interface to support dynamic dispatch.
type IGoldDefinitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	GoldClassDefinition() IGoldClassDefinitionContext
	GoldModuleDefinition() IGoldModuleDefinitionContext
	GoldFunctionDefinition() IGoldFunctionDefinitionContext
	GoldProcedureDefinition() IGoldProcedureDefinitionContext

	// IsGoldDefinitionContext differentiates from other interfaces.
	IsGoldDefinitionContext()
}

type GoldDefinitionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyGoldDefinitionContext() *GoldDefinitionContext {
	var p = new(GoldDefinitionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoldParserRULE_goldDefinition
	return p
}

func InitEmptyGoldDefinitionContext(p *GoldDefinitionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoldParserRULE_goldDefinition
}

func (*GoldDefinitionContext) IsGoldDefinitionContext() {}

func NewGoldDefinitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *GoldDefinitionContext {
	var p = new(GoldDefinitionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoldParserRULE_goldDefinition

	return p
}

func (s *GoldDefinitionContext) GetParser() antlr.Parser { return s.parser }

func (s *GoldDefinitionContext) GoldClassDefinition() IGoldClassDefinitionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IGoldClassDefinitionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IGoldClassDefinitionContext)
}

func (s *GoldDefinitionContext) GoldModuleDefinition() IGoldModuleDefinitionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IGoldModuleDefinitionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IGoldModuleDefinitionContext)
}

func (s *GoldDefinitionContext) GoldFunctionDefinition() IGoldFunctionDefinitionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IGoldFunctionDefinitionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IGoldFunctionDefinitionContext)
}

func (s *GoldDefinitionContext) GoldProcedureDefinition() IGoldProcedureDefinitionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IGoldProcedureDefinitionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IGoldProcedureDefinitionContext)
}

func (s *GoldDefinitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GoldDefinitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *GoldDefinitionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoldListener); ok {
		listenerT.EnterGoldDefinition(s)
	}
}

func (s *GoldDefinitionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoldListener); ok {
		listenerT.ExitGoldDefinition(s)
	}
}

func (p *GoldParser) GoldDefinition() (localctx IGoldDefinitionContext) {
	localctx = NewGoldDefinitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, GoldParserRULE_goldDefinition)
	p.SetState(31)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case GoldParserT__0:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(27)
			p.GoldClassDefinition()
		}

	case GoldParserT__3:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(28)
			p.GoldModuleDefinition()
		}

	case GoldParserT__9, GoldParserT__10:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(29)
			p.GoldFunctionDefinition()
		}

	case GoldParserT__4, GoldParserT__5:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(30)
			p.GoldProcedureDefinition()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IGoldClassDefinitionContext is an interface to support dynamic dispatch.
type IGoldClassDefinitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetName returns the name token.
	GetName() antlr.Token

	// GetParentClass returns the parentClass token.
	GetParentClass() antlr.Token

	// SetName sets the name token.
	SetName(antlr.Token)

	// SetParentClass sets the parentClass token.
	SetParentClass(antlr.Token)

	// Getter signatures
	AllID() []antlr.TerminalNode
	ID(i int) antlr.TerminalNode

	// IsGoldClassDefinitionContext differentiates from other interfaces.
	IsGoldClassDefinitionContext()
}

type GoldClassDefinitionContext struct {
	antlr.BaseParserRuleContext
	parser      antlr.Parser
	name        antlr.Token
	parentClass antlr.Token
}

func NewEmptyGoldClassDefinitionContext() *GoldClassDefinitionContext {
	var p = new(GoldClassDefinitionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoldParserRULE_goldClassDefinition
	return p
}

func InitEmptyGoldClassDefinitionContext(p *GoldClassDefinitionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoldParserRULE_goldClassDefinition
}

func (*GoldClassDefinitionContext) IsGoldClassDefinitionContext() {}

func NewGoldClassDefinitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *GoldClassDefinitionContext {
	var p = new(GoldClassDefinitionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoldParserRULE_goldClassDefinition

	return p
}

func (s *GoldClassDefinitionContext) GetParser() antlr.Parser { return s.parser }

func (s *GoldClassDefinitionContext) GetName() antlr.Token { return s.name }

func (s *GoldClassDefinitionContext) GetParentClass() antlr.Token { return s.parentClass }

func (s *GoldClassDefinitionContext) SetName(v antlr.Token) { s.name = v }

func (s *GoldClassDefinitionContext) SetParentClass(v antlr.Token) { s.parentClass = v }

func (s *GoldClassDefinitionContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(GoldParserID)
}

func (s *GoldClassDefinitionContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(GoldParserID, i)
}

func (s *GoldClassDefinitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GoldClassDefinitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *GoldClassDefinitionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoldListener); ok {
		listenerT.EnterGoldClassDefinition(s)
	}
}

func (s *GoldClassDefinitionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoldListener); ok {
		listenerT.ExitGoldClassDefinition(s)
	}
}

func (p *GoldParser) GoldClassDefinition() (localctx IGoldClassDefinitionContext) {
	localctx = NewGoldClassDefinitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, GoldParserRULE_goldClassDefinition)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(33)
		p.Match(GoldParserT__0)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(34)

		var _m = p.Match(GoldParserID)

		localctx.(*GoldClassDefinitionContext).name = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(38)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == GoldParserT__1 {
		{
			p.SetState(35)
			p.Match(GoldParserT__1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(36)

			var _m = p.Match(GoldParserID)

			localctx.(*GoldClassDefinitionContext).parentClass = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(37)
			p.Match(GoldParserT__2)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IGoldModuleDefinitionContext is an interface to support dynamic dispatch.
type IGoldModuleDefinitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetName returns the name token.
	GetName() antlr.Token

	// SetName sets the name token.
	SetName(antlr.Token)

	// Getter signatures
	ID() antlr.TerminalNode

	// IsGoldModuleDefinitionContext differentiates from other interfaces.
	IsGoldModuleDefinitionContext()
}

type GoldModuleDefinitionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	name   antlr.Token
}

func NewEmptyGoldModuleDefinitionContext() *GoldModuleDefinitionContext {
	var p = new(GoldModuleDefinitionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoldParserRULE_goldModuleDefinition
	return p
}

func InitEmptyGoldModuleDefinitionContext(p *GoldModuleDefinitionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoldParserRULE_goldModuleDefinition
}

func (*GoldModuleDefinitionContext) IsGoldModuleDefinitionContext() {}

func NewGoldModuleDefinitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *GoldModuleDefinitionContext {
	var p = new(GoldModuleDefinitionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoldParserRULE_goldModuleDefinition

	return p
}

func (s *GoldModuleDefinitionContext) GetParser() antlr.Parser { return s.parser }

func (s *GoldModuleDefinitionContext) GetName() antlr.Token { return s.name }

func (s *GoldModuleDefinitionContext) SetName(v antlr.Token) { s.name = v }

func (s *GoldModuleDefinitionContext) ID() antlr.TerminalNode {
	return s.GetToken(GoldParserID, 0)
}

func (s *GoldModuleDefinitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GoldModuleDefinitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *GoldModuleDefinitionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoldListener); ok {
		listenerT.EnterGoldModuleDefinition(s)
	}
}

func (s *GoldModuleDefinitionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoldListener); ok {
		listenerT.ExitGoldModuleDefinition(s)
	}
}

func (p *GoldParser) GoldModuleDefinition() (localctx IGoldModuleDefinitionContext) {
	localctx = NewGoldModuleDefinitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, GoldParserRULE_goldModuleDefinition)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(40)
		p.Match(GoldParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(41)

		var _m = p.Match(GoldParserID)

		localctx.(*GoldModuleDefinitionContext).name = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IGoldProcedureDefinitionContext is an interface to support dynamic dispatch.
type IGoldProcedureDefinitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetName returns the name token.
	GetName() antlr.Token

	// SetName sets the name token.
	SetName(antlr.Token)

	// Get_parameterDefinition returns the _parameterDefinition rule contexts.
	Get_parameterDefinition() IParameterDefinitionContext

	// Get_statement returns the _statement rule contexts.
	Get_statement() IStatementContext

	// Set_parameterDefinition sets the _parameterDefinition rule contexts.
	Set_parameterDefinition(IParameterDefinitionContext)

	// Set_statement sets the _statement rule contexts.
	Set_statement(IStatementContext)

	// GetParams returns the params rule context list.
	GetParams() []IParameterDefinitionContext

	// GetStatements returns the statements rule context list.
	GetStatements() []IStatementContext

	// SetParams sets the params rule context list.
	SetParams([]IParameterDefinitionContext)

	// SetStatements sets the statements rule context list.
	SetStatements([]IStatementContext)

	// Getter signatures
	ID() antlr.TerminalNode
	AllParameterDefinition() []IParameterDefinitionContext
	ParameterDefinition(i int) IParameterDefinitionContext
	AllStatement() []IStatementContext
	Statement(i int) IStatementContext

	// IsGoldProcedureDefinitionContext differentiates from other interfaces.
	IsGoldProcedureDefinitionContext()
}

type GoldProcedureDefinitionContext struct {
	antlr.BaseParserRuleContext
	parser               antlr.Parser
	name                 antlr.Token
	_parameterDefinition IParameterDefinitionContext
	params               []IParameterDefinitionContext
	_statement           IStatementContext
	statements           []IStatementContext
}

func NewEmptyGoldProcedureDefinitionContext() *GoldProcedureDefinitionContext {
	var p = new(GoldProcedureDefinitionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoldParserRULE_goldProcedureDefinition
	return p
}

func InitEmptyGoldProcedureDefinitionContext(p *GoldProcedureDefinitionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoldParserRULE_goldProcedureDefinition
}

func (*GoldProcedureDefinitionContext) IsGoldProcedureDefinitionContext() {}

func NewGoldProcedureDefinitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *GoldProcedureDefinitionContext {
	var p = new(GoldProcedureDefinitionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoldParserRULE_goldProcedureDefinition

	return p
}

func (s *GoldProcedureDefinitionContext) GetParser() antlr.Parser { return s.parser }

func (s *GoldProcedureDefinitionContext) GetName() antlr.Token { return s.name }

func (s *GoldProcedureDefinitionContext) SetName(v antlr.Token) { s.name = v }

func (s *GoldProcedureDefinitionContext) Get_parameterDefinition() IParameterDefinitionContext {
	return s._parameterDefinition
}

func (s *GoldProcedureDefinitionContext) Get_statement() IStatementContext { return s._statement }

func (s *GoldProcedureDefinitionContext) Set_parameterDefinition(v IParameterDefinitionContext) {
	s._parameterDefinition = v
}

func (s *GoldProcedureDefinitionContext) Set_statement(v IStatementContext) { s._statement = v }

func (s *GoldProcedureDefinitionContext) GetParams() []IParameterDefinitionContext { return s.params }

func (s *GoldProcedureDefinitionContext) GetStatements() []IStatementContext { return s.statements }

func (s *GoldProcedureDefinitionContext) SetParams(v []IParameterDefinitionContext) { s.params = v }

func (s *GoldProcedureDefinitionContext) SetStatements(v []IStatementContext) { s.statements = v }

func (s *GoldProcedureDefinitionContext) ID() antlr.TerminalNode {
	return s.GetToken(GoldParserID, 0)
}

func (s *GoldProcedureDefinitionContext) AllParameterDefinition() []IParameterDefinitionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IParameterDefinitionContext); ok {
			len++
		}
	}

	tst := make([]IParameterDefinitionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IParameterDefinitionContext); ok {
			tst[i] = t.(IParameterDefinitionContext)
			i++
		}
	}

	return tst
}

func (s *GoldProcedureDefinitionContext) ParameterDefinition(i int) IParameterDefinitionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParameterDefinitionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParameterDefinitionContext)
}

func (s *GoldProcedureDefinitionContext) AllStatement() []IStatementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStatementContext); ok {
			len++
		}
	}

	tst := make([]IStatementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStatementContext); ok {
			tst[i] = t.(IStatementContext)
			i++
		}
	}

	return tst
}

func (s *GoldProcedureDefinitionContext) Statement(i int) IStatementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *GoldProcedureDefinitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GoldProcedureDefinitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *GoldProcedureDefinitionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoldListener); ok {
		listenerT.EnterGoldProcedureDefinition(s)
	}
}

func (s *GoldProcedureDefinitionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoldListener); ok {
		listenerT.ExitGoldProcedureDefinition(s)
	}
}

func (p *GoldParser) GoldProcedureDefinition() (localctx IGoldProcedureDefinitionContext) {
	localctx = NewGoldProcedureDefinitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, GoldParserRULE_goldProcedureDefinition)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(43)
		_la = p.GetTokenStream().LA(1)

		if !(_la == GoldParserT__4 || _la == GoldParserT__5) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(44)

		var _m = p.Match(GoldParserID)

		localctx.(*GoldProcedureDefinitionContext).name = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(53)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == GoldParserT__1 {
		{
			p.SetState(45)
			p.Match(GoldParserT__1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(46)

			var _x = p.ParameterDefinition()

			localctx.(*GoldProcedureDefinitionContext)._parameterDefinition = _x
		}
		localctx.(*GoldProcedureDefinitionContext).params = append(localctx.(*GoldProcedureDefinitionContext).params, localctx.(*GoldProcedureDefinitionContext)._parameterDefinition)
		p.SetState(49)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == GoldParserT__6 {
			{
				p.SetState(47)
				p.Match(GoldParserT__6)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(48)

				var _x = p.ParameterDefinition()

				localctx.(*GoldProcedureDefinitionContext)._parameterDefinition = _x
			}
			localctx.(*GoldProcedureDefinitionContext).params = append(localctx.(*GoldProcedureDefinitionContext).params, localctx.(*GoldProcedureDefinitionContext)._parameterDefinition)

		}
		{
			p.SetState(51)
			p.Match(GoldParserT__2)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	p.SetState(58)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == GoldParserT__12 {
		{
			p.SetState(55)

			var _x = p.Statement()

			localctx.(*GoldProcedureDefinitionContext)._statement = _x
		}
		localctx.(*GoldProcedureDefinitionContext).statements = append(localctx.(*GoldProcedureDefinitionContext).statements, localctx.(*GoldProcedureDefinitionContext)._statement)

		p.SetState(60)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(61)
		_la = p.GetTokenStream().LA(1)

		if !(_la == GoldParserT__7 || _la == GoldParserT__8) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IGoldFunctionDefinitionContext is an interface to support dynamic dispatch.
type IGoldFunctionDefinitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsGoldFunctionDefinitionContext differentiates from other interfaces.
	IsGoldFunctionDefinitionContext()
}

type GoldFunctionDefinitionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyGoldFunctionDefinitionContext() *GoldFunctionDefinitionContext {
	var p = new(GoldFunctionDefinitionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoldParserRULE_goldFunctionDefinition
	return p
}

func InitEmptyGoldFunctionDefinitionContext(p *GoldFunctionDefinitionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoldParserRULE_goldFunctionDefinition
}

func (*GoldFunctionDefinitionContext) IsGoldFunctionDefinitionContext() {}

func NewGoldFunctionDefinitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *GoldFunctionDefinitionContext {
	var p = new(GoldFunctionDefinitionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoldParserRULE_goldFunctionDefinition

	return p
}

func (s *GoldFunctionDefinitionContext) GetParser() antlr.Parser { return s.parser }
func (s *GoldFunctionDefinitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GoldFunctionDefinitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *GoldFunctionDefinitionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoldListener); ok {
		listenerT.EnterGoldFunctionDefinition(s)
	}
}

func (s *GoldFunctionDefinitionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoldListener); ok {
		listenerT.ExitGoldFunctionDefinition(s)
	}
}

func (p *GoldParser) GoldFunctionDefinition() (localctx IGoldFunctionDefinitionContext) {
	localctx = NewGoldFunctionDefinitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, GoldParserRULE_goldFunctionDefinition)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(63)
		_la = p.GetTokenStream().LA(1)

		if !(_la == GoldParserT__9 || _la == GoldParserT__10) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IParameterDefinitionContext is an interface to support dynamic dispatch.
type IParameterDefinitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetName returns the name token.
	GetName() antlr.Token

	// SetName sets the name token.
	SetName(antlr.Token)

	// Getter signatures
	ID() antlr.TerminalNode
	PARAMETER_MODIFIER() antlr.TerminalNode
	VarType() IVarTypeContext

	// IsParameterDefinitionContext differentiates from other interfaces.
	IsParameterDefinitionContext()
}

type ParameterDefinitionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	name   antlr.Token
}

func NewEmptyParameterDefinitionContext() *ParameterDefinitionContext {
	var p = new(ParameterDefinitionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoldParserRULE_parameterDefinition
	return p
}

func InitEmptyParameterDefinitionContext(p *ParameterDefinitionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoldParserRULE_parameterDefinition
}

func (*ParameterDefinitionContext) IsParameterDefinitionContext() {}

func NewParameterDefinitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParameterDefinitionContext {
	var p = new(ParameterDefinitionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoldParserRULE_parameterDefinition

	return p
}

func (s *ParameterDefinitionContext) GetParser() antlr.Parser { return s.parser }

func (s *ParameterDefinitionContext) GetName() antlr.Token { return s.name }

func (s *ParameterDefinitionContext) SetName(v antlr.Token) { s.name = v }

func (s *ParameterDefinitionContext) ID() antlr.TerminalNode {
	return s.GetToken(GoldParserID, 0)
}

func (s *ParameterDefinitionContext) PARAMETER_MODIFIER() antlr.TerminalNode {
	return s.GetToken(GoldParserPARAMETER_MODIFIER, 0)
}

func (s *ParameterDefinitionContext) VarType() IVarTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVarTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVarTypeContext)
}

func (s *ParameterDefinitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParameterDefinitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParameterDefinitionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoldListener); ok {
		listenerT.EnterParameterDefinition(s)
	}
}

func (s *ParameterDefinitionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoldListener); ok {
		listenerT.ExitParameterDefinition(s)
	}
}

func (p *GoldParser) ParameterDefinition() (localctx IParameterDefinitionContext) {
	localctx = NewParameterDefinitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, GoldParserRULE_parameterDefinition)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(66)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == GoldParserPARAMETER_MODIFIER {
		{
			p.SetState(65)
			p.Match(GoldParserPARAMETER_MODIFIER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(68)

		var _m = p.Match(GoldParserID)

		localctx.(*ParameterDefinitionContext).name = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(71)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == GoldParserT__11 {
		{
			p.SetState(69)
			p.Match(GoldParserT__11)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(70)
			p.VarType()
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IStatementContext is an interface to support dynamic dispatch.
type IStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	VariableDefinition() IVariableDefinitionContext

	// IsStatementContext differentiates from other interfaces.
	IsStatementContext()
}

type StatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatementContext() *StatementContext {
	var p = new(StatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoldParserRULE_statement
	return p
}

func InitEmptyStatementContext(p *StatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoldParserRULE_statement
}

func (*StatementContext) IsStatementContext() {}

func NewStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementContext {
	var p = new(StatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoldParserRULE_statement

	return p
}

func (s *StatementContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementContext) VariableDefinition() IVariableDefinitionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVariableDefinitionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVariableDefinitionContext)
}

func (s *StatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoldListener); ok {
		listenerT.EnterStatement(s)
	}
}

func (s *StatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoldListener); ok {
		listenerT.ExitStatement(s)
	}
}

func (p *GoldParser) Statement() (localctx IStatementContext) {
	localctx = NewStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, GoldParserRULE_statement)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(73)
		p.VariableDefinition()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IVariableDefinitionContext is an interface to support dynamic dispatch.
type IVariableDefinitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetName returns the name token.
	GetName() antlr.Token

	// SetName sets the name token.
	SetName(antlr.Token)

	// Getter signatures
	VarType() IVarTypeContext
	ID() antlr.TerminalNode

	// IsVariableDefinitionContext differentiates from other interfaces.
	IsVariableDefinitionContext()
}

type VariableDefinitionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	name   antlr.Token
}

func NewEmptyVariableDefinitionContext() *VariableDefinitionContext {
	var p = new(VariableDefinitionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoldParserRULE_variableDefinition
	return p
}

func InitEmptyVariableDefinitionContext(p *VariableDefinitionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoldParserRULE_variableDefinition
}

func (*VariableDefinitionContext) IsVariableDefinitionContext() {}

func NewVariableDefinitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VariableDefinitionContext {
	var p = new(VariableDefinitionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoldParserRULE_variableDefinition

	return p
}

func (s *VariableDefinitionContext) GetParser() antlr.Parser { return s.parser }

func (s *VariableDefinitionContext) GetName() antlr.Token { return s.name }

func (s *VariableDefinitionContext) SetName(v antlr.Token) { s.name = v }

func (s *VariableDefinitionContext) VarType() IVarTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVarTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVarTypeContext)
}

func (s *VariableDefinitionContext) ID() antlr.TerminalNode {
	return s.GetToken(GoldParserID, 0)
}

func (s *VariableDefinitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VariableDefinitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VariableDefinitionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoldListener); ok {
		listenerT.EnterVariableDefinition(s)
	}
}

func (s *VariableDefinitionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoldListener); ok {
		listenerT.ExitVariableDefinition(s)
	}
}

func (p *GoldParser) VariableDefinition() (localctx IVariableDefinitionContext) {
	localctx = NewVariableDefinitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, GoldParserRULE_variableDefinition)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(75)
		p.Match(GoldParserT__12)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(76)

		var _m = p.Match(GoldParserID)

		localctx.(*VariableDefinitionContext).name = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(77)
		p.Match(GoldParserT__11)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(78)
		p.VarType()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IVarTypeContext is an interface to support dynamic dispatch.
type IVarTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ID() antlr.TerminalNode

	// IsVarTypeContext differentiates from other interfaces.
	IsVarTypeContext()
}

type VarTypeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVarTypeContext() *VarTypeContext {
	var p = new(VarTypeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoldParserRULE_varType
	return p
}

func InitEmptyVarTypeContext(p *VarTypeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoldParserRULE_varType
}

func (*VarTypeContext) IsVarTypeContext() {}

func NewVarTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VarTypeContext {
	var p = new(VarTypeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoldParserRULE_varType

	return p
}

func (s *VarTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *VarTypeContext) ID() antlr.TerminalNode {
	return s.GetToken(GoldParserID, 0)
}

func (s *VarTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VarTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VarTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoldListener); ok {
		listenerT.EnterVarType(s)
	}
}

func (s *VarTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoldListener); ok {
		listenerT.ExitVarType(s)
	}
}

func (p *GoldParser) VarType() (localctx IVarTypeContext) {
	localctx = NewVarTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, GoldParserRULE_varType)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(80)
		p.Match(GoldParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}
