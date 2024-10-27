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
		"", "'procedure'", "'proc'", "'('", "','", "')'", "'end'", "'endproc'",
		"'function'", "'func'", "':'", "'class'", "'module'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "", "", "PARAMETER_MODIFIER",
		"WHITESPACE", "INT_LITERAL", "DEICMAL_LITERAL", "ID",
	}
	staticData.RuleNames = []string{
		"goldProg", "goldDefinittion", "goldProcedureDefinition", "goldFunctionDefinition",
		"parameterDefinition", "goldClassDefinition", "goldModuleDefinition",
		"varType",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 17, 69, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 1, 0, 4, 0, 18, 8, 0, 11, 0, 12,
		0, 19, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 3, 1, 28, 8, 1, 1, 2, 1, 2,
		1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 3, 2, 37, 8, 2, 1, 2, 1, 2, 3, 2, 41, 8,
		2, 1, 2, 1, 2, 3, 2, 45, 8, 2, 1, 3, 1, 3, 1, 4, 3, 4, 50, 8, 4, 1, 4,
		1, 4, 1, 4, 3, 4, 55, 8, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 3, 5, 62, 8,
		5, 1, 6, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 0, 0, 8, 0, 2, 4, 6, 8, 10, 12,
		14, 0, 1, 1, 0, 8, 9, 71, 0, 17, 1, 0, 0, 0, 2, 27, 1, 0, 0, 0, 4, 44,
		1, 0, 0, 0, 6, 46, 1, 0, 0, 0, 8, 49, 1, 0, 0, 0, 10, 56, 1, 0, 0, 0, 12,
		63, 1, 0, 0, 0, 14, 66, 1, 0, 0, 0, 16, 18, 3, 2, 1, 0, 17, 16, 1, 0, 0,
		0, 18, 19, 1, 0, 0, 0, 19, 17, 1, 0, 0, 0, 19, 20, 1, 0, 0, 0, 20, 21,
		1, 0, 0, 0, 21, 22, 5, 0, 0, 1, 22, 1, 1, 0, 0, 0, 23, 28, 3, 10, 5, 0,
		24, 28, 3, 12, 6, 0, 25, 28, 3, 6, 3, 0, 26, 28, 3, 4, 2, 0, 27, 23, 1,
		0, 0, 0, 27, 24, 1, 0, 0, 0, 27, 25, 1, 0, 0, 0, 27, 26, 1, 0, 0, 0, 28,
		3, 1, 0, 0, 0, 29, 45, 5, 1, 0, 0, 30, 31, 5, 2, 0, 0, 31, 40, 5, 17, 0,
		0, 32, 33, 5, 3, 0, 0, 33, 36, 3, 8, 4, 0, 34, 35, 5, 4, 0, 0, 35, 37,
		3, 8, 4, 0, 36, 34, 1, 0, 0, 0, 36, 37, 1, 0, 0, 0, 37, 38, 1, 0, 0, 0,
		38, 39, 5, 5, 0, 0, 39, 41, 1, 0, 0, 0, 40, 32, 1, 0, 0, 0, 40, 41, 1,
		0, 0, 0, 41, 42, 1, 0, 0, 0, 42, 45, 5, 6, 0, 0, 43, 45, 5, 7, 0, 0, 44,
		29, 1, 0, 0, 0, 44, 30, 1, 0, 0, 0, 44, 43, 1, 0, 0, 0, 45, 5, 1, 0, 0,
		0, 46, 47, 7, 0, 0, 0, 47, 7, 1, 0, 0, 0, 48, 50, 5, 13, 0, 0, 49, 48,
		1, 0, 0, 0, 49, 50, 1, 0, 0, 0, 50, 51, 1, 0, 0, 0, 51, 54, 5, 17, 0, 0,
		52, 53, 5, 10, 0, 0, 53, 55, 3, 14, 7, 0, 54, 52, 1, 0, 0, 0, 54, 55, 1,
		0, 0, 0, 55, 9, 1, 0, 0, 0, 56, 57, 5, 11, 0, 0, 57, 61, 5, 17, 0, 0, 58,
		59, 5, 3, 0, 0, 59, 60, 5, 17, 0, 0, 60, 62, 5, 5, 0, 0, 61, 58, 1, 0,
		0, 0, 61, 62, 1, 0, 0, 0, 62, 11, 1, 0, 0, 0, 63, 64, 5, 12, 0, 0, 64,
		65, 5, 17, 0, 0, 65, 13, 1, 0, 0, 0, 66, 67, 5, 17, 0, 0, 67, 15, 1, 0,
		0, 0, 8, 19, 27, 36, 40, 44, 49, 54, 61,
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
	GoldParserPARAMETER_MODIFIER = 13
	GoldParserWHITESPACE         = 14
	GoldParserINT_LITERAL        = 15
	GoldParserDEICMAL_LITERAL    = 16
	GoldParserID                 = 17
)

// GoldParser rules.
const (
	GoldParserRULE_goldProg                = 0
	GoldParserRULE_goldDefinittion         = 1
	GoldParserRULE_goldProcedureDefinition = 2
	GoldParserRULE_goldFunctionDefinition  = 3
	GoldParserRULE_parameterDefinition     = 4
	GoldParserRULE_goldClassDefinition     = 5
	GoldParserRULE_goldModuleDefinition    = 6
	GoldParserRULE_varType                 = 7
)

// IGoldProgContext is an interface to support dynamic dispatch.
type IGoldProgContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	EOF() antlr.TerminalNode
	AllGoldDefinittion() []IGoldDefinittionContext
	GoldDefinittion(i int) IGoldDefinittionContext

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

func (s *GoldProgContext) AllGoldDefinittion() []IGoldDefinittionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IGoldDefinittionContext); ok {
			len++
		}
	}

	tst := make([]IGoldDefinittionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IGoldDefinittionContext); ok {
			tst[i] = t.(IGoldDefinittionContext)
			i++
		}
	}

	return tst
}

func (s *GoldProgContext) GoldDefinittion(i int) IGoldDefinittionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IGoldDefinittionContext); ok {
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

	return t.(IGoldDefinittionContext)
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
	p.SetState(17)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&7046) != 0) {
		{
			p.SetState(16)
			p.GoldDefinittion()
		}

		p.SetState(19)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(21)
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

// IGoldDefinittionContext is an interface to support dynamic dispatch.
type IGoldDefinittionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	GoldClassDefinition() IGoldClassDefinitionContext
	GoldModuleDefinition() IGoldModuleDefinitionContext
	GoldFunctionDefinition() IGoldFunctionDefinitionContext
	GoldProcedureDefinition() IGoldProcedureDefinitionContext

	// IsGoldDefinittionContext differentiates from other interfaces.
	IsGoldDefinittionContext()
}

type GoldDefinittionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyGoldDefinittionContext() *GoldDefinittionContext {
	var p = new(GoldDefinittionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoldParserRULE_goldDefinittion
	return p
}

func InitEmptyGoldDefinittionContext(p *GoldDefinittionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoldParserRULE_goldDefinittion
}

func (*GoldDefinittionContext) IsGoldDefinittionContext() {}

func NewGoldDefinittionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *GoldDefinittionContext {
	var p = new(GoldDefinittionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoldParserRULE_goldDefinittion

	return p
}

func (s *GoldDefinittionContext) GetParser() antlr.Parser { return s.parser }

func (s *GoldDefinittionContext) GoldClassDefinition() IGoldClassDefinitionContext {
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

func (s *GoldDefinittionContext) GoldModuleDefinition() IGoldModuleDefinitionContext {
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

func (s *GoldDefinittionContext) GoldFunctionDefinition() IGoldFunctionDefinitionContext {
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

func (s *GoldDefinittionContext) GoldProcedureDefinition() IGoldProcedureDefinitionContext {
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

func (s *GoldDefinittionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GoldDefinittionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *GoldDefinittionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoldListener); ok {
		listenerT.EnterGoldDefinittion(s)
	}
}

func (s *GoldDefinittionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoldListener); ok {
		listenerT.ExitGoldDefinittion(s)
	}
}

func (p *GoldParser) GoldDefinittion() (localctx IGoldDefinittionContext) {
	localctx = NewGoldDefinittionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, GoldParserRULE_goldDefinittion)
	p.SetState(27)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case GoldParserT__10:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(23)
			p.GoldClassDefinition()
		}

	case GoldParserT__11:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(24)
			p.GoldModuleDefinition()
		}

	case GoldParserT__7, GoldParserT__8:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(25)
			p.GoldFunctionDefinition()
		}

	case GoldParserT__0, GoldParserT__1, GoldParserT__6:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(26)
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

	// Set_parameterDefinition sets the _parameterDefinition rule contexts.
	Set_parameterDefinition(IParameterDefinitionContext)

	// GetParams returns the params rule context list.
	GetParams() []IParameterDefinitionContext

	// SetParams sets the params rule context list.
	SetParams([]IParameterDefinitionContext)

	// Getter signatures
	ID() antlr.TerminalNode
	AllParameterDefinition() []IParameterDefinitionContext
	ParameterDefinition(i int) IParameterDefinitionContext

	// IsGoldProcedureDefinitionContext differentiates from other interfaces.
	IsGoldProcedureDefinitionContext()
}

type GoldProcedureDefinitionContext struct {
	antlr.BaseParserRuleContext
	parser               antlr.Parser
	name                 antlr.Token
	_parameterDefinition IParameterDefinitionContext
	params               []IParameterDefinitionContext
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

func (s *GoldProcedureDefinitionContext) Set_parameterDefinition(v IParameterDefinitionContext) {
	s._parameterDefinition = v
}

func (s *GoldProcedureDefinitionContext) GetParams() []IParameterDefinitionContext { return s.params }

func (s *GoldProcedureDefinitionContext) SetParams(v []IParameterDefinitionContext) { s.params = v }

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
	p.EnterRule(localctx, 4, GoldParserRULE_goldProcedureDefinition)
	var _la int

	p.SetState(44)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case GoldParserT__0:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(29)
			p.Match(GoldParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case GoldParserT__1:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(30)
			p.Match(GoldParserT__1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(31)

			var _m = p.Match(GoldParserID)

			localctx.(*GoldProcedureDefinitionContext).name = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(40)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == GoldParserT__2 {
			{
				p.SetState(32)
				p.Match(GoldParserT__2)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(33)

				var _x = p.ParameterDefinition()

				localctx.(*GoldProcedureDefinitionContext)._parameterDefinition = _x
			}
			localctx.(*GoldProcedureDefinitionContext).params = append(localctx.(*GoldProcedureDefinitionContext).params, localctx.(*GoldProcedureDefinitionContext)._parameterDefinition)
			p.SetState(36)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)

			if _la == GoldParserT__3 {
				{
					p.SetState(34)
					p.Match(GoldParserT__3)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(35)

					var _x = p.ParameterDefinition()

					localctx.(*GoldProcedureDefinitionContext)._parameterDefinition = _x
				}
				localctx.(*GoldProcedureDefinitionContext).params = append(localctx.(*GoldProcedureDefinitionContext).params, localctx.(*GoldProcedureDefinitionContext)._parameterDefinition)

			}
			{
				p.SetState(38)
				p.Match(GoldParserT__4)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(42)
			p.Match(GoldParserT__5)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case GoldParserT__6:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(43)
			p.Match(GoldParserT__6)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
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
	p.EnterRule(localctx, 6, GoldParserRULE_goldFunctionDefinition)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(46)
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
	p.EnterRule(localctx, 8, GoldParserRULE_parameterDefinition)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(49)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == GoldParserPARAMETER_MODIFIER {
		{
			p.SetState(48)
			p.Match(GoldParserPARAMETER_MODIFIER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(51)

		var _m = p.Match(GoldParserID)

		localctx.(*ParameterDefinitionContext).name = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(54)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == GoldParserT__9 {
		{
			p.SetState(52)
			p.Match(GoldParserT__9)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(53)
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
	p.EnterRule(localctx, 10, GoldParserRULE_goldClassDefinition)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(56)
		p.Match(GoldParserT__10)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(57)

		var _m = p.Match(GoldParserID)

		localctx.(*GoldClassDefinitionContext).name = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(61)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == GoldParserT__2 {
		{
			p.SetState(58)
			p.Match(GoldParserT__2)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(59)

			var _m = p.Match(GoldParserID)

			localctx.(*GoldClassDefinitionContext).parentClass = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(60)
			p.Match(GoldParserT__4)
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
	p.EnterRule(localctx, 12, GoldParserRULE_goldModuleDefinition)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(63)
		p.Match(GoldParserT__11)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(64)

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
	p.EnterRule(localctx, 14, GoldParserRULE_varType)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(66)
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
