// Code generated from LangMak.g4 by ANTLR 4.13.2. DO NOT EDIT.

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

type LangMakLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var LangMakLexerLexerStaticData struct {
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

func langmaklexerLexerInit() {
	staticData := &LangMakLexerLexerStaticData
	staticData.ChannelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.ModeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.LiteralNames = []string{
		"", "';'", "'print'", "'while'", "'{'", "'}'", "'if'", "'else'", "'('",
		"')'", "'*'", "'/'", "'+'", "'-'", "'='", "'>'", "'<'", "'=='",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "", "LPAREN", "RPAREN", "MULT", "DIV", "PLUS",
		"MINUS", "ASSIGN", "GREAT", "LESS", "EQUAL", "ID", "INT", "STRING",
		"WS", "COMMENT", "LINE_COMMENT",
	}
	staticData.RuleNames = []string{
		"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "LPAREN", "RPAREN",
		"MULT", "DIV", "PLUS", "MINUS", "ASSIGN", "GREAT", "LESS", "EQUAL",
		"ID", "INT", "STRING", "WS", "COMMENT", "LINE_COMMENT",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 23, 149, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1, 4, 1, 4, 1,
		5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 7, 1, 7, 1, 8, 1, 8, 1,
		9, 1, 9, 1, 10, 1, 10, 1, 11, 1, 11, 1, 12, 1, 12, 1, 13, 1, 13, 1, 14,
		1, 14, 1, 15, 1, 15, 1, 16, 1, 16, 1, 16, 1, 17, 1, 17, 5, 17, 97, 8, 17,
		10, 17, 12, 17, 100, 9, 17, 1, 18, 4, 18, 103, 8, 18, 11, 18, 12, 18, 104,
		1, 19, 1, 19, 1, 19, 1, 19, 5, 19, 111, 8, 19, 10, 19, 12, 19, 114, 9,
		19, 1, 19, 1, 19, 1, 20, 4, 20, 119, 8, 20, 11, 20, 12, 20, 120, 1, 20,
		1, 20, 1, 21, 1, 21, 1, 21, 1, 21, 5, 21, 129, 8, 21, 10, 21, 12, 21, 132,
		9, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 22, 1, 22, 1, 22, 1, 22, 5,
		22, 143, 8, 22, 10, 22, 12, 22, 146, 9, 22, 1, 22, 1, 22, 1, 130, 0, 23,
		1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13, 7, 15, 8, 17, 9, 19, 10, 21, 11,
		23, 12, 25, 13, 27, 14, 29, 15, 31, 16, 33, 17, 35, 18, 37, 19, 39, 20,
		41, 21, 43, 22, 45, 23, 1, 0, 6, 3, 0, 65, 90, 95, 95, 97, 122, 4, 0, 48,
		57, 65, 90, 95, 95, 97, 122, 1, 0, 48, 57, 2, 0, 34, 34, 92, 92, 3, 0,
		9, 10, 13, 13, 32, 32, 2, 0, 10, 10, 13, 13, 155, 0, 1, 1, 0, 0, 0, 0,
		3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0,
		11, 1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0, 0, 0,
		0, 19, 1, 0, 0, 0, 0, 21, 1, 0, 0, 0, 0, 23, 1, 0, 0, 0, 0, 25, 1, 0, 0,
		0, 0, 27, 1, 0, 0, 0, 0, 29, 1, 0, 0, 0, 0, 31, 1, 0, 0, 0, 0, 33, 1, 0,
		0, 0, 0, 35, 1, 0, 0, 0, 0, 37, 1, 0, 0, 0, 0, 39, 1, 0, 0, 0, 0, 41, 1,
		0, 0, 0, 0, 43, 1, 0, 0, 0, 0, 45, 1, 0, 0, 0, 1, 47, 1, 0, 0, 0, 3, 49,
		1, 0, 0, 0, 5, 55, 1, 0, 0, 0, 7, 61, 1, 0, 0, 0, 9, 63, 1, 0, 0, 0, 11,
		65, 1, 0, 0, 0, 13, 68, 1, 0, 0, 0, 15, 73, 1, 0, 0, 0, 17, 75, 1, 0, 0,
		0, 19, 77, 1, 0, 0, 0, 21, 79, 1, 0, 0, 0, 23, 81, 1, 0, 0, 0, 25, 83,
		1, 0, 0, 0, 27, 85, 1, 0, 0, 0, 29, 87, 1, 0, 0, 0, 31, 89, 1, 0, 0, 0,
		33, 91, 1, 0, 0, 0, 35, 94, 1, 0, 0, 0, 37, 102, 1, 0, 0, 0, 39, 106, 1,
		0, 0, 0, 41, 118, 1, 0, 0, 0, 43, 124, 1, 0, 0, 0, 45, 138, 1, 0, 0, 0,
		47, 48, 5, 59, 0, 0, 48, 2, 1, 0, 0, 0, 49, 50, 5, 112, 0, 0, 50, 51, 5,
		114, 0, 0, 51, 52, 5, 105, 0, 0, 52, 53, 5, 110, 0, 0, 53, 54, 5, 116,
		0, 0, 54, 4, 1, 0, 0, 0, 55, 56, 5, 119, 0, 0, 56, 57, 5, 104, 0, 0, 57,
		58, 5, 105, 0, 0, 58, 59, 5, 108, 0, 0, 59, 60, 5, 101, 0, 0, 60, 6, 1,
		0, 0, 0, 61, 62, 5, 123, 0, 0, 62, 8, 1, 0, 0, 0, 63, 64, 5, 125, 0, 0,
		64, 10, 1, 0, 0, 0, 65, 66, 5, 105, 0, 0, 66, 67, 5, 102, 0, 0, 67, 12,
		1, 0, 0, 0, 68, 69, 5, 101, 0, 0, 69, 70, 5, 108, 0, 0, 70, 71, 5, 115,
		0, 0, 71, 72, 5, 101, 0, 0, 72, 14, 1, 0, 0, 0, 73, 74, 5, 40, 0, 0, 74,
		16, 1, 0, 0, 0, 75, 76, 5, 41, 0, 0, 76, 18, 1, 0, 0, 0, 77, 78, 5, 42,
		0, 0, 78, 20, 1, 0, 0, 0, 79, 80, 5, 47, 0, 0, 80, 22, 1, 0, 0, 0, 81,
		82, 5, 43, 0, 0, 82, 24, 1, 0, 0, 0, 83, 84, 5, 45, 0, 0, 84, 26, 1, 0,
		0, 0, 85, 86, 5, 61, 0, 0, 86, 28, 1, 0, 0, 0, 87, 88, 5, 62, 0, 0, 88,
		30, 1, 0, 0, 0, 89, 90, 5, 60, 0, 0, 90, 32, 1, 0, 0, 0, 91, 92, 5, 61,
		0, 0, 92, 93, 5, 61, 0, 0, 93, 34, 1, 0, 0, 0, 94, 98, 7, 0, 0, 0, 95,
		97, 7, 1, 0, 0, 96, 95, 1, 0, 0, 0, 97, 100, 1, 0, 0, 0, 98, 96, 1, 0,
		0, 0, 98, 99, 1, 0, 0, 0, 99, 36, 1, 0, 0, 0, 100, 98, 1, 0, 0, 0, 101,
		103, 7, 2, 0, 0, 102, 101, 1, 0, 0, 0, 103, 104, 1, 0, 0, 0, 104, 102,
		1, 0, 0, 0, 104, 105, 1, 0, 0, 0, 105, 38, 1, 0, 0, 0, 106, 112, 5, 34,
		0, 0, 107, 108, 5, 92, 0, 0, 108, 111, 9, 0, 0, 0, 109, 111, 8, 3, 0, 0,
		110, 107, 1, 0, 0, 0, 110, 109, 1, 0, 0, 0, 111, 114, 1, 0, 0, 0, 112,
		110, 1, 0, 0, 0, 112, 113, 1, 0, 0, 0, 113, 115, 1, 0, 0, 0, 114, 112,
		1, 0, 0, 0, 115, 116, 5, 34, 0, 0, 116, 40, 1, 0, 0, 0, 117, 119, 7, 4,
		0, 0, 118, 117, 1, 0, 0, 0, 119, 120, 1, 0, 0, 0, 120, 118, 1, 0, 0, 0,
		120, 121, 1, 0, 0, 0, 121, 122, 1, 0, 0, 0, 122, 123, 6, 20, 0, 0, 123,
		42, 1, 0, 0, 0, 124, 125, 5, 47, 0, 0, 125, 126, 5, 42, 0, 0, 126, 130,
		1, 0, 0, 0, 127, 129, 9, 0, 0, 0, 128, 127, 1, 0, 0, 0, 129, 132, 1, 0,
		0, 0, 130, 131, 1, 0, 0, 0, 130, 128, 1, 0, 0, 0, 131, 133, 1, 0, 0, 0,
		132, 130, 1, 0, 0, 0, 133, 134, 5, 42, 0, 0, 134, 135, 5, 47, 0, 0, 135,
		136, 1, 0, 0, 0, 136, 137, 6, 21, 0, 0, 137, 44, 1, 0, 0, 0, 138, 139,
		5, 47, 0, 0, 139, 140, 5, 47, 0, 0, 140, 144, 1, 0, 0, 0, 141, 143, 8,
		5, 0, 0, 142, 141, 1, 0, 0, 0, 143, 146, 1, 0, 0, 0, 144, 142, 1, 0, 0,
		0, 144, 145, 1, 0, 0, 0, 145, 147, 1, 0, 0, 0, 146, 144, 1, 0, 0, 0, 147,
		148, 6, 22, 0, 0, 148, 46, 1, 0, 0, 0, 8, 0, 98, 104, 110, 112, 120, 130,
		144, 1, 6, 0, 0,
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

// LangMakLexerInit initializes any static state used to implement LangMakLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewLangMakLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func LangMakLexerInit() {
	staticData := &LangMakLexerLexerStaticData
	staticData.once.Do(langmaklexerLexerInit)
}

// NewLangMakLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewLangMakLexer(input antlr.CharStream) *LangMakLexer {
	LangMakLexerInit()
	l := new(LangMakLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &LangMakLexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	l.channelNames = staticData.ChannelNames
	l.modeNames = staticData.ModeNames
	l.RuleNames = staticData.RuleNames
	l.LiteralNames = staticData.LiteralNames
	l.SymbolicNames = staticData.SymbolicNames
	l.GrammarFileName = "LangMak.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// LangMakLexer tokens.
const (
	LangMakLexerT__0         = 1
	LangMakLexerT__1         = 2
	LangMakLexerT__2         = 3
	LangMakLexerT__3         = 4
	LangMakLexerT__4         = 5
	LangMakLexerT__5         = 6
	LangMakLexerT__6         = 7
	LangMakLexerLPAREN       = 8
	LangMakLexerRPAREN       = 9
	LangMakLexerMULT         = 10
	LangMakLexerDIV          = 11
	LangMakLexerPLUS         = 12
	LangMakLexerMINUS        = 13
	LangMakLexerASSIGN       = 14
	LangMakLexerGREAT        = 15
	LangMakLexerLESS         = 16
	LangMakLexerEQUAL        = 17
	LangMakLexerID           = 18
	LangMakLexerINT          = 19
	LangMakLexerSTRING       = 20
	LangMakLexerWS           = 21
	LangMakLexerCOMMENT      = 22
	LangMakLexerLINE_COMMENT = 23
)
