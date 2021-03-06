// Code generated from /home/tanx-pipefy/Desktop/UAM/22-Invierno/Traductores/tarea 4/RuLang2GoLang/Ru.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser

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


type RuLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames []string
	// TODO: EOF string
}

var rulexerLexerStaticData struct {
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

func rulexerLexerInit() {
  staticData := &rulexerLexerStaticData
  staticData.channelNames = []string{
    "DEFAULT_TOKEN_CHANNEL", "HIDDEN",
  }
  staticData.modeNames = []string{
    "DEFAULT_MODE",
  }
  staticData.literalNames = []string{
    "", "'||'", "'&&'", "'=='", "'!='", "'>'", "'<'", "'<='", "'>='", "'+'", 
    "'-'", "'*'", "'/'", "'%'", "'^'", "'!'", "';'", "'='", "'('", "')'", 
    "'{'", "'}'", "'true'", "'false'", "'nil'", "'if'", "'else'", "'while'", 
    "'log'", "'imprime'",
  }
  staticData.symbolicNames = []string{
    "", "OR", "AND", "IGUAL", "DIFQ", "MAYORQ", "MENORQ", "MENIG", "MAYIG", 
    "MAS", "MENOS", "MULT", "DIV", "MOD", "POW", "NOT", "PTOCOMA", "ASIGNA", 
    "APAR", "CPAR", "ALLAVE", "CLLAVE", "TRUE", "FALSE", "NIL", "IF", "ELSE", 
    "WHILE", "LOG", "IMPRIMIR", "ID", "INT", "FLOAT", "STRING", "COMENTARIO", 
    "ESPACIO", "OTRO",
  }
  staticData.ruleNames = []string{
    "OR", "AND", "IGUAL", "DIFQ", "MAYORQ", "MENORQ", "MENIG", "MAYIG", 
    "MAS", "MENOS", "MULT", "DIV", "MOD", "POW", "NOT", "PTOCOMA", "ASIGNA", 
    "APAR", "CPAR", "ALLAVE", "CLLAVE", "TRUE", "FALSE", "NIL", "IF", "ELSE", 
    "WHILE", "LOG", "IMPRIMIR", "ID", "INT", "FLOAT", "STRING", "COMENTARIO", 
    "ESPACIO", "OTRO",
  }
  staticData.predictionContextCache = antlr.NewPredictionContextCache()
  staticData.serializedATN = []int32{
	4, 0, 36, 220, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 
	4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 
	10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 
	7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 
	20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25, 
	2, 26, 7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2, 
	31, 7, 31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 1, 0, 
	1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 1, 4, 
	1, 4, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 1, 8, 1, 8, 1, 9, 
	1, 9, 1, 10, 1, 10, 1, 11, 1, 11, 1, 12, 1, 12, 1, 13, 1, 13, 1, 14, 1, 
	14, 1, 15, 1, 15, 1, 16, 1, 16, 1, 17, 1, 17, 1, 18, 1, 18, 1, 19, 1, 19, 
	1, 20, 1, 20, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 22, 1, 22, 1, 22, 1, 
	22, 1, 22, 1, 22, 1, 23, 1, 23, 1, 23, 1, 23, 1, 24, 1, 24, 1, 24, 1, 25, 
	1, 25, 1, 25, 1, 25, 1, 25, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 
	27, 1, 27, 1, 27, 1, 27, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 
	1, 28, 1, 29, 1, 29, 5, 29, 165, 8, 29, 10, 29, 12, 29, 168, 9, 29, 1, 
	30, 4, 30, 171, 8, 30, 11, 30, 12, 30, 172, 1, 31, 4, 31, 176, 8, 31, 11, 
	31, 12, 31, 177, 1, 31, 1, 31, 5, 31, 182, 8, 31, 10, 31, 12, 31, 185, 
	9, 31, 1, 31, 1, 31, 4, 31, 189, 8, 31, 11, 31, 12, 31, 190, 3, 31, 193, 
	8, 31, 1, 32, 1, 32, 1, 32, 1, 32, 5, 32, 199, 8, 32, 10, 32, 12, 32, 202, 
	9, 32, 1, 32, 1, 32, 1, 33, 1, 33, 5, 33, 208, 8, 33, 10, 33, 12, 33, 211, 
	9, 33, 1, 33, 1, 33, 1, 34, 1, 34, 1, 34, 1, 34, 1, 35, 1, 35, 0, 0, 36, 
	1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13, 7, 15, 8, 17, 9, 19, 10, 21, 11, 
	23, 12, 25, 13, 27, 14, 29, 15, 31, 16, 33, 17, 35, 18, 37, 19, 39, 20, 
	41, 21, 43, 22, 45, 23, 47, 24, 49, 25, 51, 26, 53, 27, 55, 28, 57, 29, 
	59, 30, 61, 31, 63, 32, 65, 33, 67, 34, 69, 35, 71, 36, 1, 0, 6, 3, 0, 
	65, 90, 95, 95, 97, 122, 4, 0, 48, 57, 65, 90, 95, 95, 97, 122, 1, 0, 48, 
	57, 3, 0, 10, 10, 13, 13, 34, 34, 2, 0, 10, 10, 13, 13, 3, 0, 9, 10, 13, 
	13, 32, 32, 228, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 
	0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0, 
	0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0, 21, 1, 0, 
	0, 0, 0, 23, 1, 0, 0, 0, 0, 25, 1, 0, 0, 0, 0, 27, 1, 0, 0, 0, 0, 29, 1, 
	0, 0, 0, 0, 31, 1, 0, 0, 0, 0, 33, 1, 0, 0, 0, 0, 35, 1, 0, 0, 0, 0, 37, 
	1, 0, 0, 0, 0, 39, 1, 0, 0, 0, 0, 41, 1, 0, 0, 0, 0, 43, 1, 0, 0, 0, 0, 
	45, 1, 0, 0, 0, 0, 47, 1, 0, 0, 0, 0, 49, 1, 0, 0, 0, 0, 51, 1, 0, 0, 0, 
	0, 53, 1, 0, 0, 0, 0, 55, 1, 0, 0, 0, 0, 57, 1, 0, 0, 0, 0, 59, 1, 0, 0, 
	0, 0, 61, 1, 0, 0, 0, 0, 63, 1, 0, 0, 0, 0, 65, 1, 0, 0, 0, 0, 67, 1, 0, 
	0, 0, 0, 69, 1, 0, 0, 0, 0, 71, 1, 0, 0, 0, 1, 73, 1, 0, 0, 0, 3, 76, 1, 
	0, 0, 0, 5, 79, 1, 0, 0, 0, 7, 82, 1, 0, 0, 0, 9, 85, 1, 0, 0, 0, 11, 87, 
	1, 0, 0, 0, 13, 89, 1, 0, 0, 0, 15, 92, 1, 0, 0, 0, 17, 95, 1, 0, 0, 0, 
	19, 97, 1, 0, 0, 0, 21, 99, 1, 0, 0, 0, 23, 101, 1, 0, 0, 0, 25, 103, 1, 
	0, 0, 0, 27, 105, 1, 0, 0, 0, 29, 107, 1, 0, 0, 0, 31, 109, 1, 0, 0, 0, 
	33, 111, 1, 0, 0, 0, 35, 113, 1, 0, 0, 0, 37, 115, 1, 0, 0, 0, 39, 117, 
	1, 0, 0, 0, 41, 119, 1, 0, 0, 0, 43, 121, 1, 0, 0, 0, 45, 126, 1, 0, 0, 
	0, 47, 132, 1, 0, 0, 0, 49, 136, 1, 0, 0, 0, 51, 139, 1, 0, 0, 0, 53, 144, 
	1, 0, 0, 0, 55, 150, 1, 0, 0, 0, 57, 154, 1, 0, 0, 0, 59, 162, 1, 0, 0, 
	0, 61, 170, 1, 0, 0, 0, 63, 192, 1, 0, 0, 0, 65, 194, 1, 0, 0, 0, 67, 205, 
	1, 0, 0, 0, 69, 214, 1, 0, 0, 0, 71, 218, 1, 0, 0, 0, 73, 74, 5, 124, 0, 
	0, 74, 75, 5, 124, 0, 0, 75, 2, 1, 0, 0, 0, 76, 77, 5, 38, 0, 0, 77, 78, 
	5, 38, 0, 0, 78, 4, 1, 0, 0, 0, 79, 80, 5, 61, 0, 0, 80, 81, 5, 61, 0, 
	0, 81, 6, 1, 0, 0, 0, 82, 83, 5, 33, 0, 0, 83, 84, 5, 61, 0, 0, 84, 8, 
	1, 0, 0, 0, 85, 86, 5, 62, 0, 0, 86, 10, 1, 0, 0, 0, 87, 88, 5, 60, 0, 
	0, 88, 12, 1, 0, 0, 0, 89, 90, 5, 60, 0, 0, 90, 91, 5, 61, 0, 0, 91, 14, 
	1, 0, 0, 0, 92, 93, 5, 62, 0, 0, 93, 94, 5, 61, 0, 0, 94, 16, 1, 0, 0, 
	0, 95, 96, 5, 43, 0, 0, 96, 18, 1, 0, 0, 0, 97, 98, 5, 45, 0, 0, 98, 20, 
	1, 0, 0, 0, 99, 100, 5, 42, 0, 0, 100, 22, 1, 0, 0, 0, 101, 102, 5, 47, 
	0, 0, 102, 24, 1, 0, 0, 0, 103, 104, 5, 37, 0, 0, 104, 26, 1, 0, 0, 0, 
	105, 106, 5, 94, 0, 0, 106, 28, 1, 0, 0, 0, 107, 108, 5, 33, 0, 0, 108, 
	30, 1, 0, 0, 0, 109, 110, 5, 59, 0, 0, 110, 32, 1, 0, 0, 0, 111, 112, 5, 
	61, 0, 0, 112, 34, 1, 0, 0, 0, 113, 114, 5, 40, 0, 0, 114, 36, 1, 0, 0, 
	0, 115, 116, 5, 41, 0, 0, 116, 38, 1, 0, 0, 0, 117, 118, 5, 123, 0, 0, 
	118, 40, 1, 0, 0, 0, 119, 120, 5, 125, 0, 0, 120, 42, 1, 0, 0, 0, 121, 
	122, 5, 116, 0, 0, 122, 123, 5, 114, 0, 0, 123, 124, 5, 117, 0, 0, 124, 
	125, 5, 101, 0, 0, 125, 44, 1, 0, 0, 0, 126, 127, 5, 102, 0, 0, 127, 128, 
	5, 97, 0, 0, 128, 129, 5, 108, 0, 0, 129, 130, 5, 115, 0, 0, 130, 131, 
	5, 101, 0, 0, 131, 46, 1, 0, 0, 0, 132, 133, 5, 110, 0, 0, 133, 134, 5, 
	105, 0, 0, 134, 135, 5, 108, 0, 0, 135, 48, 1, 0, 0, 0, 136, 137, 5, 105, 
	0, 0, 137, 138, 5, 102, 0, 0, 138, 50, 1, 0, 0, 0, 139, 140, 5, 101, 0, 
	0, 140, 141, 5, 108, 0, 0, 141, 142, 5, 115, 0, 0, 142, 143, 5, 101, 0, 
	0, 143, 52, 1, 0, 0, 0, 144, 145, 5, 119, 0, 0, 145, 146, 5, 104, 0, 0, 
	146, 147, 5, 105, 0, 0, 147, 148, 5, 108, 0, 0, 148, 149, 5, 101, 0, 0, 
	149, 54, 1, 0, 0, 0, 150, 151, 5, 108, 0, 0, 151, 152, 5, 111, 0, 0, 152, 
	153, 5, 103, 0, 0, 153, 56, 1, 0, 0, 0, 154, 155, 5, 105, 0, 0, 155, 156, 
	5, 109, 0, 0, 156, 157, 5, 112, 0, 0, 157, 158, 5, 114, 0, 0, 158, 159, 
	5, 105, 0, 0, 159, 160, 5, 109, 0, 0, 160, 161, 5, 101, 0, 0, 161, 58, 
	1, 0, 0, 0, 162, 166, 7, 0, 0, 0, 163, 165, 7, 1, 0, 0, 164, 163, 1, 0, 
	0, 0, 165, 168, 1, 0, 0, 0, 166, 164, 1, 0, 0, 0, 166, 167, 1, 0, 0, 0, 
	167, 60, 1, 0, 0, 0, 168, 166, 1, 0, 0, 0, 169, 171, 7, 2, 0, 0, 170, 169, 
	1, 0, 0, 0, 171, 172, 1, 0, 0, 0, 172, 170, 1, 0, 0, 0, 172, 173, 1, 0, 
	0, 0, 173, 62, 1, 0, 0, 0, 174, 176, 7, 2, 0, 0, 175, 174, 1, 0, 0, 0, 
	176, 177, 1, 0, 0, 0, 177, 175, 1, 0, 0, 0, 177, 178, 1, 0, 0, 0, 178, 
	179, 1, 0, 0, 0, 179, 183, 5, 46, 0, 0, 180, 182, 7, 2, 0, 0, 181, 180, 
	1, 0, 0, 0, 182, 185, 1, 0, 0, 0, 183, 181, 1, 0, 0, 0, 183, 184, 1, 0, 
	0, 0, 184, 193, 1, 0, 0, 0, 185, 183, 1, 0, 0, 0, 186, 188, 5, 46, 0, 0, 
	187, 189, 7, 2, 0, 0, 188, 187, 1, 0, 0, 0, 189, 190, 1, 0, 0, 0, 190, 
	188, 1, 0, 0, 0, 190, 191, 1, 0, 0, 0, 191, 193, 1, 0, 0, 0, 192, 175, 
	1, 0, 0, 0, 192, 186, 1, 0, 0, 0, 193, 64, 1, 0, 0, 0, 194, 200, 5, 34, 
	0, 0, 195, 199, 8, 3, 0, 0, 196, 197, 5, 34, 0, 0, 197, 199, 5, 34, 0, 
	0, 198, 195, 1, 0, 0, 0, 198, 196, 1, 0, 0, 0, 199, 202, 1, 0, 0, 0, 200, 
	198, 1, 0, 0, 0, 200, 201, 1, 0, 0, 0, 201, 203, 1, 0, 0, 0, 202, 200, 
	1, 0, 0, 0, 203, 204, 5, 34, 0, 0, 204, 66, 1, 0, 0, 0, 205, 209, 5, 35, 
	0, 0, 206, 208, 8, 4, 0, 0, 207, 206, 1, 0, 0, 0, 208, 211, 1, 0, 0, 0, 
	209, 207, 1, 0, 0, 0, 209, 210, 1, 0, 0, 0, 210, 212, 1, 0, 0, 0, 211, 
	209, 1, 0, 0, 0, 212, 213, 6, 33, 0, 0, 213, 68, 1, 0, 0, 0, 214, 215, 
	7, 5, 0, 0, 215, 216, 1, 0, 0, 0, 216, 217, 6, 34, 0, 0, 217, 70, 1, 0, 
	0, 0, 218, 219, 9, 0, 0, 0, 219, 72, 1, 0, 0, 0, 10, 0, 166, 172, 177, 
	183, 190, 192, 198, 200, 209, 1, 6, 0, 0,
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

// RuLexerInit initializes any static state used to implement RuLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewRuLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func RuLexerInit() {
  staticData := &rulexerLexerStaticData
  staticData.once.Do(rulexerLexerInit)
}

// NewRuLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewRuLexer(input antlr.CharStream) *RuLexer {
  RuLexerInit()
	l := new(RuLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
  staticData := &rulexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	l.channelNames = staticData.channelNames
	l.modeNames = staticData.modeNames
	l.RuleNames = staticData.ruleNames
	l.LiteralNames = staticData.literalNames
	l.SymbolicNames = staticData.symbolicNames
	l.GrammarFileName = "Ru.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// RuLexer tokens.
const (
	RuLexerOR = 1
	RuLexerAND = 2
	RuLexerIGUAL = 3
	RuLexerDIFQ = 4
	RuLexerMAYORQ = 5
	RuLexerMENORQ = 6
	RuLexerMENIG = 7
	RuLexerMAYIG = 8
	RuLexerMAS = 9
	RuLexerMENOS = 10
	RuLexerMULT = 11
	RuLexerDIV = 12
	RuLexerMOD = 13
	RuLexerPOW = 14
	RuLexerNOT = 15
	RuLexerPTOCOMA = 16
	RuLexerASIGNA = 17
	RuLexerAPAR = 18
	RuLexerCPAR = 19
	RuLexerALLAVE = 20
	RuLexerCLLAVE = 21
	RuLexerTRUE = 22
	RuLexerFALSE = 23
	RuLexerNIL = 24
	RuLexerIF = 25
	RuLexerELSE = 26
	RuLexerWHILE = 27
	RuLexerLOG = 28
	RuLexerIMPRIMIR = 29
	RuLexerID = 30
	RuLexerINT = 31
	RuLexerFLOAT = 32
	RuLexerSTRING = 33
	RuLexerCOMENTARIO = 34
	RuLexerESPACIO = 35
	RuLexerOTRO = 36
)

