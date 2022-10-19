// Code generated from rtsp_request.g4 by ANTLR 4.10.1. DO NOT EDIT.

package request

import (
	"fmt"
	"sync"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter

type rtsp_requestLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var rtsp_requestlexerLexerStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	channelNames           []string
	modeNames              []string
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func rtsp_requestlexerLexerInit() {
	staticData := &rtsp_requestlexerLexerStaticData
	staticData.channelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.modeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.literalNames = []string{
		"", "", "", "'://'", "", "':'", "'/'", "'.'", "", "", "", "", "", "",
		"' '", "'\\r'", "'\\n'",
	}
	staticData.symbolicNames = []string{
		"", "RTSP_VERSION", "URI_SCHEME", "URI_DELIMETER", "IP", "HCOLON", "SLASH",
		"DOT", "LETTER", "TOKEN", "TEXT", "INT", "ID", "ALPHA", "SP", "CR",
		"LF", "CRLF", "CTL",
	}
	staticData.ruleNames = []string{
		"RTSP_VERSION", "URI_SCHEME", "URI_DELIMETER", "IP", "HCOLON", "SLASH",
		"DOT", "DIGIT", "LETTER", "TOKEN", "TEXT", "INT", "ID", "ALPHA", "SP",
		"CR", "LF", "CRLF", "CTL",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 18, 114, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 1, 0, 1, 0, 1, 0, 1, 0,
		1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 3, 1, 59, 8, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1,
		3, 1, 3, 1, 3, 1, 3, 1, 4, 1, 4, 1, 5, 1, 5, 1, 6, 1, 6, 1, 7, 1, 7, 1,
		8, 1, 8, 1, 9, 1, 9, 1, 10, 1, 10, 1, 11, 4, 11, 86, 8, 11, 11, 11, 12,
		11, 87, 1, 12, 1, 12, 1, 12, 5, 12, 93, 8, 12, 10, 12, 12, 12, 96, 9, 12,
		1, 13, 1, 13, 1, 14, 1, 14, 1, 15, 1, 15, 1, 16, 1, 16, 1, 17, 1, 17, 1,
		17, 1, 17, 1, 17, 3, 17, 111, 8, 17, 1, 18, 1, 18, 1, 94, 0, 19, 1, 1,
		3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13, 7, 15, 0, 17, 8, 19, 9, 21, 10, 23,
		11, 25, 12, 27, 13, 29, 14, 31, 15, 33, 16, 35, 17, 37, 18, 1, 0, 6, 1,
		0, 48, 57, 2, 0, 65, 90, 97, 122, 8, 0, 33, 33, 35, 39, 45, 46, 48, 57,
		65, 90, 94, 122, 124, 124, 126, 126, 1, 0, 33, 126, 3, 0, 48, 57, 65, 90,
		97, 122, 1, 0, 0, 31, 118, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1,
		0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13,
		1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0, 21, 1, 0, 0, 0, 0,
		23, 1, 0, 0, 0, 0, 25, 1, 0, 0, 0, 0, 27, 1, 0, 0, 0, 0, 29, 1, 0, 0, 0,
		0, 31, 1, 0, 0, 0, 0, 33, 1, 0, 0, 0, 0, 35, 1, 0, 0, 0, 0, 37, 1, 0, 0,
		0, 1, 39, 1, 0, 0, 0, 3, 58, 1, 0, 0, 0, 5, 60, 1, 0, 0, 0, 7, 64, 1, 0,
		0, 0, 9, 70, 1, 0, 0, 0, 11, 72, 1, 0, 0, 0, 13, 74, 1, 0, 0, 0, 15, 76,
		1, 0, 0, 0, 17, 78, 1, 0, 0, 0, 19, 80, 1, 0, 0, 0, 21, 82, 1, 0, 0, 0,
		23, 85, 1, 0, 0, 0, 25, 89, 1, 0, 0, 0, 27, 97, 1, 0, 0, 0, 29, 99, 1,
		0, 0, 0, 31, 101, 1, 0, 0, 0, 33, 103, 1, 0, 0, 0, 35, 110, 1, 0, 0, 0,
		37, 112, 1, 0, 0, 0, 39, 40, 5, 82, 0, 0, 40, 41, 5, 84, 0, 0, 41, 42,
		5, 83, 0, 0, 42, 43, 5, 80, 0, 0, 43, 44, 5, 47, 0, 0, 44, 45, 1, 0, 0,
		0, 45, 46, 3, 23, 11, 0, 46, 47, 5, 46, 0, 0, 47, 48, 3, 23, 11, 0, 48,
		2, 1, 0, 0, 0, 49, 50, 5, 114, 0, 0, 50, 51, 5, 116, 0, 0, 51, 52, 5, 115,
		0, 0, 52, 59, 5, 112, 0, 0, 53, 54, 5, 114, 0, 0, 54, 55, 5, 116, 0, 0,
		55, 56, 5, 115, 0, 0, 56, 57, 5, 112, 0, 0, 57, 59, 5, 117, 0, 0, 58, 49,
		1, 0, 0, 0, 58, 53, 1, 0, 0, 0, 59, 4, 1, 0, 0, 0, 60, 61, 5, 58, 0, 0,
		61, 62, 5, 47, 0, 0, 62, 63, 5, 47, 0, 0, 63, 6, 1, 0, 0, 0, 64, 65, 3,
		23, 11, 0, 65, 66, 5, 46, 0, 0, 66, 67, 3, 23, 11, 0, 67, 68, 5, 46, 0,
		0, 68, 69, 3, 23, 11, 0, 69, 8, 1, 0, 0, 0, 70, 71, 5, 58, 0, 0, 71, 10,
		1, 0, 0, 0, 72, 73, 5, 47, 0, 0, 73, 12, 1, 0, 0, 0, 74, 75, 5, 46, 0,
		0, 75, 14, 1, 0, 0, 0, 76, 77, 7, 0, 0, 0, 77, 16, 1, 0, 0, 0, 78, 79,
		7, 1, 0, 0, 79, 18, 1, 0, 0, 0, 80, 81, 7, 2, 0, 0, 81, 20, 1, 0, 0, 0,
		82, 83, 7, 3, 0, 0, 83, 22, 1, 0, 0, 0, 84, 86, 3, 15, 7, 0, 85, 84, 1,
		0, 0, 0, 86, 87, 1, 0, 0, 0, 87, 85, 1, 0, 0, 0, 87, 88, 1, 0, 0, 0, 88,
		24, 1, 0, 0, 0, 89, 94, 3, 17, 8, 0, 90, 93, 3, 17, 8, 0, 91, 93, 3, 15,
		7, 0, 92, 90, 1, 0, 0, 0, 92, 91, 1, 0, 0, 0, 93, 96, 1, 0, 0, 0, 94, 95,
		1, 0, 0, 0, 94, 92, 1, 0, 0, 0, 95, 26, 1, 0, 0, 0, 96, 94, 1, 0, 0, 0,
		97, 98, 7, 4, 0, 0, 98, 28, 1, 0, 0, 0, 99, 100, 5, 32, 0, 0, 100, 30,
		1, 0, 0, 0, 101, 102, 5, 13, 0, 0, 102, 32, 1, 0, 0, 0, 103, 104, 5, 10,
		0, 0, 104, 34, 1, 0, 0, 0, 105, 111, 3, 31, 15, 0, 106, 111, 3, 33, 16,
		0, 107, 108, 3, 31, 15, 0, 108, 109, 3, 33, 16, 0, 109, 111, 1, 0, 0, 0,
		110, 105, 1, 0, 0, 0, 110, 106, 1, 0, 0, 0, 110, 107, 1, 0, 0, 0, 111,
		36, 1, 0, 0, 0, 112, 113, 7, 5, 0, 0, 113, 38, 1, 0, 0, 0, 6, 0, 58, 87,
		92, 94, 110, 0,
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

// rtsp_requestLexerInit initializes any static state used to implement rtsp_requestLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// Newrtsp_requestLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func Rtsp_requestLexerInit() {
	staticData := &rtsp_requestlexerLexerStaticData
	staticData.once.Do(rtsp_requestlexerLexerInit)
}

// Newrtsp_requestLexer produces a new lexer instance for the optional input antlr.CharStream.
func Newrtsp_requestLexer(input antlr.CharStream) *rtsp_requestLexer {
	Rtsp_requestLexerInit()
	l := new(rtsp_requestLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &rtsp_requestlexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	l.channelNames = staticData.channelNames
	l.modeNames = staticData.modeNames
	l.RuleNames = staticData.ruleNames
	l.LiteralNames = staticData.literalNames
	l.SymbolicNames = staticData.symbolicNames
	l.GrammarFileName = "rtsp_request.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// rtsp_requestLexer tokens.
const (
	rtsp_requestLexerRTSP_VERSION  = 1
	rtsp_requestLexerURI_SCHEME    = 2
	rtsp_requestLexerURI_DELIMETER = 3
	rtsp_requestLexerIP            = 4
	rtsp_requestLexerHCOLON        = 5
	rtsp_requestLexerSLASH         = 6
	rtsp_requestLexerDOT           = 7
	rtsp_requestLexerLETTER        = 8
	rtsp_requestLexerTOKEN         = 9
	rtsp_requestLexerTEXT          = 10
	rtsp_requestLexerINT           = 11
	rtsp_requestLexerID            = 12
	rtsp_requestLexerALPHA         = 13
	rtsp_requestLexerSP            = 14
	rtsp_requestLexerCR            = 15
	rtsp_requestLexerLF            = 16
	rtsp_requestLexerCRLF          = 17
	rtsp_requestLexerCTL           = 18
)
