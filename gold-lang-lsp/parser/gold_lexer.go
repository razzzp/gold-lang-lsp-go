// Code generated from antlr4/Gold.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser

import (
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"sync"
	"unicode"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter

type GoldLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var GoldLexerLexerStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	ChannelNames           []string
	ModeNames              []string
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func goldlexerLexerInit() {
	staticData := &GoldLexerLexerStaticData
	staticData.ChannelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.ModeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.LiteralNames = []string{
		"", "'procedure'", "'proc'", "'('", "','", "')'", "'end'", "'endproc'",
		"'function'", "'func'", "':'", "'class'", "'module'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "", "", "PARAMETER_MODIFIER",
		"WHITESPACE", "INT_LITERAL", "DEICMAL_LITERAL", "ID",
	}
	staticData.RuleNames = []string{
		"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "T__8",
		"T__9", "T__10", "T__11", "PARAMETER_MODIFIER", "WHITESPACE", "INT_LITERAL",
		"DEICMAL_LITERAL", "ID",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 17, 148, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1,
		0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 1, 2, 1, 3, 1, 3, 1, 4, 1,
		4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1,
		6, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 8, 1, 8, 1,
		8, 1, 8, 1, 8, 1, 9, 1, 9, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1,
		11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 12, 1, 12, 1, 12, 1, 12,
		1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 3, 12, 111,
		8, 12, 1, 13, 4, 13, 114, 8, 13, 11, 13, 12, 13, 115, 1, 13, 1, 13, 1,
		14, 3, 14, 121, 8, 14, 1, 14, 4, 14, 124, 8, 14, 11, 14, 12, 14, 125, 1,
		15, 3, 15, 129, 8, 15, 1, 15, 4, 15, 132, 8, 15, 11, 15, 12, 15, 133, 1,
		15, 1, 15, 4, 15, 138, 8, 15, 11, 15, 12, 15, 139, 1, 16, 1, 16, 5, 16,
		144, 8, 16, 10, 16, 12, 16, 147, 9, 16, 0, 0, 17, 1, 1, 3, 2, 5, 3, 7,
		4, 9, 5, 11, 6, 13, 7, 15, 8, 17, 9, 19, 10, 21, 11, 23, 12, 25, 13, 27,
		14, 29, 15, 31, 16, 33, 17, 1, 0, 21, 2, 0, 80, 80, 112, 112, 2, 0, 82,
		82, 114, 114, 2, 0, 79, 79, 111, 111, 2, 0, 67, 67, 99, 99, 2, 0, 69, 69,
		101, 101, 2, 0, 68, 68, 100, 100, 2, 0, 85, 85, 117, 117, 2, 0, 78, 78,
		110, 110, 2, 0, 70, 70, 102, 102, 2, 0, 84, 84, 116, 116, 2, 0, 73, 73,
		105, 105, 2, 0, 76, 76, 108, 108, 2, 0, 65, 65, 97, 97, 2, 0, 83, 83, 115,
		115, 2, 0, 77, 77, 109, 109, 2, 0, 86, 86, 118, 118, 3, 0, 9, 10, 13, 13,
		32, 32, 2, 0, 43, 43, 45, 45, 1, 0, 48, 57, 3, 0, 65, 90, 95, 95, 97, 122,
		4, 0, 48, 57, 65, 90, 95, 95, 97, 122, 156, 0, 1, 1, 0, 0, 0, 0, 3, 1,
		0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1,
		0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19,
		1, 0, 0, 0, 0, 21, 1, 0, 0, 0, 0, 23, 1, 0, 0, 0, 0, 25, 1, 0, 0, 0, 0,
		27, 1, 0, 0, 0, 0, 29, 1, 0, 0, 0, 0, 31, 1, 0, 0, 0, 0, 33, 1, 0, 0, 0,
		1, 35, 1, 0, 0, 0, 3, 45, 1, 0, 0, 0, 5, 50, 1, 0, 0, 0, 7, 52, 1, 0, 0,
		0, 9, 54, 1, 0, 0, 0, 11, 56, 1, 0, 0, 0, 13, 60, 1, 0, 0, 0, 15, 68, 1,
		0, 0, 0, 17, 77, 1, 0, 0, 0, 19, 82, 1, 0, 0, 0, 21, 84, 1, 0, 0, 0, 23,
		90, 1, 0, 0, 0, 25, 110, 1, 0, 0, 0, 27, 113, 1, 0, 0, 0, 29, 120, 1, 0,
		0, 0, 31, 128, 1, 0, 0, 0, 33, 141, 1, 0, 0, 0, 35, 36, 7, 0, 0, 0, 36,
		37, 7, 1, 0, 0, 37, 38, 7, 2, 0, 0, 38, 39, 7, 3, 0, 0, 39, 40, 7, 4, 0,
		0, 40, 41, 7, 5, 0, 0, 41, 42, 7, 6, 0, 0, 42, 43, 7, 1, 0, 0, 43, 44,
		7, 4, 0, 0, 44, 2, 1, 0, 0, 0, 45, 46, 7, 0, 0, 0, 46, 47, 7, 1, 0, 0,
		47, 48, 7, 2, 0, 0, 48, 49, 7, 3, 0, 0, 49, 4, 1, 0, 0, 0, 50, 51, 5, 40,
		0, 0, 51, 6, 1, 0, 0, 0, 52, 53, 5, 44, 0, 0, 53, 8, 1, 0, 0, 0, 54, 55,
		5, 41, 0, 0, 55, 10, 1, 0, 0, 0, 56, 57, 7, 4, 0, 0, 57, 58, 7, 7, 0, 0,
		58, 59, 7, 5, 0, 0, 59, 12, 1, 0, 0, 0, 60, 61, 7, 4, 0, 0, 61, 62, 7,
		7, 0, 0, 62, 63, 7, 5, 0, 0, 63, 64, 7, 0, 0, 0, 64, 65, 7, 1, 0, 0, 65,
		66, 7, 2, 0, 0, 66, 67, 7, 3, 0, 0, 67, 14, 1, 0, 0, 0, 68, 69, 7, 8, 0,
		0, 69, 70, 7, 6, 0, 0, 70, 71, 7, 7, 0, 0, 71, 72, 7, 3, 0, 0, 72, 73,
		7, 9, 0, 0, 73, 74, 7, 10, 0, 0, 74, 75, 7, 2, 0, 0, 75, 76, 7, 7, 0, 0,
		76, 16, 1, 0, 0, 0, 77, 78, 7, 8, 0, 0, 78, 79, 7, 6, 0, 0, 79, 80, 7,
		7, 0, 0, 80, 81, 7, 3, 0, 0, 81, 18, 1, 0, 0, 0, 82, 83, 5, 58, 0, 0, 83,
		20, 1, 0, 0, 0, 84, 85, 7, 3, 0, 0, 85, 86, 7, 11, 0, 0, 86, 87, 7, 12,
		0, 0, 87, 88, 7, 13, 0, 0, 88, 89, 7, 13, 0, 0, 89, 22, 1, 0, 0, 0, 90,
		91, 7, 14, 0, 0, 91, 92, 7, 2, 0, 0, 92, 93, 7, 5, 0, 0, 93, 94, 7, 6,
		0, 0, 94, 95, 7, 11, 0, 0, 95, 96, 7, 4, 0, 0, 96, 24, 1, 0, 0, 0, 97,
		98, 7, 10, 0, 0, 98, 99, 7, 7, 0, 0, 99, 100, 7, 2, 0, 0, 100, 101, 7,
		6, 0, 0, 101, 111, 7, 9, 0, 0, 102, 103, 7, 15, 0, 0, 103, 104, 7, 12,
		0, 0, 104, 111, 7, 1, 0, 0, 105, 106, 7, 3, 0, 0, 106, 107, 7, 2, 0, 0,
		107, 108, 7, 7, 0, 0, 108, 109, 7, 13, 0, 0, 109, 111, 7, 9, 0, 0, 110,
		97, 1, 0, 0, 0, 110, 102, 1, 0, 0, 0, 110, 105, 1, 0, 0, 0, 111, 26, 1,
		0, 0, 0, 112, 114, 7, 16, 0, 0, 113, 112, 1, 0, 0, 0, 114, 115, 1, 0, 0,
		0, 115, 113, 1, 0, 0, 0, 115, 116, 1, 0, 0, 0, 116, 117, 1, 0, 0, 0, 117,
		118, 6, 13, 0, 0, 118, 28, 1, 0, 0, 0, 119, 121, 7, 17, 0, 0, 120, 119,
		1, 0, 0, 0, 120, 121, 1, 0, 0, 0, 121, 123, 1, 0, 0, 0, 122, 124, 7, 18,
		0, 0, 123, 122, 1, 0, 0, 0, 124, 125, 1, 0, 0, 0, 125, 123, 1, 0, 0, 0,
		125, 126, 1, 0, 0, 0, 126, 30, 1, 0, 0, 0, 127, 129, 7, 17, 0, 0, 128,
		127, 1, 0, 0, 0, 128, 129, 1, 0, 0, 0, 129, 131, 1, 0, 0, 0, 130, 132,
		7, 18, 0, 0, 131, 130, 1, 0, 0, 0, 132, 133, 1, 0, 0, 0, 133, 131, 1, 0,
		0, 0, 133, 134, 1, 0, 0, 0, 134, 135, 1, 0, 0, 0, 135, 137, 5, 46, 0, 0,
		136, 138, 7, 18, 0, 0, 137, 136, 1, 0, 0, 0, 138, 139, 1, 0, 0, 0, 139,
		137, 1, 0, 0, 0, 139, 140, 1, 0, 0, 0, 140, 32, 1, 0, 0, 0, 141, 145, 7,
		19, 0, 0, 142, 144, 7, 20, 0, 0, 143, 142, 1, 0, 0, 0, 144, 147, 1, 0,
		0, 0, 145, 143, 1, 0, 0, 0, 145, 146, 1, 0, 0, 0, 146, 34, 1, 0, 0, 0,
		147, 145, 1, 0, 0, 0, 9, 0, 110, 115, 120, 125, 128, 133, 139, 145, 1,
		6, 0, 0,
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

// GoldLexerInit initializes any static state used to implement GoldLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewGoldLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func GoldLexerInit() {
	staticData := &GoldLexerLexerStaticData
	staticData.once.Do(goldlexerLexerInit)
}

// NewGoldLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewGoldLexer(input antlr.CharStream) *GoldLexer {
	GoldLexerInit()
	l := new(GoldLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &GoldLexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	l.channelNames = staticData.ChannelNames
	l.modeNames = staticData.ModeNames
	l.RuleNames = staticData.RuleNames
	l.LiteralNames = staticData.LiteralNames
	l.SymbolicNames = staticData.SymbolicNames
	l.GrammarFileName = "Gold.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// GoldLexer tokens.
const (
	GoldLexerT__0               = 1
	GoldLexerT__1               = 2
	GoldLexerT__2               = 3
	GoldLexerT__3               = 4
	GoldLexerT__4               = 5
	GoldLexerT__5               = 6
	GoldLexerT__6               = 7
	GoldLexerT__7               = 8
	GoldLexerT__8               = 9
	GoldLexerT__9               = 10
	GoldLexerT__10              = 11
	GoldLexerT__11              = 12
	GoldLexerPARAMETER_MODIFIER = 13
	GoldLexerWHITESPACE         = 14
	GoldLexerINT_LITERAL        = 15
	GoldLexerDEICMAL_LITERAL    = 16
	GoldLexerID                 = 17
)
