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
		"", "'class'", "'('", "')'", "'module'", "'procedure'", "'proc'", "','",
		"'end'", "'endproc'", "'function'", "'func'", "':'", "'var'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "PARAMETER_MODIFIER",
		"WHITESPACE", "INT_LITERAL", "DEICMAL_LITERAL", "STRING_LITERAL", "ID",
	}
	staticData.RuleNames = []string{
		"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "T__8",
		"T__9", "T__10", "T__11", "T__12", "PARAMETER_MODIFIER", "WHITESPACE",
		"INT_LITERAL", "DEICMAL_LITERAL", "STRING_LITERAL", "ID",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 19, 165, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 1, 0, 1, 0, 1, 0, 1, 0,
		1, 0, 1, 0, 1, 1, 1, 1, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3,
		1, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 5,
		1, 5, 1, 5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 1, 7, 1, 8, 1, 8,
		1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9,
		1, 9, 1, 9, 1, 9, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 11, 1, 11, 1, 12,
		1, 12, 1, 12, 1, 12, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1,
		13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 3, 13, 119, 8, 13, 1, 14, 4, 14,
		122, 8, 14, 11, 14, 12, 14, 123, 1, 14, 1, 14, 1, 15, 3, 15, 129, 8, 15,
		1, 15, 4, 15, 132, 8, 15, 11, 15, 12, 15, 133, 1, 16, 3, 16, 137, 8, 16,
		1, 16, 4, 16, 140, 8, 16, 11, 16, 12, 16, 141, 1, 16, 1, 16, 4, 16, 146,
		8, 16, 11, 16, 12, 16, 147, 1, 17, 1, 17, 5, 17, 152, 8, 17, 10, 17, 12,
		17, 155, 9, 17, 1, 17, 1, 17, 1, 18, 1, 18, 5, 18, 161, 8, 18, 10, 18,
		12, 18, 164, 9, 18, 1, 153, 0, 19, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6,
		13, 7, 15, 8, 17, 9, 19, 10, 21, 11, 23, 12, 25, 13, 27, 14, 29, 15, 31,
		16, 33, 17, 35, 18, 37, 19, 1, 0, 21, 2, 0, 67, 67, 99, 99, 2, 0, 76, 76,
		108, 108, 2, 0, 65, 65, 97, 97, 2, 0, 83, 83, 115, 115, 2, 0, 77, 77, 109,
		109, 2, 0, 79, 79, 111, 111, 2, 0, 68, 68, 100, 100, 2, 0, 85, 85, 117,
		117, 2, 0, 69, 69, 101, 101, 2, 0, 80, 80, 112, 112, 2, 0, 82, 82, 114,
		114, 2, 0, 78, 78, 110, 110, 2, 0, 70, 70, 102, 102, 2, 0, 84, 84, 116,
		116, 2, 0, 73, 73, 105, 105, 2, 0, 86, 86, 118, 118, 3, 0, 9, 10, 13, 13,
		32, 32, 2, 0, 43, 43, 45, 45, 1, 0, 48, 57, 3, 0, 65, 90, 95, 95, 97, 122,
		4, 0, 48, 57, 65, 90, 95, 95, 97, 122, 174, 0, 1, 1, 0, 0, 0, 0, 3, 1,
		0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1,
		0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19,
		1, 0, 0, 0, 0, 21, 1, 0, 0, 0, 0, 23, 1, 0, 0, 0, 0, 25, 1, 0, 0, 0, 0,
		27, 1, 0, 0, 0, 0, 29, 1, 0, 0, 0, 0, 31, 1, 0, 0, 0, 0, 33, 1, 0, 0, 0,
		0, 35, 1, 0, 0, 0, 0, 37, 1, 0, 0, 0, 1, 39, 1, 0, 0, 0, 3, 45, 1, 0, 0,
		0, 5, 47, 1, 0, 0, 0, 7, 49, 1, 0, 0, 0, 9, 56, 1, 0, 0, 0, 11, 66, 1,
		0, 0, 0, 13, 71, 1, 0, 0, 0, 15, 73, 1, 0, 0, 0, 17, 77, 1, 0, 0, 0, 19,
		85, 1, 0, 0, 0, 21, 94, 1, 0, 0, 0, 23, 99, 1, 0, 0, 0, 25, 101, 1, 0,
		0, 0, 27, 118, 1, 0, 0, 0, 29, 121, 1, 0, 0, 0, 31, 128, 1, 0, 0, 0, 33,
		136, 1, 0, 0, 0, 35, 149, 1, 0, 0, 0, 37, 158, 1, 0, 0, 0, 39, 40, 7, 0,
		0, 0, 40, 41, 7, 1, 0, 0, 41, 42, 7, 2, 0, 0, 42, 43, 7, 3, 0, 0, 43, 44,
		7, 3, 0, 0, 44, 2, 1, 0, 0, 0, 45, 46, 5, 40, 0, 0, 46, 4, 1, 0, 0, 0,
		47, 48, 5, 41, 0, 0, 48, 6, 1, 0, 0, 0, 49, 50, 7, 4, 0, 0, 50, 51, 7,
		5, 0, 0, 51, 52, 7, 6, 0, 0, 52, 53, 7, 7, 0, 0, 53, 54, 7, 1, 0, 0, 54,
		55, 7, 8, 0, 0, 55, 8, 1, 0, 0, 0, 56, 57, 7, 9, 0, 0, 57, 58, 7, 10, 0,
		0, 58, 59, 7, 5, 0, 0, 59, 60, 7, 0, 0, 0, 60, 61, 7, 8, 0, 0, 61, 62,
		7, 6, 0, 0, 62, 63, 7, 7, 0, 0, 63, 64, 7, 10, 0, 0, 64, 65, 7, 8, 0, 0,
		65, 10, 1, 0, 0, 0, 66, 67, 7, 9, 0, 0, 67, 68, 7, 10, 0, 0, 68, 69, 7,
		5, 0, 0, 69, 70, 7, 0, 0, 0, 70, 12, 1, 0, 0, 0, 71, 72, 5, 44, 0, 0, 72,
		14, 1, 0, 0, 0, 73, 74, 7, 8, 0, 0, 74, 75, 7, 11, 0, 0, 75, 76, 7, 6,
		0, 0, 76, 16, 1, 0, 0, 0, 77, 78, 7, 8, 0, 0, 78, 79, 7, 11, 0, 0, 79,
		80, 7, 6, 0, 0, 80, 81, 7, 9, 0, 0, 81, 82, 7, 10, 0, 0, 82, 83, 7, 5,
		0, 0, 83, 84, 7, 0, 0, 0, 84, 18, 1, 0, 0, 0, 85, 86, 7, 12, 0, 0, 86,
		87, 7, 7, 0, 0, 87, 88, 7, 11, 0, 0, 88, 89, 7, 0, 0, 0, 89, 90, 7, 13,
		0, 0, 90, 91, 7, 14, 0, 0, 91, 92, 7, 5, 0, 0, 92, 93, 7, 11, 0, 0, 93,
		20, 1, 0, 0, 0, 94, 95, 7, 12, 0, 0, 95, 96, 7, 7, 0, 0, 96, 97, 7, 11,
		0, 0, 97, 98, 7, 0, 0, 0, 98, 22, 1, 0, 0, 0, 99, 100, 5, 58, 0, 0, 100,
		24, 1, 0, 0, 0, 101, 102, 7, 15, 0, 0, 102, 103, 7, 2, 0, 0, 103, 104,
		7, 10, 0, 0, 104, 26, 1, 0, 0, 0, 105, 106, 7, 14, 0, 0, 106, 107, 7, 11,
		0, 0, 107, 108, 7, 5, 0, 0, 108, 109, 7, 7, 0, 0, 109, 119, 7, 13, 0, 0,
		110, 111, 7, 15, 0, 0, 111, 112, 7, 2, 0, 0, 112, 119, 7, 10, 0, 0, 113,
		114, 7, 0, 0, 0, 114, 115, 7, 5, 0, 0, 115, 116, 7, 11, 0, 0, 116, 117,
		7, 3, 0, 0, 117, 119, 7, 13, 0, 0, 118, 105, 1, 0, 0, 0, 118, 110, 1, 0,
		0, 0, 118, 113, 1, 0, 0, 0, 119, 28, 1, 0, 0, 0, 120, 122, 7, 16, 0, 0,
		121, 120, 1, 0, 0, 0, 122, 123, 1, 0, 0, 0, 123, 121, 1, 0, 0, 0, 123,
		124, 1, 0, 0, 0, 124, 125, 1, 0, 0, 0, 125, 126, 6, 14, 0, 0, 126, 30,
		1, 0, 0, 0, 127, 129, 7, 17, 0, 0, 128, 127, 1, 0, 0, 0, 128, 129, 1, 0,
		0, 0, 129, 131, 1, 0, 0, 0, 130, 132, 7, 18, 0, 0, 131, 130, 1, 0, 0, 0,
		132, 133, 1, 0, 0, 0, 133, 131, 1, 0, 0, 0, 133, 134, 1, 0, 0, 0, 134,
		32, 1, 0, 0, 0, 135, 137, 7, 17, 0, 0, 136, 135, 1, 0, 0, 0, 136, 137,
		1, 0, 0, 0, 137, 139, 1, 0, 0, 0, 138, 140, 7, 18, 0, 0, 139, 138, 1, 0,
		0, 0, 140, 141, 1, 0, 0, 0, 141, 139, 1, 0, 0, 0, 141, 142, 1, 0, 0, 0,
		142, 143, 1, 0, 0, 0, 143, 145, 5, 46, 0, 0, 144, 146, 7, 18, 0, 0, 145,
		144, 1, 0, 0, 0, 146, 147, 1, 0, 0, 0, 147, 145, 1, 0, 0, 0, 147, 148,
		1, 0, 0, 0, 148, 34, 1, 0, 0, 0, 149, 153, 5, 39, 0, 0, 150, 152, 9, 0,
		0, 0, 151, 150, 1, 0, 0, 0, 152, 155, 1, 0, 0, 0, 153, 154, 1, 0, 0, 0,
		153, 151, 1, 0, 0, 0, 154, 156, 1, 0, 0, 0, 155, 153, 1, 0, 0, 0, 156,
		157, 5, 39, 0, 0, 157, 36, 1, 0, 0, 0, 158, 162, 7, 19, 0, 0, 159, 161,
		7, 20, 0, 0, 160, 159, 1, 0, 0, 0, 161, 164, 1, 0, 0, 0, 162, 160, 1, 0,
		0, 0, 162, 163, 1, 0, 0, 0, 163, 38, 1, 0, 0, 0, 164, 162, 1, 0, 0, 0,
		10, 0, 118, 123, 128, 133, 136, 141, 147, 153, 162, 1, 6, 0, 0,
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
	GoldLexerT__12              = 13
	GoldLexerPARAMETER_MODIFIER = 14
	GoldLexerWHITESPACE         = 15
	GoldLexerINT_LITERAL        = 16
	GoldLexerDEICMAL_LITERAL    = 17
	GoldLexerSTRING_LITERAL     = 18
	GoldLexerID                 = 19
)
